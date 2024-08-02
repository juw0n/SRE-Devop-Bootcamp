#!/bin/bash

# Add Helm repository for HashiCorp
echo "Adding HashiCorp Helm repository..."
helm repo add hashicorp https://helm.releases.hashicorp.com
echo "********************"

# Update Helm repositories
echo "Updating Helm repositories..."
helm repo update
echo "********************"

# List Helm repositories
echo "Listing Helm repositories..."
helm repo list
echo "********************"

# Search for Vault chart in the HashiCorp repo
echo "Searching for Vault chart in the HashiCorp repository..."
helm search repo hashicorp/vault
echo "********************"

# Install Vault Helm chart
echo "Installing Vault Helm chart..."
helm install vault hashicorp/vault --namespace vault-ns -f vault-values.yaml
echo "********************"

echo "Vault Helm chart installation complete."
echo "********************"