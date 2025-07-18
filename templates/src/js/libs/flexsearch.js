/**!
 * FlexSearch.js v0.8.200 (Bundle/Module/Debug)
 * Author and Copyright: Thomas Wilkerling
 * Licence: Apache-2.0
 * Hosted by Nextapps GmbH
 * https://github.com/nextapps-de/flexsearch
 */
var w;
function H(a, c, b) {
  const e = typeof b, d = typeof a;
  if ("undefined" !== e) {
    if ("undefined" !== d) {
      if (b) {
        if ("function" === d && e === d) {
          return function(h) {
            return a(b(h));
          };
        }
        c = a.constructor;
        if (c === b.constructor) {
          if (c === Array) {
            return b.concat(a);
          }
          if (c === Map) {
            var f = new Map(b);
            for (var g of a) {
              f.set(g[0], g[1]);
            }
            return f;
          }
          if (c === Set) {
            g = new Set(b);
            for (f of a.values()) {
              g.add(f);
            }
            return g;
          }
        }
      }
      return a;
    }
    return b;
  }
  return "undefined" === d ? c : a;
}
function aa(a, c) {
  return "undefined" === typeof a ? c : a;
}
function I() {
  return Object.create(null);
}
function N(a) {
  return "string" === typeof a;
}
function ba(a) {
  return "object" === typeof a;
}
function ca(a) {
  const c = [];
  for (const b of a.keys()) {
    c.push(b);
  }
  return c;
}
function da(a, c) {
  if (N(c)) {
    a = a[c];
  } else {
    for (let b = 0; a && b < c.length; b++) {
      a = a[c[b]];
    }
  }
  return a;
}
;const fa = /[^\p{L}\p{N}]+/u, ha = /(\d{3})/g, ia = /(\D)(\d{3})/g, ja = /(\d{3})(\D)/g, ka = /[\u0300-\u036f]/g;
function la(a = {}) {
  if (!this || this.constructor !== la) {
    return new la(...arguments);
  }
  if (arguments.length) {
    for (a = 0; a < arguments.length; a++) {
      this.assign(arguments[a]);
    }
  } else {
    this.assign(a);
  }
}
w = la.prototype;
w.assign = function(a) {
  this.normalize = H(a.normalize, !0, this.normalize);
  let c = a.include, b = c || a.exclude || a.split, e;
  if (b || "" === b) {
    if ("object" === typeof b && b.constructor !== RegExp) {
      let d = "";
      e = !c;
      c || (d += "\\p{Z}");
      b.letter && (d += "\\p{L}");
      b.number && (d += "\\p{N}", e = !!c);
      b.symbol && (d += "\\p{S}");
      b.punctuation && (d += "\\p{P}");
      b.control && (d += "\\p{C}");
      if (b = b.char) {
        d += "object" === typeof b ? b.join("") : b;
      }
      try {
        this.split = new RegExp("[" + (c ? "^" : "") + d + "]+", "u");
      } catch (f) {
        console.error("Your split configuration:", b, "is not supported on this platform. It falls back to using simple whitespace splitter instead: /s+/."), this.split = /\s+/;
      }
    } else {
      this.split = b, e = !1 === b || 2 > "a1a".split(b).length;
    }
    this.numeric = H(a.numeric, e);
  } else {
    try {
      this.split = H(this.split, fa);
    } catch (d) {
      console.warn("This platform does not support unicode regex. It falls back to using simple whitespace splitter instead: /s+/."), this.split = /\s+/;
    }
    this.numeric = H(a.numeric, H(this.numeric, !0));
  }
  this.prepare = H(a.prepare, null, this.prepare);
  this.finalize = H(a.finalize, null, this.finalize);
  b = a.filter;
  this.filter = "function" === typeof b ? b : H(b && new Set(b), null, this.filter);
  this.dedupe = H(a.dedupe, !0, this.dedupe);
  this.matcher = H((b = a.matcher) && new Map(b), null, this.matcher);
  this.mapper = H((b = a.mapper) && new Map(b), null, this.mapper);
  this.stemmer = H((b = a.stemmer) && new Map(b), null, this.stemmer);
  this.replacer = H(a.replacer, null, this.replacer);
  this.minlength = H(a.minlength, 1, this.minlength);
  this.maxlength = H(a.maxlength, 1024, this.maxlength);
  this.rtl = H(a.rtl, !1, this.rtl);
  if (this.cache = b = H(a.cache, !0, this.cache)) {
    this.I = null, this.R = "number" === typeof b ? b : 2e5, this.B = new Map(), this.H = new Map(), this.M = this.L = 128;
  }
  this.h = "";
  this.N = null;
  this.A = "";
  this.O = null;
  if (this.matcher) {
    for (const d of this.matcher.keys()) {
      this.h += (this.h ? "|" : "") + d;
    }
  }
  if (this.stemmer) {
    for (const d of this.stemmer.keys()) {
      this.A += (this.A ? "|" : "") + d;
    }
  }
  return this;
};
w.addStemmer = function(a, c) {
  this.stemmer || (this.stemmer = new Map());
  this.stemmer.set(a, c);
  this.A += (this.A ? "|" : "") + a;
  this.O = null;
  this.cache && Q(this);
  return this;
};
w.addFilter = function(a) {
  "function" === typeof a ? this.filter = a : (this.filter || (this.filter = new Set()), this.filter.add(a));
  this.cache && Q(this);
  return this;
};
w.addMapper = function(a, c) {
  if ("object" === typeof a) {
    return this.addReplacer(a, c);
  }
  if (1 < a.length) {
    return this.addMatcher(a, c);
  }
  this.mapper || (this.mapper = new Map());
  this.mapper.set(a, c);
  this.cache && Q(this);
  return this;
};
w.addMatcher = function(a, c) {
  if ("object" === typeof a) {
    return this.addReplacer(a, c);
  }
  if (2 > a.length && (this.dedupe || this.mapper)) {
    return this.addMapper(a, c);
  }
  this.matcher || (this.matcher = new Map());
  this.matcher.set(a, c);
  this.h += (this.h ? "|" : "") + a;
  this.N = null;
  this.cache && Q(this);
  return this;
};
w.addReplacer = function(a, c) {
  if ("string" === typeof a) {
    return this.addMatcher(a, c);
  }
  this.replacer || (this.replacer = []);
  this.replacer.push(a, c);
  this.cache && Q(this);
  return this;
};
w.encode = function(a, c) {
  if (this.cache && a.length <= this.L) {
    if (this.I) {
      if (this.B.has(a)) {
        return this.B.get(a);
      }
    } else {
      this.I = setTimeout(Q, 50, this);
    }
  }
  this.normalize && ("function" === typeof this.normalize ? a = this.normalize(a) : a = ka ? a.normalize("NFKD").replace(ka, "").toLowerCase() : a.toLowerCase());
  this.prepare && (a = this.prepare(a));
  this.numeric && 3 < a.length && (a = a.replace(ia, "$1 $2").replace(ja, "$1 $2").replace(ha, "$1 "));
  const b = !(this.dedupe || this.mapper || this.filter || this.matcher || this.stemmer || this.replacer);
  let e = [], d = I(), f, g, h = this.split || "" === this.split ? a.split(this.split) : [a];
  for (let m = 0, l, n; m < h.length; m++) {
    if ((l = n = h[m]) && !(l.length < this.minlength || l.length > this.maxlength)) {
      if (c) {
        if (d[l]) {
          continue;
        }
        d[l] = 1;
      } else {
        if (f === l) {
          continue;
        }
        f = l;
      }
      if (b) {
        e.push(l);
      } else {
        if (!this.filter || ("function" === typeof this.filter ? this.filter(l) : !this.filter.has(l))) {
          if (this.cache && l.length <= this.M) {
            if (this.I) {
              var k = this.H.get(l);
              if (k || "" === k) {
                k && e.push(k);
                continue;
              }
            } else {
              this.I = setTimeout(Q, 50, this);
            }
          }
          if (this.stemmer) {
            this.O || (this.O = new RegExp("(?!^)(" + this.A + ")$"));
            let u;
            for (; u !== l && 2 < l.length;) {
              u = l, l = l.replace(this.O, p => this.stemmer.get(p));
            }
          }
          if (l && (this.mapper || this.dedupe && 1 < l.length)) {
            k = "";
            for (let u = 0, p = "", t, r; u < l.length; u++) {
              t = l.charAt(u), t === p && this.dedupe || ((r = this.mapper && this.mapper.get(t)) || "" === r ? r === p && this.dedupe || !(p = r) || (k += r) : k += p = t);
            }
            l = k;
          }
          this.matcher && 1 < l.length && (this.N || (this.N = new RegExp("(" + this.h + ")", "g")), l = l.replace(this.N, u => this.matcher.get(u)));
          if (l && this.replacer) {
            for (k = 0; l && k < this.replacer.length; k += 2) {
              l = l.replace(this.replacer[k], this.replacer[k + 1]);
            }
          }
          this.cache && n.length <= this.M && (this.H.set(n, l), this.H.size > this.R && (this.H.clear(), this.M = this.M / 1.1 | 0));
          if (l) {
            if (l !== n) {
              if (c) {
                if (d[l]) {
                  continue;
                }
                d[l] = 1;
              } else {
                if (g === l) {
                  continue;
                }
                g = l;
              }
            }
            e.push(l);
          }
        }
      }
    }
  }
  this.finalize && (e = this.finalize(e) || e);
  this.cache && a.length <= this.L && (this.B.set(a, e), this.B.size > this.R && (this.B.clear(), this.L = this.L / 1.1 | 0));
  return e;
};
function Q(a) {
  a.I = null;
  a.B.clear();
  a.H.clear();
}
;function ma(a, c, b) {
  b || (c || "object" !== typeof a ? "object" === typeof c && (b = c, c = 0) : b = a);
  b && (a = b.query || a, c = b.limit || c);
  let e = "" + (c || 0);
  b && (e += (b.offset || 0) + !!b.context + !!b.suggest + (!1 !== b.resolve) + (b.resolution || this.resolution) + (b.boost || 0));
  a = ("" + a).toLowerCase();
  this.cache || (this.cache = new na());
  let d = this.cache.get(a + e);
  if (!d) {
    const f = b && b.cache;
    f && (b.cache = !1);
    d = this.search(a, c, b);
    f && (b.cache = f);
    this.cache.set(a + e, d);
  }
  return d;
}
function na(a) {
  this.limit = a && !0 !== a ? a : 1000;
  this.cache = new Map();
  this.h = "";
}
na.prototype.set = function(a, c) {
  this.cache.set(this.h = a, c);
  this.cache.size > this.limit && this.cache.delete(this.cache.keys().next().value);
};
na.prototype.get = function(a) {
  const c = this.cache.get(a);
  c && this.h !== a && (this.cache.delete(a), this.cache.set(this.h = a, c));
  return c;
};
na.prototype.remove = function(a) {
  for (const c of this.cache) {
    const b = c[0];
    c[1].includes(a) && this.cache.delete(b);
  }
};
na.prototype.clear = function() {
  this.cache.clear();
  this.h = "";
};
const qa = {normalize:!1, numeric:!1, dedupe:!1};
const ra = {};
const sa = new Map([["b", "p"], ["v", "f"], ["w", "f"], ["z", "s"], ["x", "s"], ["d", "t"], ["n", "m"], ["c", "k"], ["g", "k"], ["j", "k"], ["q", "k"], ["i", "e"], ["y", "e"], ["u", "o"]]);
const ta = new Map([["ae", "a"], ["oe", "o"], ["sh", "s"], ["kh", "k"], ["th", "t"], ["ph", "f"], ["pf", "f"]]), ua = [/([^aeo])h(.)/g, "$1$2", /([aeo])h([^aeo]|$)/g, "$1$2", /(.)\1+/g, "$1"];
const va = {a:"", e:"", i:"", o:"", u:"", y:"", b:1, f:1, p:1, v:1, c:2, g:2, j:2, k:2, q:2, s:2, x:2, z:2, "\u00df":2, d:3, t:3, l:4, m:5, n:5, r:6};
var wa = {Exact:qa, Default:ra, Normalize:ra, LatinBalance:{mapper:sa}, LatinAdvanced:{mapper:sa, matcher:ta, replacer:ua}, LatinExtra:{mapper:sa, replacer:ua.concat([/(?!^)[aeo]/g, ""]), matcher:ta}, LatinSoundex:{dedupe:!1, include:{letter:!0}, finalize:function(a) {
  for (let b = 0; b < a.length; b++) {
    var c = a[b];
    let e = c.charAt(0), d = va[e];
    for (let f = 1, g; f < c.length && (g = c.charAt(f), "h" === g || "w" === g || !(g = va[g]) || g === d || (e += g, d = g, 4 !== e.length)); f++) {
    }
    a[b] = e;
  }
}}, CJK:{split:""}, LatinExact:qa, LatinDefault:ra, LatinSimple:ra};
function za(a, c, b, e) {
  let d = [];
  for (let f = 0, g; f < a.index.length; f++) {
    if (g = a.index[f], c >= g.length) {
      c -= g.length;
    } else {
      c = g[e ? "splice" : "slice"](c, b);
      const h = c.length;
      if (h && (d = d.length ? d.concat(c) : c, b -= h, e && (a.length -= h), !b)) {
        break;
      }
      c = 0;
    }
  }
  return d;
}
function Aa(a) {
  if (!this || this.constructor !== Aa) {
    return new Aa(a);
  }
  this.index = a ? [a] : [];
  this.length = a ? a.length : 0;
  const c = this;
  return new Proxy([], {get(b, e) {
    if ("length" === e) {
      return c.length;
    }
    if ("push" === e) {
      return function(d) {
        c.index[c.index.length - 1].push(d);
        c.length++;
      };
    }
    if ("pop" === e) {
      return function() {
        if (c.length) {
          return c.length--, c.index[c.index.length - 1].pop();
        }
      };
    }
    if ("indexOf" === e) {
      return function(d) {
        let f = 0;
        for (let g = 0, h, k; g < c.index.length; g++) {
          h = c.index[g];
          k = h.indexOf(d);
          if (0 <= k) {
            return f + k;
          }
          f += h.length;
        }
        return -1;
      };
    }
    if ("includes" === e) {
      return function(d) {
        for (let f = 0; f < c.index.length; f++) {
          if (c.index[f].includes(d)) {
            return !0;
          }
        }
        return !1;
      };
    }
    if ("slice" === e) {
      return function(d, f) {
        return za(c, d || 0, f || c.length, !1);
      };
    }
    if ("splice" === e) {
      return function(d, f) {
        return za(c, d || 0, f || c.length, !0);
      };
    }
    if ("constructor" === e) {
      return Array;
    }
    if ("symbol" !== typeof e) {
      return (b = c.index[e / 2 ** 31 | 0]) && b[e];
    }
  }, set(b, e, d) {
    b = e / 2 ** 31 | 0;
    (c.index[b] || (c.index[b] = []))[e] = d;
    c.length++;
    return !0;
  }});
}
Aa.prototype.clear = function() {
  this.index.length = 0;
};
Aa.prototype.destroy = function() {
  this.proxy = this.index = null;
};
Aa.prototype.push = function() {
};
function R(a = 8) {
  if (!this || this.constructor !== R) {
    return new R(a);
  }
  this.index = I();
  this.h = [];
  this.size = 0;
  32 < a ? (this.B = Ba, this.A = BigInt(a)) : (this.B = Ca, this.A = a);
}
R.prototype.get = function(a) {
  const c = this.index[this.B(a)];
  return c && c.get(a);
};
R.prototype.set = function(a, c) {
  var b = this.B(a);
  let e = this.index[b];
  e ? (b = e.size, e.set(a, c), (b -= e.size) && this.size++) : (this.index[b] = e = new Map([[a, c]]), this.h.push(e), this.size++);
};
function S(a = 8) {
  if (!this || this.constructor !== S) {
    return new S(a);
  }
  this.index = I();
  this.h = [];
  this.size = 0;
  32 < a ? (this.B = Ba, this.A = BigInt(a)) : (this.B = Ca, this.A = a);
}
S.prototype.add = function(a) {
  var c = this.B(a);
  let b = this.index[c];
  b ? (c = b.size, b.add(a), (c -= b.size) && this.size++) : (this.index[c] = b = new Set([a]), this.h.push(b), this.size++);
};
w = R.prototype;
w.has = S.prototype.has = function(a) {
  const c = this.index[this.B(a)];
  return c && c.has(a);
};
w.delete = S.prototype.delete = function(a) {
  const c = this.index[this.B(a)];
  c && c.delete(a) && this.size--;
};
w.clear = S.prototype.clear = function() {
  this.index = I();
  this.h = [];
  this.size = 0;
};
w.values = S.prototype.values = function*() {
  for (let a = 0; a < this.h.length; a++) {
    for (let c of this.h[a].values()) {
      yield c;
    }
  }
};
w.keys = S.prototype.keys = function*() {
  for (let a = 0; a < this.h.length; a++) {
    for (let c of this.h[a].keys()) {
      yield c;
    }
  }
};
w.entries = S.prototype.entries = function*() {
  for (let a = 0; a < this.h.length; a++) {
    for (let c of this.h[a].entries()) {
      yield c;
    }
  }
};
function Ca(a) {
  let c = 2 ** this.A - 1;
  if ("number" == typeof a) {
    return a & c;
  }
  let b = 0, e = this.A + 1;
  for (let d = 0; d < a.length; d++) {
    b = (b * e ^ a.charCodeAt(d)) & c;
  }
  return 32 === this.A ? b + 2 ** 31 : b;
}
function Ba(a) {
  let c = BigInt(2) ** this.A - BigInt(1);
  var b = typeof a;
  if ("bigint" === b) {
    return a & c;
  }
  if ("number" === b) {
    return BigInt(a) & c;
  }
  b = BigInt(0);
  let e = this.A + BigInt(1);
  for (let d = 0; d < a.length; d++) {
    b = (b * e ^ BigInt(a.charCodeAt(d))) & c;
  }
  return b;
}
;let Da, T;
async function Ea(a) {
  a = a.data;
  var c = a.task;
  const b = a.id;
  let e = a.args;
  switch(c) {
    case "init":
      T = a.options || {};
      (c = a.factory) ? (Function("return " + c)()(self), Da = new self.FlexSearch.Index(T), delete self.FlexSearch) : Da = new V(T);
      postMessage({id:b});
      break;
    default:
      let d;
      if ("export" === c) {
        if (!T.export || "function" !== typeof T.export) {
          throw Error('Either no extern configuration provided for the Worker-Index or no method was defined on the config property "export".');
        }
        e[1] ? (e[0] = T.export, e[2] = 0, e[3] = 1) : e = null;
      }
      if ("import" === c) {
        if (!T.import || "function" !== typeof T.import) {
          throw Error('Either no extern configuration provided for the Worker-Index or no method was defined on the config property "import".');
        }
        e[0] && (a = await T.import.call(Da, e[0]), Da.import(e[0], a));
      } else {
        (d = e && Da[c].apply(Da, e)) && d.then && (d = await d), d && d.await && (d = await d.await), "search" === c && d.result && (d = d.result);
      }
      postMessage("search" === c ? {id:b, msg:d} : {id:b});
  }
}
;function Fa(a) {
  Ga.call(a, "add");
  Ga.call(a, "append");
  Ga.call(a, "search");
  Ga.call(a, "update");
  Ga.call(a, "remove");
  Ga.call(a, "searchCache");
}
let Ha, Ia, Ja;
function Ka() {
  Ha = Ja = 0;
}
function Ga(a) {
  this[a + "Async"] = function() {
    const c = arguments;
    var b = c[c.length - 1];
    let e;
    "function" === typeof b && (e = b, delete c[c.length - 1]);
    Ha ? Ja || (Ja = Date.now() - Ia >= this.priority * this.priority * 3) : (Ha = setTimeout(Ka, 0), Ia = Date.now());
    if (Ja) {
      const f = this;
      return new Promise(g => {
        setTimeout(function() {
          g(f[a + "Async"].apply(f, c));
        }, 0);
      });
    }
    const d = this[a].apply(this, c);
    b = d.then ? d : new Promise(f => f(d));
    e && b.then(e);
    return b;
  };
}
;let W = 0;
function La(a = {}, c) {
  function b(h) {
    function k(m) {
      m = m.data || m;
      const l = m.id, n = l && f.h[l];
      n && (n(m.msg), delete f.h[l]);
    }
    this.worker = h;
    this.h = I();
    if (this.worker) {
      d ? this.worker.on("message", k) : this.worker.onmessage = k;
      if (a.config) {
        return new Promise(function(m) {
          1e9 < W && (W = 0);
          f.h[++W] = function() {
            m(f);
          };
          f.worker.postMessage({id:W, task:"init", factory:e, options:a});
        });
      }
      this.priority = a.priority || 4;
      this.encoder = c || null;
      this.worker.postMessage({task:"init", factory:e, options:a});
      return this;
    }
    console.warn("Worker is not available on this platform. Please report on Github: https://github.com/nextapps-de/flexsearch/issues");
  }
  if (!this || this.constructor !== La) {
    return new La(a);
  }
  let e = "undefined" !== typeof self ? self._factory : "undefined" !== typeof window ? window._factory : null;
  e && (e = e.toString());
  const d = "undefined" === typeof window, f = this, g = Ma(e, d, a.worker);
  return g.then ? g.then(function(h) {
    return b.call(f, h);
  }) : b.call(this, g);
}
X("add");
X("append");
X("search");
X("update");
X("remove");
X("clear");
X("export");
X("import");
La.prototype.searchCache = ma;
Fa(La.prototype);
function X(a) {
  La.prototype[a] = function() {
    const c = this, b = [].slice.call(arguments);
    var e = b[b.length - 1];
    let d;
    "function" === typeof e && (d = e, b.pop());
    e = new Promise(function(f) {
      "export" === a && "function" === typeof b[0] && (b[0] = null);
      1e9 < W && (W = 0);
      c.h[++W] = f;
      c.worker.postMessage({task:a, id:W, args:b});
    });
    return d ? (e.then(d), this) : e;
  };
}
function Ma(a, c, b) {
  return c ? "undefined" !== typeof module ? new(require("worker_threads")["Worker"])(__dirname+"/worker/node.js") : import("worker_threads").then(function(worker){return new worker["Worker"](import.meta.dirname+"/node/node.mjs")}) : a ? new window.Worker(URL.createObjectURL(new Blob(["onmessage=" + Ea.toString()], {type:"text/javascript"}))) : new window.Worker("string" === typeof b ? b : import.meta.url.replace("/worker.js", "/worker/worker.js").replace("flexsearch.bundle.module.min.js", 
  "module/worker/worker.js"), {type:"module"});
}
;Na.prototype.add = function(a, c, b) {
  ba(a) && (c = a, a = da(c, this.key));
  if (c && (a || 0 === a)) {
    if (!b && this.reg.has(a)) {
      return this.update(a, c);
    }
    for (let h = 0, k; h < this.field.length; h++) {
      k = this.G[h];
      var e = this.index.get(this.field[h]);
      if ("function" === typeof k) {
        var d = k(c);
        d && e.add(a, d, !1, !0);
      } else {
        if (d = k.J, !d || d(c)) {
          k.constructor === String ? k = ["" + k] : N(k) && (k = [k]), Qa(c, k, this.K, 0, e, a, k[0], b);
        }
      }
    }
    if (this.tag) {
      for (e = 0; e < this.F.length; e++) {
        var f = this.F[e], g = this.P[e];
        d = this.tag.get(g);
        let h = I();
        if ("function" === typeof f) {
          if (f = f(c), !f) {
            continue;
          }
        } else {
          const k = f.J;
          if (k && !k(c)) {
            continue;
          }
          f.constructor === String && (f = "" + f);
          f = da(c, f);
        }
        if (d && f) {
          N(f) && (f = [f]);
          for (let k = 0, m, l; k < f.length; k++) {
            if (m = f[k], !h[m] && (h[m] = 1, (g = d.get(m)) ? l = g : d.set(m, l = []), !b || !l.includes(a))) {
              if (l.length === 2 ** 31 - 1) {
                g = new Aa(l);
                if (this.fastupdate) {
                  for (let n of this.reg.values()) {
                    n.includes(l) && (n[n.indexOf(l)] = g);
                  }
                }
                d.set(m, l = g);
              }
              l.push(a);
              this.fastupdate && ((g = this.reg.get(a)) ? g.push(l) : this.reg.set(a, [l]));
            }
          }
        } else {
          d || console.warn("Tag '" + g + "' was not found");
        }
      }
    }
    if (this.store && (!b || !this.store.has(a))) {
      let h;
      if (this.D) {
        h = I();
        for (let k = 0, m; k < this.D.length; k++) {
          m = this.D[k];
          if ((b = m.J) && !b(c)) {
            continue;
          }
          let l;
          if ("function" === typeof m) {
            l = m(c);
            if (!l) {
              continue;
            }
            m = [m.U];
          } else if (N(m) || m.constructor === String) {
            h[m] = c[m];
            continue;
          }
          Ra(c, h, m, 0, m[0], l);
        }
      }
      this.store.set(a, h || c);
    }
    this.worker && (this.fastupdate || this.reg.add(a));
  }
  return this;
};
function Ra(a, c, b, e, d, f) {
  a = a[d];
  if (e === b.length - 1) {
    c[d] = f || a;
  } else if (a) {
    if (a.constructor === Array) {
      for (c = c[d] = Array(a.length), d = 0; d < a.length; d++) {
        Ra(a, c, b, e, d);
      }
    } else {
      c = c[d] || (c[d] = I()), d = b[++e], Ra(a, c, b, e, d);
    }
  }
}
function Qa(a, c, b, e, d, f, g, h) {
  if (a = a[g]) {
    if (e === c.length - 1) {
      if (a.constructor === Array) {
        if (b[e]) {
          for (c = 0; c < a.length; c++) {
            d.add(f, a[c], !0, !0);
          }
          return;
        }
        a = a.join(" ");
      }
      d.add(f, a, h, !0);
    } else {
      if (a.constructor === Array) {
        for (g = 0; g < a.length; g++) {
          Qa(a, c, b, e, d, f, g, h);
        }
      } else {
        g = c[++e], Qa(a, c, b, e, d, f, g, h);
      }
    }
  } else {
    d.db && d.remove(f);
  }
}
;function Sa(a, c, b, e) {
  if (!a.length) {
    return a;
  }
  if (1 === a.length) {
    return a = a[0], a = b || a.length > c ? a.slice(b, b + c) : a, e ? Ta.call(this, a) : a;
  }
  let d = [];
  for (let f = 0, g, h; f < a.length; f++) {
    if ((g = a[f]) && (h = g.length)) {
      if (b) {
        if (b >= h) {
          b -= h;
          continue;
        }
        g = g.slice(b, b + c);
        h = g.length;
        b = 0;
      }
      h > c && (g = g.slice(0, c), h = c);
      if (!d.length && h >= c) {
        return e ? Ta.call(this, g) : g;
      }
      d.push(g);
      c -= h;
      if (!c) {
        break;
      }
    }
  }
  d = 1 < d.length ? [].concat.apply([], d) : d[0];
  return e ? Ta.call(this, d) : d;
}
;function Ua(a, c, b, e) {
  var d = e[0];
  if (d[0] && d[0].query) {
    return a[c].apply(a, d);
  }
  if (!("and" !== c && "not" !== c || a.result.length || a.await || d.suggest)) {
    return 1 < e.length && (d = e[e.length - 1]), (e = d.resolve) ? a.await || a.result : a;
  }
  let f = [], g = 0, h = 0, k, m, l;
  var n;
  let u;
  for (c = 0; c < e.length; c++) {
    if (d = e[c]) {
      var p = void 0;
      if (d.constructor === Y) {
        p = d.await || d.result;
      } else if (d.then || d.constructor === Array) {
        p = d;
      } else {
        g = d.limit || 0;
        h = d.offset || 0;
        l = d.suggest;
        m = d.resolve;
        k = (n = d.highlight && m) || d.enrich && m;
        p = d.queue;
        let t = d.async || p;
        (n = d.index) ? a.index || (a.index = n) : n = a.index;
        if (d.query || d.tag) {
          if (!a.index) {
            throw Error("Resolver can't apply because the corresponding Index was never specified");
          }
          const r = d.field || d.pluck;
          if (r) {
            if (!a.index.index) {
              throw Error("Resolver can't apply because the corresponding Document Index was not specified");
            }
            n = a.index.index.get(r);
            if (!n) {
              throw Error("Resolver can't apply because the specified Document Field '" + r + "' was not found");
            }
          }
          if (p && (u || a.await)) {
            u = 1;
            let q;
            const y = a.C.length, v = new Promise(function(A) {
              q = A;
            });
            (function(A, B) {
              v.h = function() {
                B.index = null;
                B.resolve = !1;
                let E = t ? A.searchAsync(B) : A.search(B);
                if (E.then) {
                  return E.then(function(C) {
                    a.C[y] = C = C.result || C;
                    q(C);
                    return C;
                  });
                }
                E = E.result || E;
                q(E);
                return E;
              };
            })(n, Object.assign({}, d));
            a.C.push(v);
            f[c] = v;
            continue;
          } else {
            d.resolve = !1, d.index = null, p = t ? n.searchAsync(d) : n.search(d), d.resolve = m, d.index = n;
          }
        } else if (d.and) {
          p = Va(d, "and", n);
        } else if (d.or) {
          p = Va(d, "or", n);
        } else if (d.not) {
          p = Va(d, "not", n);
        } else if (d.xor) {
          p = Va(d, "xor", n);
        } else {
          continue;
        }
      }
      p.await ? (u = 1, p = p.await) : p.then ? (u = 1, p = p.then(function(t) {
        return t.result || t;
      })) : p = p.result || p;
      f[c] = p;
    }
  }
  u && !a.await && (a.await = new Promise(function(t) {
    a.return = t;
  }));
  if (u) {
    const t = Promise.all(f).then(function(r) {
      for (let q = 0; q < a.C.length; q++) {
        if (a.C[q] === t) {
          a.C[q] = function() {
            return b.call(a, r, g, h, k, m, l);
          };
          break;
        }
      }
      Wa(a);
    });
    a.C.push(t);
  } else if (a.await) {
    a.C.push(function() {
      return b.call(a, f, g, h, k, m, l);
    });
  } else {
    return b.call(a, f, g, h, k, m, l);
  }
  return m ? a.await || a.result : a;
}
function Va(a, c, b) {
  a = a[c];
  const e = a[0] || a;
  e.index || (e.index = b);
  b = new Y(e);
  1 < a.length && (b = b[c].apply(b, a.slice(1)));
  return b;
}
;Y.prototype.or = function() {
  return Ua(this, "or", Xa, arguments);
};
function Xa(a, c, b, e, d) {
  a.length && (this.result.length && a.push(this.result), 2 > a.length ? this.result = a[0] : (this.result = Ya(a, c, b, !1, this.h), b = 0));
  d && (this.await = null);
  return d ? this.resolve(c, b, e) : this;
}
;Y.prototype.and = function() {
  return Ua(this, "and", Za, arguments);
};
function Za(a, c, b, e, d, f) {
  if (!f && !this.result.length) {
    return d ? this.result : this;
  }
  let g;
  if (a.length) {
    if (this.result.length && a.unshift(this.result), 2 > a.length) {
      this.result = a[0];
    } else {
      let h = 0;
      for (let k = 0, m, l; k < a.length; k++) {
        if ((m = a[k]) && (l = m.length)) {
          h < l && (h = l);
        } else if (!f) {
          h = 0;
          break;
        }
      }
      h ? (this.result = $a(a, h, c, b, f, this.h, d), g = !0) : this.result = [];
    }
  } else {
    f || (this.result = a);
  }
  d && (this.await = null);
  return d ? this.resolve(c, b, e, g) : this;
}
;Y.prototype.xor = function() {
  return Ua(this, "xor", ab, arguments);
};
function ab(a, c, b, e, d, f) {
  if (a.length) {
    if (this.result.length && a.unshift(this.result), 2 > a.length) {
      this.result = a[0];
    } else {
      a: {
        f = b;
        var g = this.h;
        const h = [], k = I();
        let m = 0;
        for (let l = 0, n; l < a.length; l++) {
          if (n = a[l]) {
            m < n.length && (m = n.length);
            for (let u = 0, p; u < n.length; u++) {
              if (p = n[u]) {
                for (let t = 0, r; t < p.length; t++) {
                  r = p[t], k[r] = k[r] ? 2 : 1;
                }
              }
            }
          }
        }
        for (let l = 0, n, u = 0; l < m; l++) {
          for (let p = 0, t; p < a.length; p++) {
            if (t = a[p]) {
              if (n = t[l]) {
                for (let r = 0, q; r < n.length; r++) {
                  if (q = n[r], 1 === k[q]) {
                    if (f) {
                      f--;
                    } else {
                      if (d) {
                        if (h.push(q), h.length === c) {
                          a = h;
                          break a;
                        }
                      } else {
                        const y = l + (p ? g : 0);
                        h[y] || (h[y] = []);
                        h[y].push(q);
                        if (++u === c) {
                          a = h;
                          break a;
                        }
                      }
                    }
                  }
                }
              }
            }
          }
        }
        a = h;
      }
      this.result = a;
      g = !0;
    }
  } else {
    f || (this.result = a);
  }
  d && (this.await = null);
  return d ? this.resolve(c, b, e, g) : this;
}
;Y.prototype.not = function() {
  return Ua(this, "not", bb, arguments);
};
function bb(a, c, b, e, d, f) {
  if (!f && !this.result.length) {
    return d ? this.result : this;
  }
  if (a.length && this.result.length) {
    a: {
      f = b;
      var g = [];
      a = new Set(a.flat().flat());
      for (let h = 0, k, m = 0; h < this.result.length; h++) {
        if (k = this.result[h]) {
          for (let l = 0, n; l < k.length; l++) {
            if (n = k[l], !a.has(n)) {
              if (f) {
                f--;
              } else {
                if (d) {
                  if (g.push(n), g.length === c) {
                    a = g;
                    break a;
                  }
                } else {
                  if (g[h] || (g[h] = []), g[h].push(n), ++m === c) {
                    a = g;
                    break a;
                  }
                }
              }
            }
          }
        }
      }
      a = g;
    }
    this.result = a;
    g = !0;
  }
  d && (this.await = null);
  return d ? this.resolve(c, b, e, g) : this;
}
;function Y(a, c) {
  if (!this || this.constructor !== Y) {
    return new Y(a, c);
  }
  let b = 0, e, d, f;
  if (a && a.index) {
    const g = a;
    c = g.index;
    b = g.boost || 0;
    if (g.query) {
      const h = g.resolve;
      a = g.async || g.queue;
      g.resolve = !1;
      g.index = null;
      a = a ? c.searchAsync(g) : c.search(g);
      g.resolve = h;
      g.index = c;
      a = a.result || a;
    } else {
      a = [];
    }
  }
  if (a && a.then) {
    const g = this;
    a = a.then(function(h) {
      g.C[0] = g.result = h.result || h;
      Wa(g);
    });
    e = [a];
    a = [];
    d = new Promise(function(h) {
      f = h;
    });
  }
  this.index = c || null;
  this.result = a || [];
  this.h = b;
  this.C = e || [];
  this.await = d || null;
  this.return = f || null;
}
Y.prototype.limit = function(a) {
  if (this.result.length) {
    const c = [];
    for (let b = 0, e; b < this.result.length; b++) {
      if (e = this.result[b]) {
        if (e.length <= a) {
          if (c[b] = e, a -= e.length, !a) {
            break;
          }
        } else {
          c[b] = e.slice(0, a);
          break;
        }
      }
    }
    this.result = c;
  }
  return this;
};
Y.prototype.offset = function(a) {
  if (this.result.length) {
    const c = [];
    for (let b = 0, e; b < this.result.length; b++) {
      if (e = this.result[b]) {
        e.length <= a ? a -= e.length : (c[b] = e.slice(a), a = 0);
      }
    }
    this.result = c;
  }
  return this;
};
Y.prototype.boost = function(a) {
  this.h += a;
  return this;
};
function Wa(a, c) {
  let b = a.result;
  for (let d = 0, f; d < a.C.length; d++) {
    if (f = a.C[d]) {
      if ("function" === typeof f) {
        b = f(), a.C[d] = b = b.result || b, d--;
      } else if (f.h) {
        b = f.h(), a.C[d] = b = b.result || b, d--;
      } else if (f.then) {
        return a.await;
      }
    }
  }
  const e = a.return;
  a.result = b;
  a.C = [];
  a.await = null;
  a.return = null;
  c || e(b);
  return b;
}
Y.prototype.resolve = function(a, c, b, e) {
  let d = this.await ? Wa(this, !0) : this.result;
  if (d.then) {
    const f = this;
    return d.then(function() {
      return f.resolve(a, c, b);
    });
  }
  d.length && ("object" === typeof a && (b = a.enrich, c = a.offset, a.highlight && console.warn('Highlighting results is not supported within the resolve() method. Instead pass highlight options within the last resolver stage like { query: "...", resolve: true, highlight: ... }.'), a = a.limit), d = e ? b ? Ta.call(this.index, d) : d : Sa.call(this.index, d, a || 100, c, b));
  e = this.return;
  this.return = this.await = this.C = this.result = this.index = null;
  e && e(d);
  return d;
};
function $a(a, c, b, e, d, f, g) {
  const h = a.length;
  let k = [], m, l;
  m = I();
  for (let n = 0, u, p, t, r; n < c; n++) {
    for (let q = 0; q < h; q++) {
      if (t = a[q], n < t.length && (u = t[n])) {
        for (let y = 0; y < u.length; y++) {
          p = u[y];
          (l = m[p]) ? m[p]++ : (l = 0, m[p] = 1);
          r = k[l] || (k[l] = []);
          if (!g) {
            let v = n + (q || !d ? 0 : f || 0);
            r = r[v] || (r[v] = []);
          }
          r.push(p);
          if (g && b && l === h - 1 && r.length - e === b) {
            return e ? r.slice(e) : r;
          }
        }
      }
    }
  }
  if (a = k.length) {
    if (d) {
      k = 1 < k.length ? Ya(k, b, e, g, f) : (k = k[0]) && b && k.length > b || e ? k.slice(e, b + e) : k;
    } else {
      if (a < h) {
        return [];
      }
      k = k[a - 1];
      if (b || e) {
        if (g) {
          if (k.length > b || e) {
            k = k.slice(e, b + e);
          }
        } else {
          d = [];
          for (let n = 0, u; n < k.length; n++) {
            if (u = k[n]) {
              if (e && u.length > e) {
                e -= u.length;
              } else {
                if (b && u.length > b || e) {
                  u = u.slice(e, b + e), b -= u.length, e && (e -= u.length);
                }
                d.push(u);
                if (!b) {
                  break;
                }
              }
            }
          }
          k = d;
        }
      }
    }
  }
  return k;
}
function Ya(a, c, b, e, d) {
  const f = [], g = I();
  let h;
  var k = a.length;
  let m;
  if (e) {
    for (d = k - 1; 0 <= d; d--) {
      if (m = (e = a[d]) && e.length) {
        for (k = 0; k < m; k++) {
          if (h = e[k], !g[h]) {
            if (g[h] = 1, b) {
              b--;
            } else {
              if (f.push(h), f.length === c) {
                return f;
              }
            }
          }
        }
      }
    }
  } else {
    for (let l = k - 1, n, u = 0; 0 <= l; l--) {
      n = a[l];
      for (let p = 0; p < n.length; p++) {
        if (m = (e = n[p]) && e.length) {
          for (let t = 0; t < m; t++) {
            if (h = e[t], !g[h]) {
              if (g[h] = 1, b) {
                b--;
              } else {
                let r = (p + (l < k - 1 ? d || 0 : 0)) / (l + 1) | 0;
                (f[r] || (f[r] = [])).push(h);
                if (++u === c) {
                  return f;
                }
              }
            }
          }
        }
      }
    }
  }
  return f;
}
function cb(a, c, b) {
  const e = I(), d = [];
  for (let f = 0, g; f < c.length; f++) {
    g = c[f];
    for (let h = 0; h < g.length; h++) {
      e[g[h]] = 1;
    }
  }
  if (b) {
    for (let f = 0, g; f < a.length; f++) {
      g = a[f], e[g] && (d.push(g), e[g] = 0);
    }
  } else {
    for (let f = 0, g, h; f < a.result.length; f++) {
      for (g = a.result[f], c = 0; c < g.length; c++) {
        h = g[c], e[h] && ((d[f] || (d[f] = [])).push(h), e[h] = 0);
      }
    }
  }
  return d;
}
;I();
function eb(a, c, b, e, d) {
  let f, g, h;
  "string" === typeof d ? (f = d, d = "") : f = d.template;
  if (!f) {
    throw Error('No template pattern was specified by the search option "highlight"');
  }
  g = f.indexOf("$1");
  if (-1 === g) {
    throw Error('Invalid highlight template. The replacement pattern "$1" was not found in template: ' + f);
  }
  h = f.substring(g + 2);
  g = f.substring(0, g);
  let k = d && d.boundary, m = !d || !1 !== d.clip, l = d && d.merge && h && g && new RegExp(h + " " + g, "g");
  d = d && d.ellipsis;
  var n = 0;
  if ("object" === typeof d) {
    var u = d.template;
    n = u.length - 2;
    d = d.pattern;
  }
  "string" !== typeof d && (d = !1 === d ? "" : "...");
  n && (d = u.replace("$1", d));
  u = d.length - n;
  let p, t;
  "object" === typeof k && (p = k.before, 0 === p && (p = -1), t = k.after, 0 === t && (t = -1), k = k.total || 9e5);
  n = new Map();
  for (let Oa = 0, ea, db, oa; Oa < c.length; Oa++) {
    let pa;
    if (e) {
      pa = c, oa = e;
    } else {
      var r = c[Oa];
      oa = r.field;
      if (!oa) {
        continue;
      }
      pa = r.result;
    }
    db = b.get(oa);
    ea = db.encoder;
    r = n.get(ea);
    "string" !== typeof r && (r = ea.encode(a), n.set(ea, r));
    for (let xa = 0; xa < pa.length; xa++) {
      var q = pa[xa].doc;
      if (!q) {
        continue;
      }
      q = da(q, oa);
      if (!q) {
        continue;
      }
      var y = q.trim().split(/\s+/);
      if (!y.length) {
        continue;
      }
      q = "";
      var v = [];
      let ya = [];
      var A = -1, B = -1, E = 0;
      for (var C = 0; C < y.length; C++) {
        var F = y[C], z = ea.encode(F);
        z = 1 < z.length ? z.join(" ") : z[0];
        let x;
        if (z && F) {
          var D = F.length, K = (ea.split ? F.replace(ea.split, "") : F).length - z.length, G = "", L = 0;
          for (var O = 0; O < r.length; O++) {
            var P = r[O];
            if (P) {
              var M = P.length;
              M += K;
              L && M <= L || (P = z.indexOf(P), -1 < P && (G = (P ? F.substring(0, P) : "") + g + F.substring(P, P + M) + h + (P + M < D ? F.substring(P + M) : ""), L = M, x = !0));
            }
          }
          G && (k && (0 > A && (A = q.length + (q ? 1 : 0)), B = q.length + (q ? 1 : 0) + G.length, E += D, ya.push(v.length), v.push({match:G})), q += (q ? " " : "") + G);
        }
        if (!x) {
          F = y[C], q += (q ? " " : "") + F, k && v.push({text:F});
        } else if (k && E >= k) {
          break;
        }
      }
      E = ya.length * (f.length - 2);
      if (p || t || k && q.length - E > k) {
        if (E = k + E - 2 * u, C = B - A, 0 < p && (C += p), 0 < t && (C += t), C <= E) {
          y = p ? A - (0 < p ? p : 0) : A - ((E - C) / 2 | 0), v = t ? B + (0 < t ? t : 0) : y + E, m || (0 < y && " " !== q.charAt(y) && " " !== q.charAt(y - 1) && (y = q.indexOf(" ", y), 0 > y && (y = 0)), v < q.length && " " !== q.charAt(v - 1) && " " !== q.charAt(v) && (v = q.lastIndexOf(" ", v), v < B ? v = B : ++v)), q = (y ? d : "") + q.substring(y, v) + (v < q.length ? d : "");
        } else {
          B = [];
          A = {};
          E = {};
          C = {};
          F = {};
          z = {};
          G = K = D = 0;
          for (O = L = 1;;) {
            var U = void 0;
            for (let x = 0, J; x < ya.length; x++) {
              J = ya[x];
              if (G) {
                if (K !== G) {
                  if (C[x + 1]) {
                    continue;
                  }
                  J += G;
                  if (A[J]) {
                    D -= u;
                    E[x + 1] = 1;
                    C[x + 1] = 1;
                    continue;
                  }
                  if (J >= v.length - 1) {
                    if (J >= v.length) {
                      C[x + 1] = 1;
                      J >= y.length && (E[x + 1] = 1);
                      continue;
                    }
                    D -= u;
                  }
                  q = v[J].text;
                  if (M = t && z[x]) {
                    if (0 < M) {
                      if (q.length > M) {
                        if (C[x + 1] = 1, m) {
                          q = q.substring(0, M);
                        } else {
                          continue;
                        }
                      }
                      (M -= q.length) || (M = -1);
                      z[x] = M;
                    } else {
                      C[x + 1] = 1;
                      continue;
                    }
                  }
                  if (D + q.length + 1 <= k) {
                    q = " " + q, B[x] += q;
                  } else if (m) {
                    U = k - D - 1, 0 < U && (q = " " + q.substring(0, U), B[x] += q), C[x + 1] = 1;
                  } else {
                    C[x + 1] = 1;
                    continue;
                  }
                } else {
                  if (C[x]) {
                    continue;
                  }
                  J -= K;
                  if (A[J]) {
                    D -= u;
                    C[x] = 1;
                    E[x] = 1;
                    continue;
                  }
                  if (0 >= J) {
                    if (0 > J) {
                      C[x] = 1;
                      E[x] = 1;
                      continue;
                    }
                    D -= u;
                  }
                  q = v[J].text;
                  if (M = p && F[x]) {
                    if (0 < M) {
                      if (q.length > M) {
                        if (C[x] = 1, m) {
                          q = q.substring(q.length - M);
                        } else {
                          continue;
                        }
                      }
                      (M -= q.length) || (M = -1);
                      F[x] = M;
                    } else {
                      C[x] = 1;
                      continue;
                    }
                  }
                  if (D + q.length + 1 <= k) {
                    q += " ", B[x] = q + B[x];
                  } else if (m) {
                    U = q.length + 1 - (k - D), 0 <= U && U < q.length && (q = q.substring(U) + " ", B[x] = q + B[x]), C[x] = 1;
                  } else {
                    C[x] = 1;
                    continue;
                  }
                }
              } else {
                q = v[J].match;
                p && (F[x] = p);
                t && (z[x] = t);
                x && D++;
                let Pa;
                J ? !x && u && (D += u) : (E[x] = 1, C[x] = 1);
                J >= y.length - 1 ? Pa = 1 : J < v.length - 1 && v[J + 1].match ? Pa = 1 : u && (D += u);
                D -= f.length - 2;
                if (!x || D + q.length <= k) {
                  B[x] = q;
                } else {
                  U = L = O = E[x] = 0;
                  break;
                }
                Pa && (E[x + 1] = 1, C[x + 1] = 1);
              }
              D += q.length;
              U = A[J] = 1;
            }
            if (U) {
              K === G ? G++ : K++;
            } else {
              K === G ? L = 0 : O = 0;
              if (!L && !O) {
                break;
              }
              L ? (K++, G = K) : G++;
            }
          }
          q = "";
          for (let x = 0, J; x < B.length; x++) {
            J = (x && E[x] ? " " : (x && !d ? " " : "") + d) + B[x], q += J;
          }
          d && !E[B.length] && (q += d);
        }
      }
      l && (q = q.replace(l, " "));
      pa[xa].highlight = q;
    }
    if (e) {
      break;
    }
  }
  return c;
}
;Na.prototype.search = function(a, c, b, e) {
  b || (!c && ba(a) ? (b = a, a = "") : ba(c) && (b = c, c = 0));
  let d = [];
  var f = [];
  let g;
  let h, k, m, l, n;
  let u = 0, p = !0, t;
  if (b) {
    b.constructor === Array && (b = {index:b});
    a = b.query || a;
    g = b.pluck;
    h = b.merge;
    m = b.boost;
    n = g || b.field || (n = b.index) && (n.index ? null : n);
    var r = this.tag && b.tag;
    k = b.suggest;
    p = !1 !== b.resolve;
    l = b.cache;
    this.store && b.highlight && !p ? console.warn("Highlighting results can only be done on a final resolver task or when calling .resolve({ highlight: ... })") : this.store && b.enrich && !p && console.warn("Enrich results can only be done on a final resolver task or when calling .resolve({ enrich: true })");
    t = p && this.store && b.highlight;
    var q = !!t || p && this.store && b.enrich;
    c = b.limit || c;
    var y = b.offset || 0;
    c || (c = p ? 100 : 0);
    if (r && (!this.db || !e)) {
      r.constructor !== Array && (r = [r]);
      var v = [];
      for (let F = 0, z; F < r.length; F++) {
        z = r[F];
        if (N(z)) {
          throw Error("A tag option can't be a string, instead it needs a { field: tag } format.");
        }
        if (z.field && z.tag) {
          var A = z.tag;
          if (A.constructor === Array) {
            for (var B = 0; B < A.length; B++) {
              v.push(z.field, A[B]);
            }
          } else {
            v.push(z.field, A);
          }
        } else {
          A = Object.keys(z);
          for (let D = 0, K, G; D < A.length; D++) {
            if (K = A[D], G = z[K], G.constructor === Array) {
              for (B = 0; B < G.length; B++) {
                v.push(K, G[B]);
              }
            } else {
              v.push(K, G);
            }
          }
        }
      }
      if (!v.length) {
        throw Error("Your tag definition within the search options is probably wrong. No valid tags found.");
      }
      r = v;
      if (!a) {
        f = [];
        if (v.length) {
          for (r = 0; r < v.length; r += 2) {
            if (this.db) {
              e = this.index.get(v[r]);
              if (!e) {
                console.warn("Tag '" + v[r] + ":" + v[r + 1] + "' will be skipped because there is no field '" + v[r] + "'.");
                continue;
              }
              f.push(e = e.db.tag(v[r + 1], c, y, q));
            } else {
              e = fb.call(this, v[r], v[r + 1], c, y, q);
            }
            d.push(p ? {field:v[r], tag:v[r + 1], result:e} : [e]);
          }
        }
        if (f.length) {
          const F = this;
          return Promise.all(f).then(function(z) {
            for (let D = 0; D < z.length; D++) {
              p ? d[D].result = z[D] : d[D] = z[D];
            }
            return p ? d : new Y(1 < d.length ? $a(d, 1, 0, 0, k, m) : d[0], F);
          });
        }
        return p ? d : new Y(1 < d.length ? $a(d, 1, 0, 0, k, m) : d[0], this);
      }
    }
    if (!p && !g) {
      if (n = n || this.field) {
        N(n) ? g = n : (n.constructor === Array && 1 === n.length && (n = n[0]), g = n.field || n.index);
      }
      if (!g) {
        throw Error("Apply resolver on document search requires either the option 'pluck' to be set or just select a single field name in your query.");
      }
    }
    n && n.constructor !== Array && (n = [n]);
  }
  n || (n = this.field);
  let E;
  v = (this.worker || this.db) && !e && [];
  for (let F = 0, z, D, K; F < n.length; F++) {
    D = n[F];
    if (this.db && this.tag && !this.G[F]) {
      continue;
    }
    let G;
    N(D) || (G = D, D = G.field, a = G.query || a, c = aa(G.limit, c), y = aa(G.offset, y), k = aa(G.suggest, k), t = p && this.store && aa(G.highlight, t), q = !!t || p && this.store && aa(G.enrich, q), l = aa(G.cache, l));
    if (e) {
      z = e[F];
    } else {
      A = G || b || {};
      B = A.enrich;
      var C = this.index.get(D);
      r && (this.db && (A.tag = r, E = C.db.support_tag_search, A.field = n), !E && B && (A.enrich = !1));
      z = l ? C.searchCache(a, c, A) : C.search(a, c, A);
      B && (A.enrich = B);
      if (v) {
        v[F] = z;
        continue;
      }
    }
    K = (z = z.result || z) && z.length;
    if (r && K) {
      A = [];
      B = 0;
      if (this.db && e) {
        if (!E) {
          for (C = n.length; C < e.length; C++) {
            let L = e[C];
            if (L && L.length) {
              B++, A.push(L);
            } else if (!k) {
              return p ? d : new Y(d, this);
            }
          }
        }
      } else {
        for (let L = 0, O, P; L < r.length; L += 2) {
          O = this.tag.get(r[L]);
          if (!O) {
            if (console.warn("Tag '" + r[L] + ":" + r[L + 1] + "' will be skipped because there is no field '" + r[L] + "'."), k) {
              continue;
            } else {
              return p ? d : new Y(d, this);
            }
          }
          if (P = (O = O && O.get(r[L + 1])) && O.length) {
            B++, A.push(O);
          } else if (!k) {
            return p ? d : new Y(d, this);
          }
        }
      }
      if (B) {
        z = cb(z, A, p);
        K = z.length;
        if (!K && !k) {
          return p ? z : new Y(z, this);
        }
        B--;
      }
    }
    if (K) {
      f[u] = D, d.push(z), u++;
    } else if (1 === n.length) {
      return p ? d : new Y(d, this);
    }
  }
  if (v) {
    if (this.db && r && r.length && !E) {
      for (q = 0; q < r.length; q += 2) {
        f = this.index.get(r[q]);
        if (!f) {
          if (console.warn("Tag '" + r[q] + ":" + r[q + 1] + "' was not found because there is no field '" + r[q] + "'."), k) {
            continue;
          } else {
            return p ? d : new Y(d, this);
          }
        }
        v.push(f.db.tag(r[q + 1], c, y, !1));
      }
    }
    const F = this;
    return Promise.all(v).then(function(z) {
      b && (b.resolve = p);
      z.length && (z = F.search(a, c, b, z));
      return z;
    });
  }
  if (!u) {
    return p ? d : new Y(d, this);
  }
  if (g && (!q || !this.store)) {
    return d = d[0], p ? d : new Y(d, this);
  }
  v = [];
  for (y = 0; y < f.length; y++) {
    r = d[y];
    q && r.length && "undefined" === typeof r[0].doc && (this.db ? v.push(r = this.index.get(this.field[0]).db.enrich(r)) : r = Ta.call(this, r));
    if (g) {
      return p ? t ? eb(a, r, this.index, g, t) : r : new Y(r, this);
    }
    d[y] = {field:f[y], result:r};
  }
  if (q && this.db && v.length) {
    const F = this;
    return Promise.all(v).then(function(z) {
      for (let D = 0; D < z.length; D++) {
        d[D].result = z[D];
      }
      t && (d = eb(a, d, F.index, g, t));
      return h ? gb(d) : d;
    });
  }
  t && (d = eb(a, d, this.index, g, t));
  return h ? gb(d) : d;
};
function gb(a) {
  const c = [], b = I(), e = I();
  for (let d = 0, f, g, h, k, m, l, n; d < a.length; d++) {
    f = a[d];
    g = f.field;
    h = f.result;
    for (let u = 0; u < h.length; u++) {
      if (m = h[u], "object" !== typeof m ? m = {id:k = m} : k = m.id, (l = b[k]) ? l.push(g) : (m.field = b[k] = [g], c.push(m)), n = m.highlight) {
        l = e[k], l || (e[k] = l = {}, m.highlight = l), l[g] = n;
      }
    }
  }
  return c;
}
function fb(a, c, b, e, d) {
  a = this.tag.get(a);
  if (!a) {
    return [];
  }
  a = a.get(c);
  if (!a) {
    return [];
  }
  c = a.length - e;
  if (0 < c) {
    if (b && c > b || e) {
      a = a.slice(e, e + b);
    }
    d && (a = Ta.call(this, a));
  }
  return a;
}
function Ta(a) {
  if (!this || !this.store) {
    return a;
  }
  if (this.db) {
    return this.index.get(this.field[0]).db.enrich(a);
  }
  const c = Array(a.length);
  for (let b = 0, e; b < a.length; b++) {
    e = a[b], c[b] = {id:e, doc:this.store.get(e)};
  }
  return c;
}
;function Na(a) {
  if (!this || this.constructor !== Na) {
    return new Na(a);
  }
  const c = a.document || a.doc || a;
  let b, e;
  this.G = [];
  this.field = [];
  this.K = [];
  this.key = (b = c.key || c.id) && hb(b, this.K) || "id";
  (e = a.keystore || 0) && (this.keystore = e);
  this.fastupdate = !!a.fastupdate;
  this.reg = !this.fastupdate || a.worker || a.db ? e ? new S(e) : new Set() : e ? new R(e) : new Map();
  this.D = (b = c.store || null) && b && !0 !== b && [];
  this.store = b && (e ? new R(e) : new Map());
  this.cache = (b = a.cache || null) && new na(b);
  a.cache = !1;
  this.worker = a.worker || !1;
  this.priority = a.priority || 4;
  this.index = ib.call(this, a, c);
  this.tag = null;
  if (b = c.tag) {
    if ("string" === typeof b && (b = [b]), b.length) {
      this.tag = new Map();
      this.F = [];
      this.P = [];
      for (let d = 0, f, g; d < b.length; d++) {
        f = b[d];
        g = f.field || f;
        if (!g) {
          throw Error("The tag field from the document descriptor is undefined.");
        }
        f.custom ? this.F[d] = f.custom : (this.F[d] = hb(g, this.K), f.filter && ("string" === typeof this.F[d] && (this.F[d] = new String(this.F[d])), this.F[d].J = f.filter));
        this.P[d] = g;
        this.tag.set(g, new Map());
      }
    }
  }
  if (this.worker) {
    this.fastupdate = !1;
    a = [];
    for (const d of this.index.values()) {
      d.then && a.push(d);
    }
    if (a.length) {
      const d = this;
      return Promise.all(a).then(function(f) {
        let g = 0;
        for (const h of d.index.entries()) {
          const k = h[0];
          let m = h[1];
          m.then && (m = f[g], d.index.set(k, m), g++);
        }
        return d;
      });
    }
  } else {
    a.db && (this.fastupdate = !1, this.mount(a.db));
  }
}
w = Na.prototype;
w.mount = function(a) {
  if (this.worker) {
    throw Error("You can't use Worker-Indexes on a persistent model. That would be useless, since each of the persistent model acts like Worker-Index by default (Master/Slave).");
  }
  let c = this.field;
  if (this.tag) {
    for (let f = 0, g; f < this.P.length; f++) {
      g = this.P[f];
      var b = void 0;
      this.index.set(g, b = new V({}, this.reg));
      c === this.field && (c = c.slice(0));
      c.push(g);
      b.tag = this.tag.get(g);
    }
  }
  b = [];
  const e = {db:a.db, type:a.type, fastupdate:a.fastupdate};
  for (let f = 0, g, h; f < c.length; f++) {
    e.field = h = c[f];
    g = this.index.get(h);
    const k = new a.constructor(a.id, e);
    k.id = a.id;
    b[f] = k.mount(g);
    g.document = !0;
    f ? g.bypass = !0 : g.store = this.store;
  }
  const d = this;
  return this.db = Promise.all(b).then(function() {
    d.db = !0;
  });
};
w.commit = async function(a, c) {
  const b = [];
  for (const e of this.index.values()) {
    b.push(e.commit(a, c));
  }
  await Promise.all(b);
  this.reg.clear();
};
w.destroy = function() {
  const a = [];
  for (const c of this.index.values()) {
    a.push(c.destroy());
  }
  return Promise.all(a);
};
function ib(a, c) {
  const b = new Map();
  let e = c.index || c.field || c;
  N(e) && (e = [e]);
  for (let f = 0, g, h; f < e.length; f++) {
    g = e[f];
    N(g) || (h = g, g = g.field);
    h = ba(h) ? Object.assign({}, a, h) : a;
    if (this.worker) {
      var d = void 0;
      d = (d = h.encoder) && d.encode ? d : new la("string" === typeof d ? wa[d] : d || {});
      d = new La(h, d);
      b.set(g, d);
    }
    this.worker || b.set(g, new V(h, this.reg));
    h.custom ? this.G[f] = h.custom : (this.G[f] = hb(g, this.K), h.filter && ("string" === typeof this.G[f] && (this.G[f] = new String(this.G[f])), this.G[f].J = h.filter));
    this.field[f] = g;
  }
  if (this.D) {
    a = c.store;
    N(a) && (a = [a]);
    for (let f = 0, g, h; f < a.length; f++) {
      g = a[f], h = g.field || g, g.custom ? (this.D[f] = g.custom, g.custom.U = h) : (this.D[f] = hb(h, this.K), g.filter && ("string" === typeof this.D[f] && (this.D[f] = new String(this.D[f])), this.D[f].J = g.filter));
    }
  }
  return b;
}
function hb(a, c) {
  const b = a.split(":");
  let e = 0;
  for (let d = 0; d < b.length; d++) {
    a = b[d], "]" === a[a.length - 1] && (a = a.substring(0, a.length - 2)) && (c[e] = !0), a && (b[e++] = a);
  }
  e < b.length && (b.length = e);
  return 1 < e ? b : b[0];
}
w.append = function(a, c) {
  return this.add(a, c, !0);
};
w.update = function(a, c) {
  return this.remove(a).add(a, c);
};
w.remove = function(a) {
  ba(a) && (a = da(a, this.key));
  for (var c of this.index.values()) {
    c.remove(a, !0);
  }
  if (this.reg.has(a)) {
    if (this.tag && !this.fastupdate) {
      for (let b of this.tag.values()) {
        for (let e of b) {
          c = e[0];
          const d = e[1], f = d.indexOf(a);
          -1 < f && (1 < d.length ? d.splice(f, 1) : b.delete(c));
        }
      }
    }
    this.store && this.store.delete(a);
    this.reg.delete(a);
  }
  this.cache && this.cache.remove(a);
  return this;
};
w.clear = function() {
  const a = [];
  for (const c of this.index.values()) {
    const b = c.clear();
    b.then && a.push(b);
  }
  if (this.tag) {
    for (const c of this.tag.values()) {
      c.clear();
    }
  }
  this.store && this.store.clear();
  this.cache && this.cache.clear();
  return a.length ? Promise.all(a) : this;
};
w.contain = function(a) {
  return this.db ? this.index.get(this.field[0]).db.has(a) : this.reg.has(a);
};
w.cleanup = function() {
  for (const a of this.index.values()) {
    a.cleanup();
  }
  return this;
};
w.get = function(a) {
  return this.db ? this.index.get(this.field[0]).db.enrich(a).then(function(c) {
    return c[0] && c[0].doc || null;
  }) : this.store.get(a) || null;
};
w.set = function(a, c) {
  "object" === typeof a && (c = a, a = da(c, this.key));
  this.store.set(a, c);
  return this;
};
w.searchCache = ma;
w.export = jb;
w.import = kb;
Fa(Na.prototype);
function lb(a, c = 0) {
  let b = [], e = [];
  c && (c = 250000 / c * 5000 | 0);
  for (const d of a.entries()) {
    e.push(d), e.length === c && (b.push(e), e = []);
  }
  e.length && b.push(e);
  return b;
}
function mb(a, c) {
  c || (c = new Map());
  for (let b = 0, e; b < a.length; b++) {
    e = a[b], c.set(e[0], e[1]);
  }
  return c;
}
function nb(a, c = 0) {
  let b = [], e = [];
  c && (c = 250000 / c * 1000 | 0);
  for (const d of a.entries()) {
    e.push([d[0], lb(d[1])[0]]), e.length === c && (b.push(e), e = []);
  }
  e.length && b.push(e);
  return b;
}
function ob(a, c) {
  c || (c = new Map());
  for (let b = 0, e, d; b < a.length; b++) {
    e = a[b], d = c.get(e[0]), c.set(e[0], mb(e[1], d));
  }
  return c;
}
function pb(a) {
  let c = [], b = [];
  for (const e of a.keys()) {
    b.push(e), 250000 === b.length && (c.push(b), b = []);
  }
  b.length && c.push(b);
  return c;
}
function qb(a, c) {
  c || (c = new Set());
  for (let b = 0; b < a.length; b++) {
    c.add(a[b]);
  }
  return c;
}
function rb(a, c, b, e, d, f, g = 0) {
  const h = e && e.constructor === Array;
  var k = h ? e.shift() : e;
  if (!k) {
    return this.export(a, c, d, f + 1);
  }
  if ((k = a((c ? c + "." : "") + (g + 1) + "." + b, JSON.stringify(k))) && k.then) {
    const m = this;
    return k.then(function() {
      return rb.call(m, a, c, b, h ? e : null, d, f, g + 1);
    });
  }
  return rb.call(this, a, c, b, h ? e : null, d, f, g + 1);
}
function jb(a, c, b = 0, e = 0) {
  if (b < this.field.length) {
    const g = this.field[b];
    if ((c = this.index.get(g).export(a, g, b, e = 1)) && c.then) {
      const h = this;
      return c.then(function() {
        return h.export(a, g, b + 1);
      });
    }
    return this.export(a, g, b + 1);
  }
  let d, f;
  switch(e) {
    case 0:
      d = "reg";
      f = pb(this.reg);
      c = null;
      break;
    case 1:
      d = "tag";
      f = this.tag && nb(this.tag, this.reg.size);
      c = null;
      break;
    case 2:
      d = "doc";
      f = this.store && lb(this.store);
      c = null;
      break;
    default:
      return;
  }
  return rb.call(this, a, c, d, f, b, e);
}
function kb(a, c) {
  var b = a.split(".");
  "json" === b[b.length - 1] && b.pop();
  const e = 2 < b.length ? b[0] : "";
  b = 2 < b.length ? b[2] : b[1];
  if (this.worker && e) {
    return this.index.get(e).import(a);
  }
  if (c) {
    "string" === typeof c && (c = JSON.parse(c));
    if (e) {
      return this.index.get(e).import(b, c);
    }
    switch(b) {
      case "reg":
        this.fastupdate = !1;
        this.reg = qb(c, this.reg);
        for (let d = 0, f; d < this.field.length; d++) {
          f = this.index.get(this.field[d]), f.fastupdate = !1, f.reg = this.reg;
        }
        if (this.worker) {
          c = [];
          for (const d of this.index.values()) {
            c.push(d.import(a));
          }
          return Promise.all(c);
        }
        break;
      case "tag":
        this.tag = ob(c, this.tag);
        break;
      case "doc":
        this.store = mb(c, this.store);
    }
  }
}
function sb(a, c) {
  let b = "";
  for (const e of a.entries()) {
    a = e[0];
    const d = e[1];
    let f = "";
    for (let g = 0, h; g < d.length; g++) {
      h = d[g] || [""];
      let k = "";
      for (let m = 0; m < h.length; m++) {
        k += (k ? "," : "") + ("string" === c ? '"' + h[m] + '"' : h[m]);
      }
      k = "[" + k + "]";
      f += (f ? "," : "") + k;
    }
    f = '["' + a + '",[' + f + "]]";
    b += (b ? "," : "") + f;
  }
  return b;
}
;V.prototype.remove = function(a, c) {
  const b = this.reg.size && (this.fastupdate ? this.reg.get(a) : this.reg.has(a));
  if (b) {
    if (this.fastupdate) {
      for (let e = 0, d, f; e < b.length; e++) {
        if ((d = b[e]) && (f = d.length)) {
          if (d[f - 1] === a) {
            d.pop();
          } else {
            const g = d.indexOf(a);
            0 <= g && d.splice(g, 1);
          }
        }
      }
    } else {
      tb(this.map, a), this.depth && tb(this.ctx, a);
    }
    c || this.reg.delete(a);
  }
  this.db && (this.commit_task.push({del:a}), this.S && ub(this));
  this.cache && this.cache.remove(a);
  return this;
};
function tb(a, c) {
  let b = 0;
  var e = "undefined" === typeof c;
  if (a.constructor === Array) {
    for (let d = 0, f, g, h; d < a.length; d++) {
      if ((f = a[d]) && f.length) {
        if (e) {
          return 1;
        }
        g = f.indexOf(c);
        if (0 <= g) {
          if (1 < f.length) {
            return f.splice(g, 1), 1;
          }
          delete a[d];
          if (b) {
            return 1;
          }
          h = 1;
        } else {
          if (h) {
            return 1;
          }
          b++;
        }
      }
    }
  } else {
    for (let d of a.entries()) {
      e = d[0], tb(d[1], c) ? b++ : a.delete(e);
    }
  }
  return b;
}
;const vb = {memory:{resolution:1}, performance:{resolution:3, fastupdate:!0, context:{depth:1, resolution:1}}, match:{tokenize:"forward"}, score:{resolution:9, context:{depth:2, resolution:3}}};
V.prototype.add = function(a, c, b, e) {
  if (c && (a || 0 === a)) {
    if (!e && !b && this.reg.has(a)) {
      return this.update(a, c);
    }
    e = this.depth;
    c = this.encoder.encode(c, !e);
    const m = c.length;
    if (m) {
      const l = I(), n = I(), u = this.resolution;
      for (let p = 0; p < m; p++) {
        let t = c[this.rtl ? m - 1 - p : p];
        var d = t.length;
        if (d && (e || !n[t])) {
          var f = this.score ? this.score(c, t, p, null, 0) : wb(u, m, p), g = "";
          switch(this.tokenize) {
            case "full":
              if (2 < d) {
                for (let r = 0, q; r < d; r++) {
                  for (f = d; f > r; f--) {
                    g = t.substring(r, f);
                    q = this.rtl ? d - 1 - r : r;
                    var h = this.score ? this.score(c, t, p, g, q) : wb(u, m, p, d, q);
                    xb(this, n, g, h, a, b);
                  }
                }
                break;
              }
            case "bidirectional":
            case "reverse":
              if (1 < d) {
                for (h = d - 1; 0 < h; h--) {
                  g = t[this.rtl ? d - 1 - h : h] + g;
                  var k = this.score ? this.score(c, t, p, g, h) : wb(u, m, p, d, h);
                  xb(this, n, g, k, a, b);
                }
                g = "";
              }
            case "forward":
              if (1 < d) {
                for (h = 0; h < d; h++) {
                  g += t[this.rtl ? d - 1 - h : h], xb(this, n, g, f, a, b);
                }
                break;
              }
            default:
              if (xb(this, n, t, f, a, b), e && 1 < m && p < m - 1) {
                for (d = I(), g = this.T, f = t, h = Math.min(e + 1, this.rtl ? p + 1 : m - p), d[f] = 1, k = 1; k < h; k++) {
                  if ((t = c[this.rtl ? m - 1 - p - k : p + k]) && !d[t]) {
                    d[t] = 1;
                    const r = this.score ? this.score(c, f, p, t, k - 1) : wb(g + (m / 2 > g ? 0 : 1), m, p, h - 1, k - 1), q = this.bidirectional && t > f;
                    xb(this, l, q ? f : t, r, a, b, q ? t : f);
                  }
                }
              }
          }
        }
      }
      this.fastupdate || this.reg.add(a);
    } else {
      c = "";
    }
  }
  this.db && (c || this.commit_task.push({del:a}), this.S && ub(this));
  return this;
};
function xb(a, c, b, e, d, f, g) {
  let h = g ? a.ctx : a.map, k;
  if (!c[b] || g && !(k = c[b])[g]) {
    if (g ? (c = k || (c[b] = I()), c[g] = 1, (k = h.get(g)) ? h = k : h.set(g, h = new Map())) : c[b] = 1, (k = h.get(b)) ? h = k : h.set(b, h = k = []), h = h[e] || (h[e] = []), !f || !h.includes(d)) {
      if (h.length === 2 ** 31 - 1) {
        c = new Aa(h);
        if (a.fastupdate) {
          for (let m of a.reg.values()) {
            m.includes(h) && (m[m.indexOf(h)] = c);
          }
        }
        k[e] = h = c;
      }
      h.push(d);
      a.fastupdate && ((e = a.reg.get(d)) ? e.push(h) : a.reg.set(d, [h]));
    }
  }
}
function wb(a, c, b, e, d) {
  return b && 1 < a ? c + (e || 0) <= a ? b + (d || 0) : (a - 1) / (c + (e || 0)) * (b + (d || 0)) + 1 | 0 : 0;
}
;V.prototype.search = function(a, c, b) {
  b || (c || "object" !== typeof a ? "object" === typeof c && (b = c, c = 0) : (b = a, a = ""));
  if (b && b.cache) {
    return b.cache = !1, a = this.searchCache(a, c, b), b.cache = !0, a;
  }
  let e = [], d, f, g, h = 0, k, m, l, n, u;
  b && (a = b.query || a, c = b.limit || c, h = b.offset || 0, f = b.context, g = b.suggest, u = (k = b.resolve) && b.enrich, l = b.boost, n = b.resolution, m = this.db && b.tag);
  "undefined" === typeof k && (k = this.resolve);
  f = this.depth && !1 !== f;
  let p = this.encoder.encode(a, !f);
  d = p.length;
  c = c || (k ? 100 : 0);
  if (1 === d) {
    return yb.call(this, p[0], "", c, h, k, u, m);
  }
  if (2 === d && f && !g) {
    return yb.call(this, p[1], p[0], c, h, k, u, m);
  }
  let t = I(), r = 0, q;
  f && (q = p[0], r = 1);
  n || 0 === n || (n = q ? this.T : this.resolution);
  if (this.db) {
    if (this.db.search && (b = this.db.search(this, p, c, h, g, k, u, m), !1 !== b)) {
      return b;
    }
    const y = this;
    return async function() {
      for (let v, A; r < d; r++) {
        if ((A = p[r]) && !t[A]) {
          t[A] = 1;
          v = await zb(y, A, q, 0, 0, !1, !1);
          if (v = Ab(v, e, g, n)) {
            e = v;
            break;
          }
          q && (g && v && e.length || (q = A));
        }
        g && q && r === d - 1 && !e.length && (n = y.resolution, q = "", r = -1, t = I());
      }
      return Bb(e, n, c, h, g, l, k);
    }();
  }
  for (let y, v; r < d; r++) {
    if ((v = p[r]) && !t[v]) {
      t[v] = 1;
      y = zb(this, v, q, 0, 0, !1, !1);
      if (y = Ab(y, e, g, n)) {
        e = y;
        break;
      }
      q && (g && y && e.length || (q = v));
    }
    g && q && r === d - 1 && !e.length && (n = this.resolution, q = "", r = -1, t = I());
  }
  return Bb(e, n, c, h, g, l, k);
};
function Bb(a, c, b, e, d, f, g) {
  let h = a.length, k = a;
  if (1 < h) {
    k = $a(a, c, b, e, d, f, g);
  } else if (1 === h) {
    return g ? Sa.call(null, a[0], b, e) : new Y(a[0], this);
  }
  return g ? k : new Y(k, this);
}
function yb(a, c, b, e, d, f, g) {
  a = zb(this, a, c, b, e, d, f, g);
  return this.db ? a.then(function(h) {
    return d ? h || [] : new Y(h, this);
  }) : a && a.length ? d ? Sa.call(this, a, b, e) : new Y(a, this) : d ? [] : new Y([], this);
}
function Ab(a, c, b, e) {
  let d = [];
  if (a && a.length) {
    if (a.length <= e) {
      c.push(a);
      return;
    }
    for (let f = 0, g; f < e; f++) {
      if (g = a[f]) {
        d[f] = g;
      }
    }
    if (d.length) {
      c.push(d);
      return;
    }
  }
  if (!b) {
    return d;
  }
}
function zb(a, c, b, e, d, f, g, h) {
  let k;
  b && (k = a.bidirectional && c > b) && (k = b, b = c, c = k);
  if (a.db) {
    return a.db.get(c, b, e, d, f, g, h);
  }
  a = b ? (a = a.ctx.get(b)) && a.get(c) : a.map.get(c);
  return a;
}
;function V(a, c) {
  if (!this || this.constructor !== V) {
    return new V(a);
  }
  if (a) {
    var b = N(a) ? a : a.preset;
    b && (vb[b] || console.warn("Preset not found: " + b), a = Object.assign({}, vb[b], a));
  } else {
    a = {};
  }
  b = a.context;
  const e = !0 === b ? {depth:1} : b || {}, d = N(a.encoder) ? wa[a.encoder] : a.encode || a.encoder || {};
  this.encoder = d.encode ? d : "object" === typeof d ? new la(d) : {encode:d};
  this.resolution = a.resolution || 9;
  this.tokenize = b = (b = a.tokenize) && "default" !== b && "exact" !== b && b || "strict";
  this.depth = "strict" === b && e.depth || 0;
  this.bidirectional = !1 !== e.bidirectional;
  this.fastupdate = !!a.fastupdate;
  this.score = a.score || null;
  e && e.depth && "strict" !== this.tokenize && console.warn('Context-Search could not applied, because it is just supported when using the tokenizer "strict".');
  (b = a.keystore || 0) && (this.keystore = b);
  this.map = b ? new R(b) : new Map();
  this.ctx = b ? new R(b) : new Map();
  this.reg = c || (this.fastupdate ? b ? new R(b) : new Map() : b ? new S(b) : new Set());
  this.T = e.resolution || 3;
  this.rtl = d.rtl || a.rtl || !1;
  this.cache = (b = a.cache || null) && new na(b);
  this.resolve = !1 !== a.resolve;
  if (b = a.db) {
    this.db = this.mount(b);
  }
  this.S = !1 !== a.commit;
  this.commit_task = [];
  this.commit_timer = null;
  this.priority = a.priority || 4;
}
w = V.prototype;
w.mount = function(a) {
  this.commit_timer && (clearTimeout(this.commit_timer), this.commit_timer = null);
  return a.mount(this);
};
w.commit = function(a, c) {
  this.commit_timer && (clearTimeout(this.commit_timer), this.commit_timer = null);
  return this.db.commit(this, a, c);
};
w.destroy = function() {
  this.commit_timer && (clearTimeout(this.commit_timer), this.commit_timer = null);
  return this.db.destroy();
};
function ub(a) {
  a.commit_timer || (a.commit_timer = setTimeout(function() {
    a.commit_timer = null;
    a.db.commit(a, void 0, void 0);
  }, 1));
}
w.clear = function() {
  this.map.clear();
  this.ctx.clear();
  this.reg.clear();
  this.cache && this.cache.clear();
  this.db && (this.commit_timer && clearTimeout(this.commit_timer), this.commit_timer = null, this.commit_task = [{clear:!0}]);
  return this;
};
w.append = function(a, c) {
  return this.add(a, c, !0);
};
w.contain = function(a) {
  return this.db ? this.db.has(a) : this.reg.has(a);
};
w.update = function(a, c) {
  const b = this, e = this.remove(a);
  return e && e.then ? e.then(() => b.add(a, c)) : this.add(a, c);
};
w.cleanup = function() {
  if (!this.fastupdate) {
    return console.info('Cleanup the index isn\'t required when not using "fastupdate".'), this;
  }
  tb(this.map);
  this.depth && tb(this.ctx);
  return this;
};
w.searchCache = ma;
w.export = function(a, c, b = 0, e = 0) {
  let d, f;
  switch(e) {
    case 0:
      d = "reg";
      f = pb(this.reg);
      break;
    case 1:
      d = "cfg";
      f = null;
      break;
    case 2:
      d = "map";
      f = lb(this.map, this.reg.size);
      break;
    case 3:
      d = "ctx";
      f = nb(this.ctx, this.reg.size);
      break;
    default:
      return;
  }
  return rb.call(this, a, c, d, f, b, e);
};
w.import = function(a, c) {
  if (c) {
    switch("string" === typeof c && (c = JSON.parse(c)), a = a.split("."), "json" === a[a.length - 1] && a.pop(), 3 === a.length && a.shift(), a = 1 < a.length ? a[1] : a[0], a) {
      case "reg":
        this.fastupdate = !1;
        this.reg = qb(c, this.reg);
        break;
      case "map":
        this.map = mb(c, this.map);
        break;
      case "ctx":
        this.ctx = ob(c, this.ctx);
    }
  }
};
w.serialize = function(a = !0) {
  let c = "", b = "", e = "";
  if (this.reg.size) {
    let f;
    for (var d of this.reg.keys()) {
      f || (f = typeof d), c += (c ? "," : "") + ("string" === f ? '"' + d + '"' : d);
    }
    c = "index.reg=new Set([" + c + "]);";
    b = sb(this.map, f);
    b = "index.map=new Map([" + b + "]);";
    for (const g of this.ctx.entries()) {
      d = g[0];
      let h = sb(g[1], f);
      h = "new Map([" + h + "])";
      h = '["' + d + '",' + h + "]";
      e += (e ? "," : "") + h;
    }
    e = "index.ctx=new Map([" + e + "]);";
  }
  return a ? "function inject(index){" + c + b + e + "}" : c + b + e;
};
Fa(V.prototype);
const Cb = "undefined" !== typeof window && (window.indexedDB || window.mozIndexedDB || window.webkitIndexedDB || window.msIndexedDB), Db = ["map", "ctx", "tag", "reg", "cfg"], Eb = I();
function Fb(a, c = {}) {
  if (!this || this.constructor !== Fb) {
    return new Fb(a, c);
  }
  "object" === typeof a && (c = a, a = a.name);
  a || console.info("Default storage space was used, because a name was not passed.");
  this.id = "flexsearch" + (a ? ":" + a.toLowerCase().replace(/[^a-z0-9_\-]/g, "") : "");
  this.field = c.field ? c.field.toLowerCase().replace(/[^a-z0-9_\-]/g, "") : "";
  this.type = c.type;
  this.fastupdate = this.support_tag_search = !1;
  this.db = null;
  this.h = {};
}
w = Fb.prototype;
w.mount = function(a) {
  if (a.index) {
    return a.mount(this);
  }
  a.db = this;
  return this.open();
};
w.open = function() {
  if (this.db) {
    return this.db;
  }
  let a = this;
  navigator.storage && navigator.storage.persist();
  Eb[a.id] || (Eb[a.id] = []);
  Eb[a.id].push(a.field);
  const c = Cb.open(a.id, 1);
  c.onupgradeneeded = function() {
    const b = a.db = this.result;
    for (let e = 0, d; e < Db.length; e++) {
      d = Db[e];
      for (let f = 0, g; f < Eb[a.id].length; f++) {
        g = Eb[a.id][f], b.objectStoreNames.contains(d + ("reg" !== d ? g ? ":" + g : "" : "")) || b.createObjectStore(d + ("reg" !== d ? g ? ":" + g : "" : ""));
      }
    }
  };
  return a.db = Z(c, function(b) {
    a.db = b;
    a.db.onversionchange = function() {
      a.close();
    };
  });
};
w.close = function() {
  this.db && this.db.close();
  this.db = null;
};
w.destroy = function() {
  const a = Cb.deleteDatabase(this.id);
  return Z(a);
};
w.clear = function() {
  const a = [];
  for (let b = 0, e; b < Db.length; b++) {
    e = Db[b];
    for (let d = 0, f; d < Eb[this.id].length; d++) {
      f = Eb[this.id][d], a.push(e + ("reg" !== e ? f ? ":" + f : "" : ""));
    }
  }
  const c = this.db.transaction(a, "readwrite");
  for (let b = 0; b < a.length; b++) {
    c.objectStore(a[b]).clear();
  }
  return Z(c);
};
w.get = function(a, c, b = 0, e = 0, d = !0, f = !1) {
  a = this.db.transaction((c ? "ctx" : "map") + (this.field ? ":" + this.field : ""), "readonly").objectStore((c ? "ctx" : "map") + (this.field ? ":" + this.field : "")).get(c ? c + ":" + a : a);
  const g = this;
  return Z(a).then(function(h) {
    let k = [];
    if (!h || !h.length) {
      return k;
    }
    if (d) {
      if (!b && !e && 1 === h.length) {
        return h[0];
      }
      for (let m = 0, l; m < h.length; m++) {
        if ((l = h[m]) && l.length) {
          if (e >= l.length) {
            e -= l.length;
            continue;
          }
          const n = b ? e + Math.min(l.length - e, b) : l.length;
          for (let u = e; u < n; u++) {
            k.push(l[u]);
          }
          e = 0;
          if (k.length === b) {
            break;
          }
        }
      }
      return f ? g.enrich(k) : k;
    }
    return h;
  });
};
w.tag = function(a, c = 0, b = 0, e = !1) {
  a = this.db.transaction("tag" + (this.field ? ":" + this.field : ""), "readonly").objectStore("tag" + (this.field ? ":" + this.field : "")).get(a);
  const d = this;
  return Z(a).then(function(f) {
    if (!f || !f.length || b >= f.length) {
      return [];
    }
    if (!c && !b) {
      return f;
    }
    f = f.slice(b, b + c);
    return e ? d.enrich(f) : f;
  });
};
w.enrich = function(a) {
  "object" !== typeof a && (a = [a]);
  const c = this.db.transaction("reg", "readonly").objectStore("reg"), b = [];
  for (let e = 0; e < a.length; e++) {
    b[e] = Z(c.get(a[e]));
  }
  return Promise.all(b).then(function(e) {
    for (let d = 0; d < e.length; d++) {
      e[d] = {id:a[d], doc:e[d] ? JSON.parse(e[d]) : null};
    }
    return e;
  });
};
w.has = function(a) {
  a = this.db.transaction("reg", "readonly").objectStore("reg").getKey(a);
  return Z(a).then(function(c) {
    return !!c;
  });
};
w.search = null;
w.info = function() {
};
w.transaction = function(a, c, b) {
  a += "reg" !== a ? this.field ? ":" + this.field : "" : "";
  let e = this.h[a + ":" + c];
  if (e) {
    return b.call(this, e);
  }
  let d = this.db.transaction(a, c);
  this.h[a + ":" + c] = e = d.objectStore(a);
  const f = b.call(this, e);
  this.h[a + ":" + c] = null;
  return Z(d).finally(function() {
    d = e = null;
    return f;
  });
};
w.commit = async function(a, c, b) {
  if (c) {
    await this.clear(), a.commit_task = [];
  } else {
    let e = a.commit_task;
    a.commit_task = [];
    for (let d = 0, f; d < e.length; d++) {
      if (f = e[d], f.clear) {
        await this.clear();
        c = !0;
        break;
      } else {
        e[d] = f.del;
      }
    }
    c || (b || (e = e.concat(ca(a.reg))), e.length && await this.remove(e));
  }
  a.reg.size && (await this.transaction("map", "readwrite", function(e) {
    for (const d of a.map) {
      const f = d[0], g = d[1];
      g.length && (c ? e.put(g, f) : e.get(f).onsuccess = function() {
        let h = this.result;
        var k;
        if (h && h.length) {
          const m = Math.max(h.length, g.length);
          for (let l = 0, n, u; l < m; l++) {
            if ((u = g[l]) && u.length) {
              if ((n = h[l]) && n.length) {
                for (k = 0; k < u.length; k++) {
                  n.push(u[k]);
                }
              } else {
                h[l] = u;
              }
              k = 1;
            }
          }
        } else {
          h = g, k = 1;
        }
        k && e.put(h, f);
      });
    }
  }), await this.transaction("ctx", "readwrite", function(e) {
    for (const d of a.ctx) {
      const f = d[0], g = d[1];
      for (const h of g) {
        const k = h[0], m = h[1];
        m.length && (c ? e.put(m, f + ":" + k) : e.get(f + ":" + k).onsuccess = function() {
          let l = this.result;
          var n;
          if (l && l.length) {
            const u = Math.max(l.length, m.length);
            for (let p = 0, t, r; p < u; p++) {
              if ((r = m[p]) && r.length) {
                if ((t = l[p]) && t.length) {
                  for (n = 0; n < r.length; n++) {
                    t.push(r[n]);
                  }
                } else {
                  l[p] = r;
                }
                n = 1;
              }
            }
          } else {
            l = m, n = 1;
          }
          n && e.put(l, f + ":" + k);
        });
      }
    }
  }), a.store ? await this.transaction("reg", "readwrite", function(e) {
    for (const d of a.store) {
      const f = d[0], g = d[1];
      e.put("object" === typeof g ? JSON.stringify(g) : 1, f);
    }
  }) : a.bypass || await this.transaction("reg", "readwrite", function(e) {
    for (const d of a.reg.keys()) {
      e.put(1, d);
    }
  }), a.tag && await this.transaction("tag", "readwrite", function(e) {
    for (const d of a.tag) {
      const f = d[0], g = d[1];
      g.length && (e.get(f).onsuccess = function() {
        let h = this.result;
        h = h && h.length ? h.concat(g) : g;
        e.put(h, f);
      });
    }
  }), a.map.clear(), a.ctx.clear(), a.tag && a.tag.clear(), a.store && a.store.clear(), a.document || a.reg.clear());
};
function Gb(a, c, b) {
  const e = a.value;
  let d, f = 0;
  for (let g = 0, h; g < e.length; g++) {
    if (h = b ? e : e[g]) {
      for (let k = 0, m, l; k < c.length; k++) {
        if (l = c[k], m = h.indexOf(l), 0 <= m) {
          if (d = 1, 1 < h.length) {
            h.splice(m, 1);
          } else {
            e[g] = [];
            break;
          }
        }
      }
      f += h.length;
    }
    if (b) {
      break;
    }
  }
  f ? d && a.update(e) : a.delete();
  a.continue();
}
w.remove = function(a) {
  "object" !== typeof a && (a = [a]);
  return Promise.all([this.transaction("map", "readwrite", function(c) {
    c.openCursor().onsuccess = function() {
      const b = this.result;
      b && Gb(b, a);
    };
  }), this.transaction("ctx", "readwrite", function(c) {
    c.openCursor().onsuccess = function() {
      const b = this.result;
      b && Gb(b, a);
    };
  }), this.transaction("tag", "readwrite", function(c) {
    c.openCursor().onsuccess = function() {
      const b = this.result;
      b && Gb(b, a, !0);
    };
  }), this.transaction("reg", "readwrite", function(c) {
    for (let b = 0; b < a.length; b++) {
      c.delete(a[b]);
    }
  })]);
};
function Z(a, c) {
  return new Promise((b, e) => {
    a.onsuccess = a.oncomplete = function() {
      c && c(this.result);
      c = null;
      b(this.result);
    };
    a.onerror = a.onblocked = e;
    a = null;
  });
}
;export default {Index:V, Charset:wa, Encoder:la, Document:Na, Worker:La, Resolver:Y, IndexedDB:Fb, Language:{}};

export const Index=V;export const  Charset=wa;export const  Encoder=la;export const  Document=Na;export const  Worker=La;export const  Resolver=Y;export const  IndexedDB=Fb;export const  Language={};