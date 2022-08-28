import appConfig from "@/configs/appConfig";
import {ApolloClient, InMemoryCache} from "@apollo/client";

export enum GraphqlErrorCodes {
  NotFound = "NOT_FOUND",
  InputMalformed = "INPUT_MALFORMED",
  ValidationError = "VALIDATION_ERROR",
  UnexpectedError = "UNEXPECTED_ERROR",
}

export const apolloClient = new ApolloClient({
  uri: appConfig.hosts.graphql,
  cache: new InMemoryCache(),
});

export default apolloClient;
