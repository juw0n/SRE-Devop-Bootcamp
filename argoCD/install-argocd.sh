#!/bin/bash

# Define the download URL and the destination path
URL="https://github.com/argoproj/argo-cd/releases/latest/download/argocd-linux-amd64"
DEST="/usr/local/bin/argocd"
HELM_REPO_NAME="argo"
HELM_REPO_URL="https://argoproj.github.io/argo-helm"
HELM_CHART_NAME="argo-cd"
NAMESPACE="argocd"
RELEASE_NAME="argocd"

# Check if Argo CD CLI is already installed
if command -v argocd &> /dev/null
then
    echo "Argo CD CLI is already installed."
else
    # If not installed, proceed with the installation
    echo "Argo CD CLI not found. Proceeding with installation..."

    # Download the Argo CD CLI binary
    echo "Downloading Argo CD CLI..."
    curl -sSL -o argocd-linux-amd64 "$URL"

    # Make the binary executable
    echo "Setting executable permissions..."
    chmod +x argocd-linux-amd64

    # Move the binary to /usr/local/bin
    echo "Moving the binary to $DEST..."
    sudo mv argocd-linux-amd64 "$DEST"

    # Confirm installation
    if command -v argocd &> /dev/null
    then
        echo "Argo CD CLI installed successfully!"
    else
        echo "Failed to install Argo CD CLI."
        exit 1
    fi
fi
echo
# Add the Argo Helm repository
echo "Adding Argo Helm repository..."
helm repo add "$HELM_REPO_NAME" "$HELM_REPO_URL" && \
helm repo update
echo
# create argdocd namespace
kubectl create namespace "$NAMESPACE"
echo
# Install Argo CD using Helm on a specified node using node-selector.yaml
echo "Installing Argo CD with Helm..."
helm install "$RELEASE_NAME" "$HELM_REPO_NAME/$HELM_CHART_NAME" \
  --namespace "$NAMESPACE" \
  -f node-selector.yaml \
  --set nodeSelector.service=dependent-services-node

# Confirm Argo CD installation
echo "Checking Argo CD installation..."
kubectl get pods -n "$NAMESPACE"

# Change the argocd-server service type to LoadBalancer OR NodePort:
echo "Change argo-server service type"
kubectl patch svc argocd-server -n "$NAMESPACE" -p '{"spec": {"type": "LoadBalancer"}}'
echo
echo "See argocd running services"
kubectl get svc -n "$NAMESPACE" -o wide
echo 
echo "All tasks completed successfully!"
echo