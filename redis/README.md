## Setup

1. Replication Mode
2. Auth enabled
<!-- 3. Volume & Claim in [persistence.yaml](persistence.yaml) -->

```bash
# kubectl apply -f persistence.yaml
helm install redis-7 bitnami/redis --values values.yaml --set auth.password={password}
```

## Resulting Service Endpoints

1. redis-7-master.default.svc.cluster.local for read/write operations (port 6379)
2. redis-7-replicas.default.svc.cluster.local for read-only operations (port 6379)

## Accessing Service

```bash
npm i -g redis-commander
kubectl port-forward --namespace default svc/redis-7-master 6379:6379

redis-commander --redis-password {password}
```
