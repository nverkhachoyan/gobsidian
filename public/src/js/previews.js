let previewStack = [];
let showTimeout;
let hideTimeout;

const hideCardsFromIndex = (startIndex, useTimeout = true, force = false) => {
  clearTimeout(showTimeout);

  const performHide = () => {
    // Before hiding, check if the mouse is over any card that is a candidate for removal.
    // If so, it means the user moved into a child card, so we shouldn't hide.
    if (!force) {
      for (let i = startIndex; i < previewStack.length; i++) {
        if (previewStack[i].matches(":hover")) {
          return; // Abort hide
        }
      }
    }

    // Iterate backwards to safely remove items from the stack
    for (let i = previewStack.length - 1; i >= startIndex; i--) {
      const card = previewStack[i];
      if (force || card.dataset.pinned !== "true") {
        card.remove();
        previewStack.splice(i, 1);
      }
    }
  };

  hideTimeout = setTimeout(performHide, useTimeout ? 500 : 0);
};

const showPreview = async (e, link) => {
  clearTimeout(hideTimeout);

  const href = link.getAttribute("href");

  showTimeout = setTimeout(async () => {
    const previewCard = document.createElement("div");
    previewCard.classList.add("link-preview");
    document.body.appendChild(previewCard);

    // When the mouse enters the card, cancel the hide timer
    previewCard.addEventListener("mouseenter", () => clearTimeout(hideTimeout));
    // When the mouse leaves the card, hide it and any "child" previews
    previewCard.addEventListener("mouseleave", () => {
      const cardIndex = previewStack.indexOf(previewCard);
      if (cardIndex !== -1) {
        hideCardsFromIndex(cardIndex, true);
      }
    });

    previewStack.push(previewCard);

    // Construct preview URL
    const previewUrl = `/previews${href}`;

    previewCard.innerHTML = "Loading...";
    updateCardPosition(e, previewCard);

    try {
      const response = await fetch(previewUrl);
      if (!response.ok) {
        previewCard.innerHTML = "";
        previewCard.classList.add("broken-link");

        const newChildElement = document.createElement("p");
        newChildElement.textContent = "Oh no! This link is broken.";
        newChildElement.classList.add("broken-link-text");
        previewCard.appendChild(newChildElement);
        updateCardPosition(e, previewCard);
        return;
      }
      const content = await response.text();
      previewCard.innerHTML = content;
      Alpine.initTree(previewCard);

      const dragHandle = previewCard.querySelector(".drag-handle");
      if (dragHandle) {
        makeDraggable(previewCard, dragHandle);
      }

      const resizeHandle = previewCard.querySelector(".resize-handle");
      if (resizeHandle) {
        makeResizable(previewCard, resizeHandle);
      }

      updateCardPosition(e, previewCard);
      const cardIndex = previewStack.indexOf(previewCard);
      initPreviews(previewCard, cardIndex); // Recursively enable previews

      const closeBtn = previewCard.querySelector(".close-preview-btn");
      if (closeBtn) {
        closeBtn.addEventListener("click", () => {
          const cardIndex = previewStack.indexOf(previewCard);
          if (cardIndex !== -1) {
            hideCardsFromIndex(cardIndex, false, true); // Force close
          }
        });
      }
      const pinBtn = previewCard.querySelector(".pin-preview-btn");
      if (pinBtn) {
        pinBtn.addEventListener("click", () => {
          const isPinned = previewCard.dataset.pinned === "true";
          previewCard.dataset.pinned = isPinned ? "false" : "true";
          pinBtn.classList.toggle("active", !isPinned);
        });
      }
    } catch (err) {
      console.error("Failed to load preview", err);
      previewCard.innerHTML = "";
      previewCard.classList.add("broken-link");

      const newChildElement = document.createElement("p");
      newChildElement.textContent = "Oh no! This link is broken.";
      newChildElement.classList.add("broken-link-text");
      previewCard.appendChild(newChildElement);
      updateCardPosition(e, previewCard);
    }
  }, 300);
};

const initPreviews = (scope, parentCardIndex = -1) => {
  const internalLinks = scope.querySelectorAll('a[href^="/"]');
  internalLinks.forEach((link) => {
    if (link.dataset.previewInitialized) {
      return;
    }
    if (
      !link.href.endsWith(".html") ||
      link.closest(".post-meta") ||
      link.closest("h1")
    ) {
      return;
    }
    link.dataset.previewInitialized = "true";

    link.addEventListener("mouseenter", (e) => showPreview(e, link));
    const indexToHideFrom = parentCardIndex + 1;
    link.addEventListener("mouseleave", () =>
      hideCardsFromIndex(indexToHideFrom, true)
    );
  });
};

function updateCardPosition(e, card) {
  if (!card) return;

  const cardWidth = 300;
  card.style.width = `${cardWidth}px`;
  card.style.position = "absolute";

  const level = Math.max(0, previewStack.indexOf(card));
  const offset = level * 15;
  card.style.zIndex = 100 + level;

  card.style.top = `${e.pageY + offset}px`;

  let leftPosition = e.pageX + offset;
  // if it overflows the viewport, position it to the left of the cursor
  if (leftPosition + cardWidth > window.innerWidth) {
    leftPosition = e.pageX - cardWidth - offset;
  }
  card.style.left = `${leftPosition}px`;
}

function makeResizable(element, handle) {
  let isResizing = false;
  let startX, startY, startWidth, startHeight;

  const onMouseDown = (e) => {
    e.preventDefault();
    isResizing = true;
    startX = e.clientX;
    startY = e.clientY;
    const rect = element.getBoundingClientRect();
    startWidth = rect.width;
    startHeight = rect.height;

    document.addEventListener("mousemove", onMouseMove);
    document.addEventListener("mouseup", onMouseUp);
  };

  const onMouseMove = (e) => {
    if (!isResizing) return;
    const newWidth = startWidth + (e.clientX - startX);
    const newHeight = startHeight + (e.clientY - startY);
    element.style.width = `${newWidth}px`;
    element.style.height = `${newHeight}px`;
  };

  const onMouseUp = () => {
    isResizing = false;
    document.removeEventListener("mousemove", onMouseMove);
    document.removeEventListener("mouseup", onMouseUp);
  };

  handle.addEventListener("mousedown", onMouseDown);
}

function makeDraggable(element, handle) {
  let isDragging = false;
  let offsetX, offsetY;

  const onMouseDown = (e) => {
    isDragging = true;

    // Bring card to front when dragging starts
    const allPreviews = document.querySelectorAll(".link-preview");
    const maxZ = Math.max(
      ...Array.from(allPreviews).map(
        (el) => parseFloat(el.style.zIndex) || 100
      )
    );
    element.style.zIndex = maxZ + 1;

    const initialLeft = parseInt(element.style.left || 0, 10);
    const initialTop = parseInt(element.style.top || 0, 10);

    offsetX = e.clientX - initialLeft;
    offsetY = e.clientY - initialTop;

    document.addEventListener("mousemove", onMouseMove);
    document.addEventListener("mouseup", onMouseUp);

    // Prevent text selection
    e.preventDefault();
  };

  const onMouseMove = (e) => {
    if (!isDragging) return;
    element.style.left = `${e.clientX - offsetX}px`;
    element.style.top = `${e.clientY - offsetY}px`;
  };

  const onMouseUp = () => {
    isDragging = false;
    document.removeEventListener("mousemove", onMouseMove);
    document.removeEventListener("mouseup", onMouseUp);
  };

  handle.addEventListener("mousedown", onMouseDown);
}

initPreviews(document);
