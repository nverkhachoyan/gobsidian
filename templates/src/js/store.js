// Global Alpine.js Stores
const graphCanvas = document.getElementById("graph-network-canvas");
const modalGraphCanvas = document.getElementById("graph-network-modal-canvas");
window.ObsiGraph = window.ObsiGraph.default || ObsiGraph.default;

const graphConfig = {
  initialZoomFactor: 0.7,
  zoomToFit: false,
  nodesHaveLinks: true,
  colors: {
    light: {
      node: 'rgb(73 77 83)',
      nodeHighlight: '#9165ea',
      nodeDimmed: 'rgb(229 229 229)',
      nodeStroke: '#b2b2b2',
      nodeDimmedStroke: 'rgb(229 229 229)',
      nodeHighlightStroke: '#9165ea',
      nodeFont: '#3f3f3f',
      edge: 'rgba(98, 98, 98, 1)',
      edgeHighlight: '#9165ea',
      background: '#ffffff',
    },
    dark: {
      node: '#b2b2b2',
      nodeHighlight: '#9165ea',
      nodeDimmed: 'rgb(38, 40, 41)',
      nodeStroke: '#b2b2b2',
      nodeDimmedStroke: 'rgb(38, 40, 41)',
      nodeHighlightStroke: '#9165ea',
      nodeFont: '#e8f4f8',
      edge: 'rgb(224, 224, 224)',
      edgeHighlight: 'rgb(145, 101, 234)',
      background: '#0d1117',
    },
  },
  sizing: {
    minNodeSize: 5,
    maxNodeSize: 50,
  },
  simulation: {
    linkDistance: 150,
    linkStrength: 0.8,
    chargeStrength: -300,
    collisionStrength: 0.2,
    alpha: 0.9,
    finalAlpha: 0.005,
    simulationDelay: 10000,
  },
  interaction: {
    zoomMin: 0.01,
    zoomMax: 5,
    labelShowThreshold: 0.8,
  },
  culling: {
    enabled: true,
    margin: 100,
    autoEnable: true,
    nodeThreshold: 50,
    edgeThreshold: 50,
  },
  font: 'sans-serif',
  disableTransitions: true,
}

document.addEventListener("alpine:init", () => {
  Alpine.store("theme", {
    current:
      localStorage.getItem("theme") ||
      (window.matchMedia("(prefers-color-scheme: dark)").matches
        ? "dark"
        : "light"),

    init() {
      this.setTheme(this.current);
    },

    toggle() {
      const newTheme = this.current === "light" ? "dark" : "light";
      this.setTheme(newTheme);
    },

    setTheme(newTheme) {
      this.current = newTheme;
      localStorage.setItem("theme", newTheme);

      // Update DOM
      if (newTheme === "dark") {
        document.documentElement.classList.add("dark-mode");
      } else {
        document.documentElement.classList.remove("dark-mode");
      }

      // Trigger custom event for other components
      window.dispatchEvent(
        new CustomEvent("theme-changed", {
          detail: { theme: newTheme },
        })
      );

      window.dispatchEvent(
        new CustomEvent("theme-initialized", {
          detail: { theme: newTheme },
        })
      );
    },
  });

  Alpine.store("fileTree", {
    data: window.fileTreeData || null,
    expanded: localStorage.getItem("fileTreeExpanded") === "true",

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
    data: window.tagsData || [],
    showAll: localStorage.getItem("tagsShowAll") === "true",

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
    data: window.graphData || null,
    modalOpen: false,
    modalHasBeenInitialized: false,
    graphObject: null,
    graphModalObject: null,

 
    initSmallGraph() {
      if (this.graphObject) return;
        const nodes = this.data.nodes.map((node) => ({ ...node }));
        const edges = this.data.edges.map((edge) => ({ ...edge }));
        
        this.graphObject = new ObsiGraph(
          graphCanvas,
          nodes,
          edges,
          {...graphConfig}
        );
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
        this.graphModalObject = new ObsiGraph(
          modalGraphCanvas,
          nodes,
          edges,
          {...graphConfig}
        );
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
    leftCollapsed: localStorage.getItem("leftSidebarCollapsed") === "true",
    rightCollapsed: localStorage.getItem("rightSidebarCollapsed") === "true",

    toggleLeft() {
      this.leftCollapsed = !this.leftCollapsed;
      localStorage.setItem("leftSidebarCollapsed", this.leftCollapsed);
    },

    toggleRight() {
      this.rightCollapsed = !this.rightCollapsed;
      localStorage.setItem("rightSidebarCollapsed", this.rightCollapsed);
    },
  });


  window.addEventListener("theme-initialized", (e) => {
    const isInitialized = e.detail;
    if (isInitialized.theme) {
      Alpine.store("graph").initSmallGraph();
    }
  });
});
