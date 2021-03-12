import fastify from "fastify";

const app = fastify();
const port = 3999;

app.get("/", async (req, res) => {
  return { message: "hello world" };
});

app
  .listen(port)
  .then((v) =>
    console.log(`Example app listening at http://localhost:${port}`),
  );
