#!/bin/bash

# Define the download URL and the destination path
URL="https://github.com/argoproj/argo-cd/releases/latest/download/argocd-linux-amd64"
DEST="/usr/local/bin/argocd"
HELM_REPO_NAME="argo"
HELM_REPO_URL="https://argoproj.github.io/argo-helm"
HELM_CHART_NAME="argo-cd"
NAMESPACE="argocd"
RELEASE_NAME="argocd"

# Function to check if Argo CD CLI is installed
check_argocd_installed() {
    if command -v argocd &> /dev/null; then
        echo "*****Argo CD CLI is already installed.*****"
        return 0
    else
        return 1
    fi
}

# Function to check if Argo CD is installed in the cluster
check_argocd_in_cluster() {
    if kubectl get namespaces | grep -q "$NAMESPACE"; then
        echo "*****Argo CD is already installed in the Kubernetes cluster.*****"
        echo
        return 0
    else
        return 1
    fi
}

# Check if Argo CD CLI is installed
check_argocd_installed
if [ $? -ne 0 ]; then
    # Download the Argo CD CLI binary
    echo "Downloading Argo CD CLI..."
    curl -sSL -o argocd-linux-amd64 "$URL" && \
    echo "*****Done Downloading ArgoCD CLI Binary.*****"
    echo

    # Make the binary executable
    echo "Setting executable permissions..."
    chmod +x argocd-linux-amd64
    echo "*****Done setting execution permission*****"
    echo

    # Move the binary to /usr/local/bin
    echo "Moving the binary to $DEST..."
    sudo mv argocd-linux-amd64 "$DEST"
    echo "*****Done moving argocd-linux-amd64 to the right destination*****"
    echo

    # Confirm installation
    if command -v argocd &> /dev/null; then
        echo "*****Argo CD CLI installed successfully!*****"
        echo
    else
        echo "*****Failed to install Argo CD CLI.*****"
        echo
        exit 1
    fi
else
    echo "*****Argo CD CLI is already installed. Skipping installation.*****"
    echo
fi
echo
# Check if Argo CD is installed in the cluster
check_argocd_in_cluster
if [ $? -ne 0 ]; then
    # Add the Argo Helm repository
    echo "Adding Argo Helm repository..."
    helm repo add "$HELM_REPO_NAME" "$HELM_REPO_URL" && helm repo update
    echo "*****Done adding ArgoCD helm repo*****"
    echo

    # Create argocd namespace
    echo "Creating namespace '$NAMESPACE'..."
    kubectl create namespace "$NAMESPACE"
    echo "*****Namespace Created*****"
    echo

    # Install Argo CD using Helm on a specified node using node-selector.yaml
    echo "Installing Argo CD with Helm..."
    helm install "$RELEASE_NAME" "$HELM_REPO_NAME/$HELM_CHART_NAME" --namespace "$NAMESPACE" -f node-selector.yaml
    echo "*****Done installing ArgoCD with helm*****"
    echo

    # Confirm Argo CD installation
    echo "Checking Argo CD installation..."
    kubectl get pods -n "$NAMESPACE"
    echo
else
    echo "*****Argo CD is already installed in the cluster. Skipping installation.*****"
    echo
fi

# List Helm repositories
echo "Listing Helm repositories..."
helm repo list
echo "*****Helm repositories listed.*****"
echo