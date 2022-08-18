# HASURA!

## 1. Prerequisite

1. node v16
2. postgres v12
3. hasura cli
4. Prepare `.env` file in `/hasura` directory

```bash
# Local Postgres Config
POSTGRES_DB=""
POSTGRES_USER=""
POSTGRES_PASSWORD=""

# Hasura Config
PG_DATABASE_URL=""
HASURA_GRAPHQL_METADATA_DATABASE_URL=""
HASURA_GRAPHQL_JWT_SECRET='{"type":"HS256|HS512","key":"32ByteForHS256|64ByteForH512"}'
HASURA_GRAPHQL_ADMIN_SECRET=""
```

## 2. Quickstart

1. docker-compose up -d
2. hasura console

## 3. Applying Migrations

Importing metadata from local server to metadata directory

```bash
hasura metadata export
```

Creating first migration from existing PostgreSQL

```
hasura migrate create init --sql-from-server --endpoint <hasura-project-url> --admin-secret <admin-secret>
```

Deploying to remote server

```bash
hasura migrate apply --endpoint <hasura-project-url> --admin-secret <admin-secret>
hasura seed apply --endpoint <hasura-project-url> --admin-secret <admin-secret>
hasura metadata apply --endpoint <hasura-project-url> --admin-secret <admin-secret>
# OR
hasura deploy --endpoint <hasura-project-url> --admin-secret <admin-secret>
```

## References

1. [Hasura CLI](https://hasura.io/docs/latest/hasura-cli/commands/index/)
2. [3 Factor App](https://3factor.app/)
