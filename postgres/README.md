

## Setup

1. Replication Mode
2. One Primary, Three Read Replicas
3. Volume & Claim in [persistence.yaml](persistence.yaml)

```bash
kubectl apply -f persistence.yaml
helm install postgres-14 bitnami/postgresql --values values.yaml \
  --set auth.username={} --set auth.password={} --set auth.database={} \
  --set auth.postgresPassword={} --set auth.replicationPassword={}
```

## Resulting Service Endpoints

1. postgres-14-postgresql-primary.default.svc.cluster.local - Read/Write connection
2. postgres-14-postgresql-read.default.svc.cluster.local - Read only connection

## Accessing Service

```bash
kubectl port-forward --namespace default svc/postgres-14-postgresql-primary 5432:5432
```