# Fibonacci API

A fibonacci calculator API written in Golang.

## Kubernetes

This project is meant to be executed by a kubernetes cluster, all the configurations are under `k8s` folder, the API deployment
is linked to the `Dockerfile` present in the root directory.

### Running the cluster

To run this project locally you will need:

- Minikube installed
- Docker installed


### Setting up and running

To start the project simply run:

```bash
# Start your local cluster
minukube start

# Deploy the api
kubectl create -f k8s/api/deployment.yaml
```


**Check if the project was deployed**
```bash
kubectl get deployments
```

expected output:

```text
NAME      READY   UP-TO-DATE   AVAILABLE   AGE
fib-api   1/1     1            1           58s
```
