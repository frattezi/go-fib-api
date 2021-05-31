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

---

## About the test (to Argyle)

Hi, I'm really happy with the test result, I've decided to use Golang for some reasons and I would like to
ponctuate them here:

- I have low experience with Golang yeat I've got excited with the oportunity to create a project using it.
- I knew and already made some Golang API, but only in personal projects or using larger tools (gorilla mux, Kong).
- To me, I thinks it's really cool to create something without depending a lot of third parties.

With that in mind, I'm finishing the core solution (Basic API + Kubernetes) within ~5 hours, but I think I would like
to enhance the system a little more.

