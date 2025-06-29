document.addEventListener("DOMContentLoaded", () => {
  document.querySelectorAll(".file-tree__toggle").forEach((toggle) => {
    toggle.addEventListener("click", () => {
      const folderItem = toggle.closest(".file-tree__item--folder");
      if (folderItem) {
        folderItem.classList.toggle("expanded");
      }
    });
  });
});
