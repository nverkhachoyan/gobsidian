let mermaidInitialized = false;
const originalMermaidContent = new Map();

function initializeMermaid(theme) {
  // Convert theme to Mermaid theme
  const mermaidTheme = theme === "dark" ? "dark" : "neutral";

  mermaid.initialize({
    startOnLoad: false, // We'll handle this manually
    theme: mermaidTheme,
  });

  // Find all Mermaid diagrams
  const diagrams = document.querySelectorAll(".mermaid");

  diagrams.forEach((diagram, index) => {
    // Store original content on first initialization
    if (!mermaidInitialized) {
      const originalContent = diagram.textContent.trim();
      originalMermaidContent.set(diagram, originalContent);
    }

    // For re-initialization, restore original content
    if (mermaidInitialized) {
      const originalContent = originalMermaidContent.get(diagram);
      if (originalContent) {
        diagram.innerHTML = originalContent;
        diagram.removeAttribute("data-processed");
      }
    }
  });

  // Render all diagrams
  if (diagrams.length > 0) {
    mermaid.run({
      nodes: diagrams,
    });
  }

  mermaidInitialized = true;
}

// Initialize Mermaid when Alpine is ready
document.addEventListener("alpine:init", () => {
  const currentTheme = Alpine.store("theme").current;
  initializeMermaid(currentTheme);
});

// Listen for theme changes and re-initialize Mermaid
window.addEventListener("theme-changed", (event) => {
  initializeMermaid(event.detail.theme);
});