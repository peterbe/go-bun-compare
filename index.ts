Bun.serve({
  port: 3000,
  async fetch(_) {
    const jsonData = await Bun.file("data.json").arrayBuffer();
    return new Response(jsonData, {headers: {'content-type': "application/json"}});
  },
});
