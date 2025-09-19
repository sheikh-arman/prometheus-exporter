
`go mod init`

`gmv`

`cobra-cli init`

`cobra-cli add start`

-> then write the business logic in exporter.go file

-> create dockerfile

`docker build  -t skaliarman/prometheus:latest .`

`docker push skaliarman/prometheus:latest`

-> test

`docker run --rm -it -p 8080:8080 skaliarman/prometheus:latest`

-> install prometheus stack

```helm
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install prometheus prometheus-community/kube-prometheus-stack -n monitoring --set grafana.image.tag=7.5.5 --create-namespace
```

deploy a deployment, service, service monitor kept on  `./helm` directory

`helm create prometheus-exporter`

metrics we are exposing is `http_requests_total` on '/metrics' endpoint.

graphana user: admin

password: prom-operator

-> port forward the deployment

`kubectl port-forward service/prometheus-exporter 8080:8080`

-> visit `http://localhost:8080/hi`

-> port forward prometheus and grafana

`kubectl port-forward -n monitoring service/prometheus-kube-prometheus-prometheus 9090:9090`

`kubectl port-forward -n monitoring service/prometheus-grafana 3000:80`

check `http_requests_total` on prometheus










