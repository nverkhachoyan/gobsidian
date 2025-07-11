// Global Alpine.js Stores
const graphContainer = document.getElementById('graph-network');
const modalGraphContainer = document.getElementById('graph-network-modal');


const initGraph = (container, nodes, edges, config) => {
    if (container) {
        new D3Graph(
            container,
            nodes,
            edges,
            config,
        );
    }
}


document.addEventListener('alpine:init', () => {
  Alpine.store('theme', {
    current: localStorage.getItem('theme') || (window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'),
    
    init() {
      this.setTheme(this.current);
    },
    
    toggle() {
      const newTheme = this.current === 'light' ? 'dark' : 'light';
      this.setTheme(newTheme);
    },
    
    setTheme(newTheme) {
      this.current = newTheme;
      localStorage.setItem('theme', newTheme);
      
      // Update DOM
      if (newTheme === 'dark') {
        document.documentElement.classList.add('dark-mode');
      } else {
        document.documentElement.classList.remove('dark-mode');
      }
      
      // Trigger custom event for other components
      window.dispatchEvent(new CustomEvent('theme-changed', {
        detail: { theme: newTheme }
      }));
    }
  });

  Alpine.store('fileTree', {
    data: window.fileTreeData || null,
    expanded: localStorage.getItem('fileTreeExpanded') === 'true',
    
    toggle() {
      this.expanded = !this.expanded;
      localStorage.setItem('fileTreeExpanded', this.expanded);
    },
    
    setExpanded(state) {
      this.expanded = state;
      localStorage.setItem('fileTreeExpanded', state);
    }
  });

  Alpine.store('tags', {
    data: window.tagsData || [],
    showAll: localStorage.getItem('tagsShowAll') === 'true',
    
    toggle() {
      this.showAll = !this.showAll;
      localStorage.setItem('tagsShowAll', this.showAll);
    },
    
    getVisible() {
      return this.showAll ? this.data : this.data.slice(0, 6);
    },
    
    hasMore() {
      return this.data.length > 6;
    }
  });

  Alpine.store('graph', {
    data: window.graphData || null,
    modalOpen: false,
    modalHasBeenInitialized: false,
    config: {
      initialZoomFactor: 0.7,
      nodesHaveLinks: true,
      colors: {
        node: '#b2b2b2',
        nodeStroke: '#bca1f2',
        nodeFont: Alpine.store('theme').current === 'dark' ? '#e8f4f8' : '#3f3f3f',
        edge: '#626262',
        nodeHighlight: '#9165ea',
        edgeHighlight: '#9165ea',
        background: Alpine.store('theme').current === 'dark' ? '#0d1117' : '#ffffff',
        grid: Alpine.store('theme').current === 'dark' ? 'rgba(26, 26, 26, 0.1)' : 'rgba(240, 240, 240, 0.1)',
        globeBoundary: Alpine.store('theme').current === 'dark' ? 'rgba(76, 159, 206, 0.1)' : 'rgba(33, 150, 243, 0.1)',
      },
      sizing: {
        minNodeSize: 4,
        maxNodeSize: 20,
        globeRadiusMultiplier: 0.18,
      },
      simulation: {
        linkDistance: 100,
        linkStrength: 1,
        chargeStrength: -300,
        collisionStrength: 0.2,
        globeStrength: 0.3,
        alpha: 0.5,
        finalAlpha: 0.005,
        simulationTime: 10000,
      },
      interaction: {
        dragThreshold: 5, // pixels
        zoomMin: 0.1,
        zoomMax: 5,
      },
      font: 'sans-serif',
    },
    
    openModal() {
      this.modalOpen = true;
      if (!this.modalHasBeenInitialized) {
       setTimeout(() => {
          initGraph(modalGraphContainer, this.data.nodes, this.data.edges, this.config);
          this.modalHasBeenInitialized = true;
       }, 200);
      }
    },
    
    closeModal() {
      this.modalOpen = false;
    },
    
    toggleModal() {
      this.modalOpen = !this.modalOpen;
    }
  });

  Alpine.store('sidebar', {
    leftCollapsed: localStorage.getItem('leftSidebarCollapsed') === 'true',
    rightCollapsed: localStorage.getItem('rightSidebarCollapsed') === 'true',
    
    toggleLeft() {
      this.leftCollapsed = !this.leftCollapsed;
      localStorage.setItem('leftSidebarCollapsed', this.leftCollapsed);
    },
    
    toggleRight() {
      this.rightCollapsed = !this.rightCollapsed;
      localStorage.setItem('rightSidebarCollapsed', this.rightCollapsed);
    }
  });


  // Initialize the graph
  const graph = Alpine.store("graph");
  const nodes = graph.data.nodes;
  const edges = graph.data.edges;
  initGraph(graphContainer, nodes, edges, graph.config);
}); 



