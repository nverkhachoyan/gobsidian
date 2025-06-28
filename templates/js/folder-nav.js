document.addEventListener("DOMContentLoaded", () => {
  document
    .querySelectorAll(".file-tree-nav .folder-toggle")
    .forEach((toggle) => {
      toggle.addEventListener("click", () => {
        const folderItem = toggle.closest(".folder-item");
        if (folderItem) {
          folderItem.classList.toggle("expanded");
        }
      });
    });
});
