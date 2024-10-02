#!/bin/bash

# Check if Helm is installed
if ! command -v helm &> /dev/null
then
    echo "Helm is not installed. Please install Helm before running this script."
    exit 1
fi
echo 
# Create Vault Namespace if not exists
echo "Checking if vault-ns namespace exists..."
kubectl get namespace vault-ns &> /dev/null

if [ $? -ne 0 ]; then
  echo "Namespace vault-ns not found. Creating namespace..."
  kubectl create namespace vault-ns
else
  echo "Namespace vault-ns already exists."
fi
echo 
# Add Helm repository for HashiCorp
echo "Adding HashiCorp Helm repository..."
helm repo add hashicorp https://helm.releases.hashicorp.com
if [ $? -ne 0 ]; then
  echo "Failed to add HashiCorp Helm repository. Exiting."
  exit 1
fi
echo "*******Done adding Hashicorp Helm Repo*************"
echo 
# Update Helm repositories
echo "Updating Helm repositories..."
helm repo update
if [ $? -ne 0 ]; then
  echo "Failed to update Helm repositories. Exiting."
  exit 1
fi
echo "********Done updating Helm Repo************"
echo 
# Install Vault Helm chart
echo "Installing Vault Helm chart..."
helm install vault hashicorp/vault --namespace vault-ns -f vault-values.yaml
if [ $? -ne 0 ]; then
  echo "Failed to install Vault Helm chart. Exiting."
  exit 1
fi
echo "********Done Installing Hashicorp-Vault using Helm chart************"
echo 
# List Helm repositories
echo "Listing Helm repositories..."
helm repo list
echo "********************"
echo 
echo "Vault Helm chart installation complete."
echo
echo "See running vault pods"
kubectl get pods -n vault-ns -o wide
echo
echo "Vault services"
kubectl get svc -n vault-ns -o wide