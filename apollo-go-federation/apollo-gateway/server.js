const { ApolloServer } = require('apollo-server');
const { ApolloGateway } = require("@apollo/gateway");

const gateway = new ApolloGateway({
    serviceList: [
        { name: 'accounts', url: 'http://gql-accounts:80/query' },
        { name: 'products', url: 'http://gql-products:80/query' },
        { name: 'reviews', url: 'http://gql-reviews:80/query' }
    ],
});

const server = new ApolloServer({
    gateway,

    subscriptions: false,
});

server.listen({ port: 4004 }).then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
});
