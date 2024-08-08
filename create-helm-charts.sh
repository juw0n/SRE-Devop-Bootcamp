#!/bin/bash

# Create base directories
mkdir -p helm/charts/api
mkdir -p helm/charts/db
mkdir -p helm/charts/vault

# Create Helm charts
helm create helm/charts/api
helm create helm/charts/db
helm create helm/charts/vault

# Remove files in each templates/ directory
rm -rf helm/charts/api/templates/*
rm -rf helm/charts/db/templates/*
rm -rf helm/charts/vault/templates/*

# Clear content in values.yaml files
echo -n > helm/charts/api/values.yaml
echo -n > helm/charts/db/values.yaml
echo -n > helm/charts/vault/values.yaml

# Print message
echo "Helm charts structure created and cleaned up successfully."