import mermaid from './libs/mermaid/mermaid.esm.min.mjs'

mermaid.initialize({ 
  startOnLoad: false
});

export async function initializeMermaid(theme) {
  const mermaidTheme = theme === "dark" ? "dark" : "neutral";

  mermaid.initialize({ 
    startOnLoad: false,
    theme: mermaidTheme 
  });

  const diagrams = document.querySelectorAll('.mermaid');
  for (const diagram of diagrams) {
    // Storing original content before rendering
    // Since it breaks if theme is changed after render
    if (!diagram.dataset.originalContent) {
      diagram.dataset.originalContent = diagram.textContent.trim();
    }
    
    try {
      const id = 'mermaid-' + Math.random().toString(36).substring(2);
      const { svg } = await mermaid.render(id, diagram.dataset.originalContent);
      diagram.innerHTML = svg;
    } catch (err) {
      console.error('Failed to render mermaid diagram:', err);
      diagram.innerHTML = `<pre>Failed to render diagram: ${err.message}</pre>`;
    }
  }
}