// Simple live reload script with reconnect logic

if (
  window.location.hostname === "localhost" ||
  window.location.hostname === "127.0.0.1"
) {
  let ws;
  const reconnectInterval = 1000;

  function connect() {
    ws = new WebSocket("ws://" + window.location.host + "/ws");

    ws.onopen = (event) => {
      console.log("Live reload: Connection established.");
    };

    ws.onclose = (event) => {
      console.log(
        "Live reload: Connection closed. Attempting to reconnect...",
        event.code,
        event.reason
      );
      setTimeout(connect, reconnectInterval);
    };

    ws.onerror = (error) => {
      console.error("Live reload: WebSocket error:", error);
    };

    ws.onmessage = (event) => {
      if (event.data === "reload") {
        console.log("Live reload: Reloading page...");
        location.reload();
      }
    };
  }

  connect();
}
