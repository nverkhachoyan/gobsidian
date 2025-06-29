document.addEventListener("DOMContentLoaded", () => {
  const folderContainer = document.querySelector(".file-tree-container");
  const folderTitle = folderContainer?.querySelector(
    ".left-sidebar-item-title"
  );

  if (folderTitle) {
    folderTitle.addEventListener("click", () => {
      // Only toggle if in mobile view.
      if (window.innerWidth <= 768) {
        folderContainer.classList.toggle("expanded");
      }
    });
  }
});
