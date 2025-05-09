#!/bin/bash

# BEFORE RUNNING:

# 1. Install docker AND run the docker daemon
# 2. Install kubectl
# 3. Install kind 
# 4. Install helm

# All of the above can be installed with brew on macOS

set -e

for cmd in docker kubectl kind helm; do
  command -v $cmd >/dev/null 2>&1 || { echo "$cmd is not installed. Aborting."; exit 1; }
done

echo "Creating Kind cluster..."
kind create cluster --name manaflo-cluster

echo "Applying backend infrastructure..."
cd backend/infrastructure
kubectl apply -f .
cd ..

echo "Building and deploying user service..."
cd user-service
docker build -t user-service:latest .
kind load docker-image user-service:latest --name manaflo-cluster
cd infrastructure
kubectl apply -f .
cd ..
cd ..

echo "Building and deploying project service..."
cd project-service
docker build -t project-service:latest .
kind load docker-image project-service:latest --name manaflo-cluster
cd infrastructure
kubectl apply -f .
cd ..
cd ..

echo "Building and deploying task service..."
cd task-service
docker build -t task-service:latest .
kind load docker-image task-service:latest --name manaflo-cluster
cd infrastructure
kubectl apply -f .
cd ..
cd ..
cd ..

echo "Building and deploying frontend..."
cd frontend
docker build -t manaflo-frontend:latest .
kind load docker-image manaflo-frontend:latest --name manaflo-cluster
cd infrastructure
kubectl apply -f .

echo "Setting up NGINX Ingress Controller..."
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install ingress-nginx ingress-nginx/ingress-nginx

echo "Cluster setup complete. Listing pods:"
kubectl get pods --all-namespaces





