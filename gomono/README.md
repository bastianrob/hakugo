ls# GOMONO

## Build Docker Image

```
docker build -t {image-name} --build-arg service={cmd-dirname} --build-arg port={service-port} --build-arg version={service-version} .
```

## Build Helm

```
helm template {service-name} ./gomono -f ./cmd/{service}/values.yaml \
  --set secret.APP_JWT_SECRET="" \
  --set secret.APP_GQL_AUTH_SECRET="" \
  --set configmap.APP_GQL_HOST="" \
  --set configmap.APP_GQL_AUTH_HEADER="" \
  --set configmap.APP_REDIS_HOST="" \
  --set configmap.APP_REDIS_PASS="" \
  --set configmap.APP_REDIS_DB="" \
  --set configmap.APP_MAILJET_APIKEY="" \
  --set configmap.APP_MAILJET_SECRET="" \
  --set image.tag="";
```
