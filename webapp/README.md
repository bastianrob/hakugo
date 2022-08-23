# WEBAPP

This is a [Next.js](https://nextjs.org/) project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app) and uses [TypesSript](https://www.typescriptlang.org/)

## Prerequisite

1. Node Version Manager
2. Node v16.x.x
3. PNPM! not NPM not YARN!

```bash
nvm use v16;      # use Node v16
npm i -g pnpm;    # install pnpm
pnpm install;     # install all dependencies
```

Create `.env` file in the webapp worksapce directory

```bash
APP_HOST_GQL={proto}://{host}:{port}/{path}     # GraphQL server host
```

## Getting Started

```bash
pnpm run dev;
```

- Open [http://localhost](http://localhost) with your browser to see the result.
- Check [package.json](package.json) for available scripts, or to cconfigure running port