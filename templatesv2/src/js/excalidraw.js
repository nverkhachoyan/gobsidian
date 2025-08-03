// See https://www.npmjs.com/package/@excalidraw/excalidraw documentation.
import * as ExcalidrawLib from "https://esm.sh/@excalidraw/excalidraw@0.18.0/dist/dev/index.js?external=react,react-dom";
import React from "https://esm.sh/react@19.0.0";
import { createRoot } from "https://esm.sh/react-dom@19.0.0/client";
import ReactDOM from "https://esm.sh/react-dom@19.0.0";
import LZString from "https://esm.sh/lz-string@1.5.0";

window.ExcalidrawLib = ExcalidrawLib;

const App = ({ theme }) => {
  const filename = window.location.href.split("/").pop();
  const filenameJson = filename.replace(".html", ".json");
  const addressJson = `/generated/excalidraws/${filenameJson}`;
  const [initialSceneData, setInitialSceneData] = React.useState(null);
  const [isLoading, setIsLoading] = React.useState(true);
  const [error, setError] = React.useState(null);

  React.useEffect(() => {
    const loadData = async () => {
      try {
        const response = await fetch(addressJson);
        const jsonData = await response.json();
        const data = jsonData.data;
        const decompressedString = LZString.decompressFromBase64(data);
        if (!decompressedString) {
          throw new Error("Failed to decompress data");
        }

        let parsedData = JSON.parse(decompressedString);
        parsedData = {
          ...parsedData,
          appState: {
            ...parsedData.appState,
            theme: theme,
          },
        };

        console.log("parsedData", parsedData);

        setInitialSceneData(parsedData);
      } catch (err) {
        console.error("Failed to load/decode Excalidraw data:", err);
        setError(err);
      } finally {
        setIsLoading(false);
      }
    };

    loadData();
  }, []);

  if (isLoading) {
    return React.createElement("div", null, "Loading...");
  }

  if (error || !initialSceneData) {
    return React.createElement("div", null, "Error loading drawing data.");
  }

  return React.createElement(
    "div",
    {
      className: "excalidraw-wrapper",
      style: { height: "600px" },
    },
    React.createElement(ExcalidrawLib.Excalidraw, {
      initialData: initialSceneData,
    })
  );
};

export function renderExcalidraw(theme) {
  const excalidrawWrapper = document.getElementById("excalidraw");
  const root = createRoot(excalidrawWrapper);
  root.render(React.createElement(App, { theme: theme }));
}
