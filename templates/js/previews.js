const internalLinks = document.querySelectorAll('a[href^="/"]');
let previewCard = null;
let showTimeout;
let hideTimeout;

const hidePreview = (useTimeout = true) => {
  clearTimeout(showTimeout);
  hideTimeout = setTimeout(
    () => {
      if (previewCard) {
        previewCard.remove();
        previewCard = null;
      }
    },
    useTimeout ? 300 : 0
  );
};

const showPreview = async (e, link) => {
  clearTimeout(hideTimeout);

  const href = link.getAttribute("href");

  showTimeout = setTimeout(async () => {
    if (previewCard) {
      previewCard.remove();
    }

    previewCard = document.createElement("div");
    previewCard.classList.add("link-preview");
    document.body.appendChild(previewCard);

    // When the mouse enters the card, cancel the hide timer
    previewCard.addEventListener("mouseenter", () => clearTimeout(hideTimeout));
    // When the mouse leaves the card, hide it
    previewCard.addEventListener("mouseleave", hidePreview);

    // Construct preview URL
    const previewUrl = `/previews${href}`;

    previewCard.innerHTML = "Loading...";
    updateCardPosition(e);

    try {
      const response = await fetch(previewUrl);
      if (!response.ok) {
        previewCard.innerHTML = "";
        previewCard.classList.add("broken-link");

        const newChildElement = document.createElement("p");
        newChildElement.textContent = "Oh no! This link is broken.";
        newChildElement.classList.add("broken-link-text");
        previewCard.appendChild(newChildElement);
        updateCardPosition(e);
        return;
      }
      const content = await response.text();
      previewCard.innerHTML = content;
      updateCardPosition(e);
      const closeBtn = previewCard.querySelector(".close-preview-btn");
      if (closeBtn) {
        closeBtn.addEventListener("click", () => hidePreview(false));
      }
    } catch (err) {
      console.error("Failed to load preview", err);
      previewCard.innerHTML = "";
      previewCard.classList.add("broken-link");

      const newChildElement = document.createElement("p");
      newChildElement.textContent = "Oh no! This link is broken.";
      newChildElement.classList.add("broken-link-text");
      previewCard.appendChild(newChildElement);
      updateCardPosition(e);
    }
  }, 500);
};

internalLinks.forEach((link) => {
  // We only want previews for post pages
  if (
    !link.href.endsWith(".html") ||
    link.closest(".post-meta") ||
    link.closest("h1")
  ) {
    return;
  }

  link.addEventListener("mouseenter", (e) => showPreview(e, link));
  link.addEventListener("mouseleave", hidePreview);
});

function updateCardPosition(e) {
  if (!previewCard) return;

  const cardWidth = 300;
  previewCard.style.width = `${cardWidth}px`;
  previewCard.style.position = "absolute";
  previewCard.style.top = `${e.pageY + 15}px`;

  let leftPosition = e.pageX + 15;
  // if it overflows the viewport, position it to the left of the cursor
  if (leftPosition + cardWidth > window.innerWidth) {
    leftPosition = e.pageX - cardWidth - 15;
  }
  previewCard.style.left = `${leftPosition}px`;
}
