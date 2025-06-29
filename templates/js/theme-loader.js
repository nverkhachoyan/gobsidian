(() => {
  const savedTheme = localStorage.getItem("theme");
  if (savedTheme === "dark-mode") {
    document.documentElement.classList.add("dark-mode");
  } else if (savedTheme === "light-mode") {
    // Do nothing, light is the default
  } else if (window.matchMedia("(prefers-color-scheme: dark)").matches) {
    document.documentElement.classList.add("dark-mode");
  }
})();
