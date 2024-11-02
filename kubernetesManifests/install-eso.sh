#!/bin/bash

# Check if Helm is installed
if ! command -v helm &> /dev/null; then
  echo "Helm is not installed. Please install Helm and try again."
  exit 1
fi
echo "*****Helm is installed.*****"
echo

# Create namespace if not exists
NAMESPACE="external-secrets-ns"
if ! kubectl get namespace $NAMESPACE > /dev/null 2>&1; then
  echo "Namespace $NAMESPACE not found. Creating..."
  kubectl create namespace $NAMESPACE
  echo "*****Namespace $NAMESPACE created.*****"
fi

# Add External Secrets Helm repository
echo "Adding External Secrets Helm repository..."
helm repo add external-secrets https://charts.external-secrets.io
echo "*****Helm repository added.*****"
echo

# Update Helm repositories
echo "Updating Helm repositories..."
helm repo update
echo "*****Helm repositories updated.*****"
echo

# Install External Secrets Operator (ESO)
echo "Installing External Secrets Operator..."
helm install external-secrets external-secrets/external-secrets \
  --namespace external-secrets-ns \
  --set nodeSelector.service=dependent-services-node \
  -f eso-node-values.yaml
echo "*****External Secrets Operator installed.*****"
echo

# List Helm repositories
echo "Listing Helm repositories..."
helm repo list
echo "*****Helm repositories listed.*****"
echo
echo "External-secrets Helm chart installation complete."
echo
echo "See running external-secrets pods"
kubectl get pods -n external-secrets-ns -o wide
echo
echo "See external-secrets services"
kubectl get svc -n external-secrets-ns -o wide