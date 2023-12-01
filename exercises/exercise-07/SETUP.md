# Einrichten eines Monitoring-Stacks auf Kubernetes
Die entsprechenden Manifests finden Sie unter dem Verzeichnis [monitoring](monitoring/). Hier einmal der Aufbau für den ersten Überblick:

```
monitoring
├── namespace.yaml
├── grafana
│   ├── datasource.yaml
│   ├── deployment.yaml
│   └── service.yaml
├── kube-state-metrics
│   ├── cluster-role-binding.yaml
│   ├── cluster-role.yaml
│   ├── deployment.yaml
│   ├── service-account.yaml
│   └── service.yaml
└── prometheus
    ├── cluster-role.yaml
    ├── config-map.yaml
    ├── deployment.yaml
    ├── node-exporter-daemonset.yaml
    ├── node-exporter-service.yaml
    └── service.yaml
```

## Initialisieren benötigter Resourcen für den Stack
```bash
kubectl apply -f monitoring/
```

## Deployment von Prometheus als Datenquelle
```bash
kubectl apply -f monitoring/prometheus/
```

## Deployment von Grafana als Visualisierungsplatform
```bash
kubectl apply -f monitoring/grafana/
```

## Deployment von Kube-State-Metrics für das Überwachen des Kubernetes API-Servers
```bash
kubectl apply -f monitoring/kube-state-metrics/
```

## Deployment aller Stacks
```bash
kubectl apply \
    -f monitoring/ \
    -f monitoring/prometheus/ \
    -f monitoring/kube-state-metrics/ \
    -f monitoring/grafana/
```