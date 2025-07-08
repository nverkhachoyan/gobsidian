const zoomGraphButton = document.querySelector(".zoom-graph-button");
const modal = document.getElementById("graph-modal");
const modalClose = document.getElementById("graph-modal-close");

function getCssVariable(variableName) {
  return getComputedStyle(document.documentElement)
    .getPropertyValue(variableName)
    .trim();
}

const isDarkMode = document.documentElement.classList.contains("dark-mode");

const nodeBgColor = "#aaaab3";
const nodeBorderColor = getCssVariable("--border-color");
const nodeFontColor = isDarkMode ? "#e0e0e0" : "#212529";
const edgeColor = getCssVariable("--subtle-text-color");
const nodeHighlightColor = "#8a5cf5";
const edgeHighlightColor = "#8153ec";
const edgeHoverColor = isDarkMode ? "#a0a0a0" : "#8a5cf5";

const rawData = window.graphData;
console.log(rawData);

// Calculate degree (number of connections) for each node
const degreeMap = {};
rawData.nodes.forEach((node) => {
  degreeMap[node.id] = 0;
});

rawData.edges.forEach((edge) => {
  if (degreeMap[edge.from] !== undefined) degreeMap[edge.from]++;
  if (degreeMap[edge.to] !== undefined) degreeMap[edge.to]++;
});

// Find min and max degrees for scaling
const degrees = Object.values(degreeMap);
const minDegree = Math.min(...degrees);
const maxDegree = Math.max(...degrees);

// Set node sizes based on degree
const nodesWithSize = rawData.nodes.map((node) => {
  const degree = degreeMap[node.id];
  // Scale size between 8 and 25 based on degree
  const minSize = 8;
  const maxSize = 25;
  const size =
    maxDegree > minDegree
      ? minSize +
        ((degree - minDegree) * (maxSize - minSize)) / (maxDegree - minDegree)
      : minSize;

  return {
    ...node,
    size: Math.max(size, minSize), // Ensure minimum size
    title: node.title
      ? `${node.title}<br/>Connections: ${degree}`
      : `Connections: ${degree}`, // Add degree to tooltip
  };
});

const nodes = new vis.DataSet(nodesWithSize);
const edges = new vis.DataSet(rawData.edges);
const container = document.getElementById("graph-network");
const data = {
  nodes: nodes,
  edges: edges,
};

const options = {
  height: "100%",
  width: "100%",
  autoResize: true,
  nodes: {
    shape: "dot",
    font: {
      size: 14,
      color: nodeFontColor,
      multi: true,
      align: "center",
      strokeWidth: 0,
      vadjust: 0,
    },
    widthConstraint: {
      maximum: 120,
    },
    color: {
      background: nodeBgColor,
      border: nodeBorderColor,
      highlight: {
        background: nodeHighlightColor,
        border: isDarkMode ? "#5a5a5a" : "#c0c0c0",
      },
      hover: {
        background: nodeHighlightColor,
        border: isDarkMode ? "#4a4a4a" : "#d0d0d0",
      },
    },
    borderWidth: 1,
    shadow: {
      enabled: true,
      color: getCssVariable("--shadow-color"),
      size: 8,
      x: 0,
      y: 0,
    },
  },
  edges: {
    color: {
      color: edgeColor,
      highlight: edgeHighlightColor,
      hover: edgeHoverColor,
      inherit: "from",
      opacity: 0.8,
    },
    arrows: {
      to: {
        enabled: false,
        scaleFactor: 0.5,
      },
    },
    smooth: {
      enabled: true,
      type: "continuous",
      roundness: 1,
    },
    width: 0.5,
    shadow: {
      enabled: false,
    },
  },
  physics: {
    enabled: true,
    barnesHut: {
      gravitationalConstant: -8000,
      centralGravity: 0.3,
      springLength: 150,
      springConstant: 0.05,
      damping: 0.09,
      avoidOverlap: 0.8,
    },
    maxVelocity: 50,
    minVelocity: 0.75,
    solver: "barnesHut",
    stabilization: {
      enabled: true,
      iterations: 2000,
      updateInterval: 25,
    },
  },
  interaction: {
    hover: true,
    zoomView: true,
    dragView: true,
    tooltipDelay: 300,
  },
  layout: {
    randomSeed: undefined,
    improvedLayout: true,
  },
};

var network = new vis.Network(container, data, options);

network.on("hoverNode", function (params) {
  container.style.cursor = "pointer";
});

network.on("blurNode", function (params) {
  container.style.cursor = "default";
});

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

function openModal() {
  modal.style.display = "block";
  document.body.style.overflow = "hidden";

  if (!modalNetwork) {
    const modalContainer = document.getElementById("graph-network-modal");

    const modalOptions = {
      ...options,
      nodes: {
        ...options.nodes,
        font: {
          ...options.nodes.font,
          size: 16,
        },
      },
      physics: {
        ...options.physics,
        stabilization: {
          enabled: true,
          iterations: 1000,
          updateInterval: 25,
        },
      },
    };

    modalNetwork = new vis.Network(modalContainer, data, modalOptions);

    modalNetwork.on("hoverNode", function (params) {
      modalContainer.style.cursor = "pointer";
    });

    modalNetwork.on("blurNode", function (params) {
      modalContainer.style.cursor = "default";
    });

    modalNetwork.on("click", function (params) {
      if (params.nodes.length > 0) {
        const nodeId = params.nodes[0];
        const nodeData = nodes.get(nodeId);
        if (nodeData && nodeData.url) {
          window.location.href = nodeData.url;
        }
      }
    });

    setTimeout(() => {
      modalNetwork.fit({
        animation: true,
        duration: 1000,
        easingFunction: "easeInOutQuad",
      });
    }, 100);
  } else {
    setTimeout(() => {
      modalNetwork.redraw();
      modalNetwork.fit({
        animation: true,
        duration: 1000,
        easingFunction: "easeInOutQuad",
      });
    }, 100);
  }
}

function closeModal() {
  modal.style.display = "none";
  document.body.style.overflow = "auto";
}

if (zoomGraphButton) {
  zoomGraphButton.addEventListener("click", function () {
    openModal();
  });
}

if (modalClose) {
  modalClose.addEventListener("click", function () {
    closeModal();
  });
}

if (modal) {
  modal.addEventListener("click", function (event) {
    if (event.target === modal) {
      closeModal();
    }
  });
}

document.addEventListener("keydown", function (event) {
  if (event.key === "Escape" && modal && modal.style.display === "block") {
    closeModal();
  }
});

window.addEventListener("resize", function () {
  if (modalNetwork && modal && modal.style.display === "block") {
    modalNetwork.redraw();
  }
});
