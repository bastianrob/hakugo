import pkg from "package.json";

export const appConfig = {
  name: pkg.name,
  version: pkg.version,
  hosts: {
    graphql: process.env.NEXT_PUBLIC_APP_HOST_GQL,
  },
};

export default appConfig;
