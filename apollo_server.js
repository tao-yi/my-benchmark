import { ApolloServer, gql } from "apollo-server";

const typeDefs = gql`
  type Query {
    hello: String
  }
`;

const resolvers = {
  Query: {
    hello: () => "world",
  },
};
const port = 3999;

const app = new ApolloServer({ typeDefs, resolvers });

app.listen(port, () =>
  console.log(`Example app listening at http://localhost:${port}`),
);
