// AlpineJS Core
import { initLiveReload } from "./live-reload.js";
import { initPreviews } from "./previews.js";
import { initializeMermaid } from "./mermaid.js";
import "./libs/alpine/alpine.min.js";
import "./libs/alpine/alpine-collapse.min.js";

import { initializeStores } from "./store.js";

initializeStores(window.Alpine);

window.onload = () => {
  const currentTheme = window.Alpine.store("theme").current;
  initializeMermaid(currentTheme);
  window.addEventListener("theme-changed", (e) => {
    const theme = e.detail.theme;
    if (theme) {
      initializeMermaid(theme);
    }
  });
};

initPreviews(document);

// KaTeX
import "./libs/katex/katex.min.js";
import "./libs/katex/contrib/auto-render.min.js";
import "./libs/katex/contrib/copy-tex.min.js";
import "./libs/katex/contrib/render-a11y-string.min.js";

document.addEventListener("DOMContentLoaded", function () {
  if (typeof renderMathInElement === "function") {
    renderMathInElement(document.body, {
      delimiters: [
        { left: "$$", right: "$$", display: true },
        { left: "$", right: "$", display: false },
        { left: "\\[", right: "\\]", display: true },
        { left: "\\(", right: "\\)", display: false },
      ],
    });
  }
});

initLiveReload();
