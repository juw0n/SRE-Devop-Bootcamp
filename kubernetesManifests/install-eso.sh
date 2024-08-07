#!/bin/bash

# Exit on error
set -e

# Add External Secrets Helm repository
echo "Adding External Secrets Helm repository..."
helm repo add external-secrets https://charts.external-secrets.io
echo "Helm repository added."
echo "********************"

# Update Helm repositories
echo "Updating Helm repositories..."
helm repo update
echo "Helm repositories updated."
echo "********************"

# Install External Secrets Operator (ESO)
echo "Installing External Secrets Operator..."
helm install external-secrets external-secrets/external-secrets \
  --namespace external-secrets-ns \
  -f eso-node-values.yaml
echo "External Secrets Operator installed."
echo "********************"

# List Helm repositories
echo "Listing Helm repositories..."
helm repo list
echo "Helm repositories listed."
echo "********************"


# Search for ESO chart in External Secrets repository
echo "Searching for ESO chart in External Secrets repository..."
helm search repo external-secrets/external-secrets
echo "Chart search complete."
echo "********************"