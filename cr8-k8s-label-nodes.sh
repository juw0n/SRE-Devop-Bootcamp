#!/bin/bash

# Create Kubernetes Cluster (using Minikube)
minikube start --driver=kvm2 --nodes 3 -p sre-project

echo "Done creating the 3-Node cluster"
echo "**************************************"

# Apply the first set of labels
kubectl label node sre-project type=dependent-services-node
kubectl label node sre-project-m02 type=database-node
kubectl label node sre-project-m03 type=application-node

echo "Done applying label-1 to the 3-Node cluster"
echo "**************************************"

# Apply the second set of labels
kubectl label node sre-project service=dependent-services-node
kubectl label node sre-project-m02 database=database-node
kubectl label node sre-project-m03 application=application-node

echo "Done applying label-2 to the 3-Node cluster"
echo "**************************************"

# Create Namespaces
kubectl apply -f - <<EOF
# Hashicorp Vault Namespace
apiVersion: v1
kind: Namespace
metadata:
  name: vault-ns
  labels:
    name: vault-ns
---
# External Secrets Operator Namespace
apiVersion: v1
kind: Namespace
metadata:
  name: external-secrets-ns
  labels:
    name: external-secrets-ns
---
# Student API Namespace
apiVersion: v1
kind: Namespace
metadata:
  name: student-api-ns
  labels:
    name: student-api-ns
# ---
# # argocd
# apiVersion: v1
# kind: Namespace
# metadata:
#   name: argocd
#   labels:
#     name: argocd
---
# observability
apiVersion: v1
kind: Namespace
metadata:
  name: observability
  labels:
    name: observability
EOF

echo "Done creating namespaces"
echo "**************************************"
# Make the script executable ==> chmod +x nodes-label.sh
# Run the script: ==> ./nodes-label.sh

# PS: you can also taint the nodes. but this gave me some issues 
# kubectl taint nodes sre-project service=dependent-services-node:NoSchedule
# kubectl taint nodes sre-project-m02 database=database-node:NoSchedule
# kubectl taint nodes sre-project-m03 application=application-node:NoSchedule