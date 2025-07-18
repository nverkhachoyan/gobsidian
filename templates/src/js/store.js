// Global Alpine.js Stores
import ObsiGraph from "./libs/obsigraph.min.js";
import { loadJsonData, graphConfig } from "./utils.js";
import explorerComponent from "./explorer.js";
import { Document, Charset } from "./libs/flexsearch.js";

const graphCanvas = document.getElementById("graph-network-canvas");
const modalGraphCanvas = document.getElementById("graph-network-modal-canvas");

export function initializeStores(Alpine) {
  Alpine.data("explorerComponent", explorerComponent);
  Alpine.store("theme", {
    themeInitialized: false,
    current:
      localStorage.getItem("theme") ||
      (window.matchMedia("(prefers-color-scheme: dark)").matches
        ? "dark"
        : "light"),

    init() {
      this.setTheme(this.current, false);
      this.themeInitialized = true;
      window.dispatchEvent(
        new CustomEvent("theme-initialized", {
          detail: { theme: this.current },
        })
      );
    },

    toggle() {
      const newTheme = this.current === "light" ? "dark" : "light";
      this.setTheme(newTheme);
    },

    setTheme(newTheme, updateDOM = true) {
      this.current = newTheme;
      localStorage.setItem("theme", newTheme);

      // Update DOM
      if (updateDOM) {
        if (newTheme === "dark") {
          document.documentElement.classList.add("dark-mode");
        } else {
          document.documentElement.classList.remove("dark-mode");
        }
      }

      // Trigger custom event for other components
      window.dispatchEvent(
        new CustomEvent("theme-changed", {
          detail: { theme: newTheme },
        })
      );
    },
  });

  Alpine.store("fileTree", {
    data: null,
    expanded: localStorage.getItem("fileTreeExpanded") === "true",

    async init() {
      this.data = await loadJsonData("/generated/explorer.json");
    },

    toggle() {
      this.expanded = !this.expanded;
      localStorage.setItem("fileTreeExpanded", this.expanded);
    },

    setExpanded(state) {
      this.expanded = state;
      localStorage.setItem("fileTreeExpanded", state);
    },
  });

  Alpine.store("tags", {
    data: [],
    showAll: localStorage.getItem("tagsShowAll") === "true",

    async init() {
      this.data = await loadJsonData("/generated/tags.json");
    },

    toggle() {
      this.showAll = !this.showAll;
      localStorage.setItem("tagsShowAll", this.showAll);
    },

    getVisible() {
      return this.showAll ? this.data : this.data.slice(0, 6);
    },

    hasMore() {
      return this.data.length > 6;
    },
  });

  Alpine.store("graph", {
    data: null,
    modalOpen: false,
    modalHasBeenInitialized: false,
    graphObject: null,
    graphModalObject: null,

    async init() {
      this.data = await loadJsonData("/generated/graph.json");
      this.initSmallGraph();
    },

    initSmallGraph() {
      if (this.graphObject || !this.data) return;
      if (!Alpine.store("theme").themeInitialized) {
        // Defer initialization until theme is ready
        window.addEventListener(
          "theme-initialized",
          () => this.initSmallGraph(),
          { once: true }
        );
        return;
      }

      const nodes = this.data.nodes.map((node) => ({ ...node }));
      const edges = this.data.edges.map((edge) => ({ ...edge }));

      this.graphObject = new ObsiGraph(graphCanvas, nodes, edges, {
        ...graphConfig,
      });
      this.graphObject.setTheme(Alpine.store("theme").current);

      window.addEventListener("theme-changed", (e) => {
        this.graphObject.setTheme(e.detail.theme);
      });
    },

    openModal() {
      this.modalOpen = true;
      if (!this.modalHasBeenInitialized) {
        const nodes = this.data.nodes.map((node) => ({ ...node }));
        const edges = this.data.edges.map((edge) => ({ ...edge }));
        this.graphModalObject = new ObsiGraph(modalGraphCanvas, nodes, edges, {
          ...graphConfig,
        });
        this.graphModalObject.setTheme(Alpine.store("theme").current);
        this.modalHasBeenInitialized = true;

        window.addEventListener("theme-changed", (e) => {
          this.graphModalObject.setTheme(e.detail.theme);
        });
      }
    },

    closeModal() {
      this.modalOpen = false;
    },

    toggleModal() {
      this.modalOpen = !this.modalOpen;
    },
  });

  Alpine.store("sidebar", {
    mobileOpen: false,
    toggleMobile() {
      this.mobileOpen = !this.mobileOpen;
    },
    closeMobile() {
      this.mobileOpen = false;
    },
  });

  Alpine.store("search", {
    index: null,
    results: [],
    isOpen: false,
    query: "",
    loading: false,
    data: [],

    async init() {
      const data = await loadJsonData("/generated/search-index.json");
      this.data = data.notes;

      this.index = new Document({
        charset: Charset.LatinAdvanced,
        tokenize: "forward",
        resolution: 9,
        document: {
          id: "id",
          store: ["title", "body", "img", "url"],
          index: ["title", "body"],
        },
      });

      for (const record of this.data) {
        this.index.add(record);
      }
    },

    async search(query) {
      this.query = query;
      if (!query) {
        this.results = [];
        this.isOpen = false;
        return;
      }

      this.loading = true;
      this.isOpen = true;
      this.results = [];

      const searchResults = await this.index.searchAsync(query, {
        limit: 10,
        enrich: true,
      });

      // Flatten results from different fields
      const flatResults = {};
      searchResults.forEach((fieldResult) => {
        fieldResult.result.forEach((doc) => {
          if (!flatResults[doc.id]) {
            flatResults[doc.id] = { id: doc.id, ...doc.doc };
          }
        });
      });

      this.results = Object.values(flatResults);
      this.loading = false;
    },

    highlight(text) {
      if (!this.query || !text) {
        return text;
      }
      const query = this.query.replace(/[-\/\\^$*+?.()|[\]{}]/g, "\\$&");
      const regex = new RegExp(`(${query})`, "gi");
      return text.replace(regex, "<strong>$1</strong>");
    },

    open() {
      this.isOpen = true;
    },

    close() {
      this.isOpen = false;
      this.query = "";
      this.results = [];
      this.loading = false;
    },
  });

  Alpine.store("collapsible", {
    states: {},
    init(id, initialState = false) {
      if (this.states[id] === undefined) {
        const storedState = localStorage.getItem(`collapsible_${id}`);
        this.states[id] =
          storedState !== null ? JSON.parse(storedState) : initialState;
      }
    },
    toggle(id) {
      if (this.states[id] !== undefined) {
        this.states[id] = !this.states[id];
        localStorage.setItem(
          `collapsible_${id}`,
          JSON.stringify(this.states[id])
        );
        this.states = { ...this.states };
      }
    },
    isExpanded(id) {
      return this.states[id] === true;
    },
  });
}
