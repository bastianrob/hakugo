```bash
kubectl apply -f {yaml}
kubectl get all
kubectl get node -o wide;
```

```
minikube dashboard
minikube image load image:tag
minikube service list
minikube service {service-name} --url
```

```
helm install {service-name} --values {values.yaml} {chart-name}
helm upgrade {service-name} --values {values.yaml} {chart-name}
``