import { Document, Charset } from "./libs/flexsearch.js";

export const initFlexSearch = async () => {
  const data = [
    {
      id: 1,
      title: "Carmencita",
      body: "Carmencita is a short film directed by Méliès. It is a silent film and was released in 1894. It is a documentary film about the life of a young woman named Carmencita.",
      tags: ["Documentary", "Short"],
    },
    {
      id: 2,
      title: "Le clown et ses chiens",
      body: "Le clown et ses chiens is a short film directed by Georges Méliès. It is a silent film and was released in 1892. It is a documentary film about the life of a young woman named Carmencita.",
      tags: ["Animation", "Short"],
    },
  ];

  const index = new Document({
    charset: Charset.LatinAdvanced,
    tokenize: "forward",
    resolution: 9,
    document: {
      id: "id",
      store: true,
      index: [
        {
          field: "title",
        },
        {
          field: "body",
        },
      ],
    },
  });

  for (let i = 0; i < data.length; i++) {
    await index.add(data[i]);
  }

  const result = await index.search("Georges", {
    field: ["title", "body"],
    limit: 10,
    suggest: true,
    enrich: true,
  });

  console.log(result);
  console.log(JSON.stringify(result, null, 2));
};
