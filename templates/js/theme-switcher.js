document.addEventListener("DOMContentLoaded", () => {
  const themeSwitcher = document.getElementById("theme-switcher");
  const html = document.documentElement;

  themeSwitcher.addEventListener("click", () => {
    html.classList.toggle("dark-mode");
    if (html.classList.contains("dark-mode")) {
      localStorage.setItem("theme", "dark-mode");
    } else {
      localStorage.setItem("theme", "light-mode");
    }
  });
});
