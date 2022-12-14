import appConfig from "@/configs/appConfig";
import {ApolloClient, createHttpLink, InMemoryCache} from "@apollo/client";
import {setContext} from "@apollo/client/link/context";
import {getCookie} from "cookies-next";

export enum GraphqlErrorCodes {
  NotFound = "NOT_FOUND",
  InputMalformed = "INPUT_MALFORMED",
  ValidationError = "VALIDATION_ERROR",
  UnexpectedError = "UNEXPECTED_ERROR",
}

const httpLink = createHttpLink({
  uri: appConfig.hosts.graphql,
});

const authLink = setContext((_, {headers}) => {
  const token = getCookie("access-token");

  return !token
    ? headers
    : {
        headers: {...headers, authorization: `Bearer ${token}`},
      };
});

export const apolloClient = new ApolloClient({
  link: authLink.concat(httpLink),
  cache: new InMemoryCache(),
});

export default apolloClient;
