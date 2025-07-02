
document.addEventListener('DOMContentLoaded', function() {
  const zoomGraphButton = document.querySelector('.zoom-graph-button');
  const modal = document.getElementById('graph-modal');
  const modalClose = document.getElementById('graph-modal-close');

  function getCssVariable(variableName) {
      return getComputedStyle(document.documentElement).getPropertyValue(variableName).trim();
  }

  const isDarkMode = document.documentElement.classList.contains('dark-mode');

  const nodeBgColor = '#aaaab3'
  const nodeBorderColor = getCssVariable('--border-color');
  const nodeFontColor = getCssVariable('--text-color');
  const edgeColor = getCssVariable('--subtle-text-color');
  const edgeHighlightColor = isDarkMode ? '#8a5cf5' : '#8a5cf5'; 
  const edgeHoverColor = isDarkMode ? '#a0a0a0' : '#8a5cf5';

 
  var rawData = window.graphData;
  console.log(rawData);
 
  var nodes = new vis.DataSet(rawData.nodes);
  var edges = new vis.DataSet(rawData.edges);
  var container = document.getElementById('graph-network');
  var data = {
      nodes: nodes,
      edges: edges
  };
  
  var options = {
    height: '100%',
    width: '100%',
    autoResize: true,
    nodes: {
      shape: 'dot',
      size: 10,
      font: {
        size: 14,
        color: nodeFontColor, 
        multi: true,
        align: 'center',
        strokeWidth: 0,
        vadjust: 0
      },
      widthConstraint: {
        maximum: 120
      },
      color: {
        background: nodeBgColor, 
        border: nodeBorderColor,     
        highlight: {
          background: isDarkMode ? '#3a3a3a' : '#e0e0e0',
          border: isDarkMode ? '#5a5a5a' : '#c0c0c0',
        },
        hover: {
          background: isDarkMode ? '#2a2a2a' : '#f0f0f0',
          border: isDarkMode ? '#4a4a4a' : '#d0d0d0'
        }
      },
      borderWidth: 1,
      shadow: {
        enabled: true,
        color: getCssVariable('--shadow-color'),  
        size: 8,
        x: 0,
        y: 0
      }
    },
    edges: {
      color: {
        color: edgeColor, 
        highlight: edgeHighlightColor,  
        hover: edgeHoverColor,  
        inherit: 'from',
        opacity: 0.8,
    
      },
      arrows: {
        to: {
          enabled: false,
          scaleFactor: 0.5
        }
      },
      smooth: {
        enabled: true,
        type: 'continuous',
        roundness: 1
      },
      width: 0.5,
      shadow: {
        enabled: false
      }
    },
    physics: {
      enabled: true,
      barnesHut: {
        gravitationalConstant: -8000,
        centralGravity: 0.3,
        springLength: 150,
        springConstant: 0.05,
        damping: 0.09,
        avoidOverlap: 0.8
      },
      maxVelocity: 50,
      minVelocity: 0.75,
      solver: 'barnesHut',
      stabilization: {
        enabled: true,
        iterations: 2000,
        updateInterval: 25
      }
    },
    interaction: {
      hover: true,
      zoomView: true,
      dragView: true,
      tooltipDelay: 300
    },
    layout: {
      randomSeed: undefined,
      improvedLayout: true
    }
  };

  
  var network = new vis.Network(container, data, options);
  
  // Add hover cursor functionality for main graph
  network.on("hoverNode", function (params) {
    container.style.cursor = 'pointer';
  });
  
  network.on("blurNode", function (params) {
    container.style.cursor = 'default';
  });

  // Add click functionality for main graph
  network.on("click", function (params) {
    if (params.nodes.length > 0) {
      const nodeId = params.nodes[0];
      const nodeData = nodes.get(nodeId);
      if (nodeData && nodeData.url) {
        window.location.href = nodeData.url;
      }
    }
  });

  
  var modalNetwork = null;

  // Function to open modal
  function openModal() {
    modal.style.display = 'block';
    document.body.style.overflow = 'hidden'; // Prevent background scrolling
    
    // Create modal graph instance if it doesn't exist
    if (!modalNetwork) {
      const modalContainer = document.getElementById('graph-network-modal');
      
      // Enhanced options for modal graph (larger, more interactive)
      const modalOptions = {
        ...options,
        nodes: {
          ...options.nodes,
          size: 15, // Larger nodes in modal
          font: {
            ...options.nodes.font,
            size: 16 // Larger font in modal
          }
        },
        physics: {
          ...options.physics,
          stabilization: {
            enabled: true,
            iterations: 1000,
            updateInterval: 25
          }
        }
      };
      
      modalNetwork = new vis.Network(modalContainer, data, modalOptions);
      
      // Add hover cursor functionality for modal graph
      modalNetwork.on("hoverNode", function (params) {
        modalContainer.style.cursor = 'pointer';
      });
      
      modalNetwork.on("blurNode", function (params) {
        modalContainer.style.cursor = 'default';
      });

      // Add click functionality for modal graph
      modalNetwork.on("click", function (params) {
        if (params.nodes.length > 0) {
          const nodeId = params.nodes[0];
          const nodeData = nodes.get(nodeId);
          if (nodeData && nodeData.url) {
            window.location.href = nodeData.url;
          }
        }
      });
      
      // Fit the modal graph after a short delay to ensure proper rendering
      setTimeout(() => {
        modalNetwork.fit({
          animation: true,
          duration: 1000,
          easingFunction: 'easeInOutQuad'
        });
      }, 100);
    } else {
      // If network already exists, just fit it
      setTimeout(() => {
        modalNetwork.redraw();
        modalNetwork.fit({
          animation: true,
          duration: 1000,
          easingFunction: 'easeInOutQuad'
        });
      }, 100);
    }
  }

  // Function to close modal
  function closeModal() {
    modal.style.display = 'none';
    document.body.style.overflow = 'auto'; // Restore background scrolling
  }

  // Event listeners
  if (zoomGraphButton) {
    zoomGraphButton.addEventListener('click', function() {
      openModal();
    });
  }

  if (modalClose) {
    modalClose.addEventListener('click', function() {
      closeModal();
    });
  }

  // Close modal when clicking outside the modal content
  if (modal) {
    modal.addEventListener('click', function(event) {
      if (event.target === modal) {
        closeModal();
      }
    });
  }

  // Close modal with Escape key
  document.addEventListener('keydown', function(event) {
    if (event.key === 'Escape' && modal && modal.style.display === 'block') {
      closeModal();
    }
  });

  // Handle window resize for modal
  window.addEventListener('resize', function() {
    if (modalNetwork && modal && modal.style.display === 'block') {
      modalNetwork.redraw();
    }
  });
});