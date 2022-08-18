```bash
kubectl apply -f {yaml}
kubectl get all
kubectl get node -o wide;
```

```bash
minikube dashboard
minikube image load image:tag
minikube service list
minikube service {service-name} --url
```

```bash
helm install {service-name} --values {values.yaml} {chart-name}
helm upgrade {service-name} --values {values.yaml} {chart-name}
```

```bash
git diff --dirstat=files,0
```