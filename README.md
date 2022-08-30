# GoMemoryLeak
Demo repo for Golang memory leak

## Pre-requisite
1. Install minikube: https://minikube.sigs.k8s.io/docs/start/
2. Install Tilt: https://docs.tilt.dev/install.html

## Run service
```bash
minikube addons enable metrics-server
minikbue start
tilt up
# In another terminal tab
minikube dashboard
```

