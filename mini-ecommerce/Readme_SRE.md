# SRE Stack запуск

## 1. Запуск Docker инфраструктуры

```bash id="1m8v5g"
docker compose up --build
```

---

# Проверка сервисов

## Backend

```text id="2n7w1f"
http://localhost:8080
```

---

## Prometheus

```text id="3p2v8b"
http://localhost:9090
```

---

## Grafana

```text id="4q1m9l"
http://localhost:3000
```

Login:

```text id="5r7w2c"
admin
admin
```

---

## Kibana

```text id="6t4v1d"
http://localhost:5601
```

---

## Alertmanager

```text id="7v1m8e"
http://localhost:9093
```

---

# Prometheus

## Проверка targets

Prometheus → Status → Targets

Backend target должен быть:

```text id="8x2w5c"
UP
```

---

# Grafana

## Добавление Prometheus datasource

Connections → Data Sources → Add Data Source

Тип:

```text id="9p6m1f"
Prometheus
```

URL:

```text id="0q7v2l"
http://prometheus:9090
```

---

# Grafana Dashboard Queries

## HTTP Requests

```promql id="1r4m8c"
http_requests_total
```

---

## Request Latency

```promql id="2t9w1z"
rate(http_request_duration_seconds_sum[1m])
/
rate(http_request_duration_seconds_count[1m])
```

---

## Orders

```promql id="3v1m7b"
orders_created_total
```

---

## Heavy Endpoint

```promql id="4x8p2z"
heavy_requests_total
```

---

# Load Testing

## Создание venv

```bash id="5y1m9z"
python3 -m venv venv
```

---

## Активация

```bash id="6z4w2v"
source venv/bin/activate
```

---

## Установка Locust

```bash id="7a1m8z"
pip install locust
```

---

## Запуск Locust

```bash id="8b2v5x"
locust
```

---

## Locust UI

```text id="9c7w1x"
http://localhost:8089
```

Host:

```text id="0d4m8z"
http://localhost:8080
```

---

# Kubernetes

## Запуск Minikube

```bash id="1e9w2w"
minikube start
```

---

## Включение Metrics Server

```bash id="2f7v1r"
minikube addons enable metrics-server
```

---

## Применение Kubernetes manifests

```bash id="3g2m8v"
kubectl apply -f deployments/kubernetes/
```

---

## Проверка pod-ов

```bash id="4h9w1w"
kubectl get pods
```

---

## Проверка HPA

```bash id="5i4v8t"
kubectl get hpa
```

---

## Monitoring scaling

```bash id="6j1m7x"
kubectl get hpa -w
```

---

# Terraform

## Инициализация

```bash id="7k8w2t"
terraform init
```

---

## Проверка конфигурации

```bash id="8l2v5s"
terraform validate
```

---

## План инфраструктуры

```bash id="9m7w1w"
terraform plan
```

---

# GitHub Actions

## Push проекта

```bash id="0n4m8y"
git add .
git commit -m "update"
git push
```

---

## Проверка pipeline

GitHub Repository → Actions

Pipeline должен успешно:

* Build Go application
* Run tests
* Build Docker image
* Push Docker image
