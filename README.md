# Fibonacci API

A fibonacci calculator API written in Golang. This API uses Redis and Golang to
achieve high Fib numbers with fast responses as well.

Stack:
- Golang
    - net/http
    - redis v8
- Redis 

---

## Running locally

Since this is a conteinerized aplication there's a `docker-compose.yaml` that will:

- Setup a local Redis
- Setup and run the api code based on the `Dockerfile`

To run the project simply run:

```bash
docker-compose up
```

Using (Postman)[https://www.postman.com/] or `Curl` you can call the api by:

```bash
# Check if the API is UP and running
curl localhost:8000/health-check

# Calculate fibonacci sequence for fib(5)
curl localhost:8000/fib?n=5
```

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
