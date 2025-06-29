// Simple live reload script
(() => {
  if (
    window.location.hostname === "localhost" ||
    window.location.hostname === "127.0.0.1"
  ) {
    const conn = new WebSocket("ws://" + window.location.host + "/ws");
    conn.onclose = (evt) => {
      console.log("Connection to live reload server closed.");
    };
    conn.onmessage = (evt) => {
      if (evt.data === "reload") {
        console.log("Reloading page...");
        location.reload();
      }
    };
  }
})();
