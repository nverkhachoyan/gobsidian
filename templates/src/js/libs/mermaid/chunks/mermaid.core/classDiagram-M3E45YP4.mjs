import {
  ClassDB,
  classDiagram_default,
  classRenderer_v3_unified_default,
  styles_default
} from "./chunk-SZ463SBG.mjs";
import "./chunk-E2GYISFI.mjs";
import "./chunk-BFAMUDN2.mjs";
import "./chunk-SKB7J2MH.mjs";
import "./chunk-IWUHOULB.mjs";
import "./chunk-M6DAPIYF.mjs";
import "./chunk-MXNHSMXR.mjs";
import "./chunk-JW4RIYDF.mjs";
import "./chunk-AC5SNWB5.mjs";
import "./chunk-UWXLY5YG.mjs";
import "./chunk-QESNASVV.mjs";
import "./chunk-55PJQP7W.mjs";
import {
  __name
} from "./chunk-3XYRH5AP.mjs";

// src/diagrams/class/classDiagram.ts
var diagram = {
  parser: classDiagram_default,
  get db() {
    return new ClassDB();
  },
  renderer: classRenderer_v3_unified_default,
  styles: styles_default,
  init: /* @__PURE__ */ __name((cnf) => {
    if (!cnf.class) {
      cnf.class = {};
    }
    cnf.class.arrowMarkerAbsolute = cnf.arrowMarkerAbsolute;
  }, "init")
};
export {
  diagram
};
