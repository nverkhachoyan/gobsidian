// Live reload script with exponential backoff reconnection and status notifications
export function initLiveReload() {
  if (
    window.location.hostname !== "localhost" &&
    window.location.hostname !== "127.0.0.1"
  ) {
    return;
  }

  let ws;
  let reconnectAttempts = 0;
  const maxReconnectAttempts = 10;
  const baseReconnectInterval = 1000;
  let reconnectTimeout;

  function getReconnectDelay() {
    return Math.min(baseReconnectInterval * Math.pow(2, reconnectAttempts), 10000);
  }

  function connect() {
    if (ws) {
      ws.close();
    }

    try {
      ws = new WebSocket(`ws://${window.location.host}/ws`);

      ws.onopen = () => {
        console.log("%cLive reload connected", "color: #4CAF50; font-weight: bold");
        reconnectAttempts = 0;
      };

      ws.onclose = () => {
        const delay = getReconnectDelay();
        if (reconnectAttempts < maxReconnectAttempts) {
          reconnectAttempts++;
          console.log(
            `%cLive reload disconnected. Reconnecting in ${delay/1000}s... (${reconnectAttempts}/${maxReconnectAttempts})`,
            "color: #FFA500"
          );
          reconnectTimeout = setTimeout(connect, delay);
        } else {
          console.log(
            "%cLive reload max reconnection attempts reached. Please restart the server.",
            "color: #FF0000; font-weight: bold"
          );
        }
      };

      ws.onerror = (error) => {
        console.error("Live reload connection error:", error);
        ws.close();
      };

      ws.onmessage = (event) => {
        if (event.data === "reload") {
          console.log("%cLive reload: Refreshing page...", "color: #2196F3");
          window.location.reload();
        }
      };

    } catch (err) {
      console.error("Live reload connection failed:", err);
      if (reconnectAttempts < maxReconnectAttempts) {
        reconnectTimeout = setTimeout(connect, getReconnectDelay());
      }
    }
  }

  connect();

  window.addEventListener('beforeunload', () => {
    clearTimeout(reconnectTimeout);
    if (ws) {
      ws.close();
    }
  });
}