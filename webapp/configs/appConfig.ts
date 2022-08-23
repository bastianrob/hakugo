import pkg from "package.json";

export const appConfig = {
  name: pkg.name,
  version: pkg.version,
  hosts: {
    graphql: process.env.APP_HOST_GQL,
  },
};

export default appConfig;
