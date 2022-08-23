import appConfig from "@/configs/appConfig";
import {ApolloClient, InMemoryCache} from "@apollo/client";

export const apolloClient = new ApolloClient({
  uri: appConfig.hosts.graphql,
  cache: new InMemoryCache(),
});

export default apolloClient;
