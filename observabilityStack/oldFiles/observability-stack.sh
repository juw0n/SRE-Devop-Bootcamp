#!/bin/bash

# Define namespace
NAMESPACE="observability"

# Check if namespace exists, if not, create it
kubectl get namespace $NAMESPACE &>/dev/null
if [ $? -ne 0 ]; then
  echo "Namespace '$NAMESPACE' does not exist. Creating namespace..."
  kubectl create namespace $NAMESPACE
else
  echo "Namespace '$NAMESPACE' already exists."
fi

# Add Prometheus and Grafana Helm repositories
echo "Adding Helm repositories..."
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add grafana https://grafana.github.io/helm-charts
echo

# Update Helm repositories to ensure the latest versions
echo "Updating Helm repositories..."
helm repo update
echo

# Install Prometheus with node selector configuration
echo "Installing Prometheus..."
helm install prometheus prometheus-community/prometheus \
  --namespace $NAMESPACE \
  --set server.nodeSelector.service=dependent-services-node \
  --set alertmanager.nodeSelector.service=dependent-services-node \
  --set nodeExporter.nodeSelector.service=dependent-services-node \
  --set kube-state-metrics.nodeSelector.service=dependent-services-node \
  --set pushgateway.nodeSelector.service=dependent-services-node
echo

# Install Loki with Promtail using Grafana repo and node selector configuration
echo "Installing Loki (with Promtail)..."
helm install loki grafana/loki-stack --namespace $NAMESPACE -f node-selector.yaml
echo

# Install Grafana with node selector configuration
echo "Installing Grafana..."
helm install grafana grafana/grafana --namespace $NAMESPACE -f node-selector.yaml
echo

# Install Promtail
echo "Installing Promtail and configured to send only application logs to Loki"
helm upgrade --values promtail-config.yaml --install promtail grafana/promtail

# Deploy a DB Metrics Exporter which exposes database metrics that Prometheus can scrape.
helm install postgres-exporter prometheus-community/prometheus-postgres-exporter --namespace $NAMESPACE -f db-exporter-config.yaml

# Blackbox exporter is used to check the availability and latency of services.
helm install blackbox-exporter prometheus-community/prometheus-blackbox-exporter --namespace $NAMESPACE -f blackbox-config.yaml

# Verify the deployment
echo "Checking the status of the deployed services..."
kubectl get pods -n $NAMESPACE -o wide
echo

echo "All tasks completed successfully!"
echo

# To Uninstall the helm release
# # Uninstall Helm releases
# # helm list -n observability
# helm uninstall prometheus -n observability
# helm uninstall loki -n observability
# helm uninstall grafana -n observability

# Had error with prometheus server crashing and fix it with this link
# ==> https://github.com/helm/charts/issues/15742