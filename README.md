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

---

## About the test (to Argyle)

### Day one

Hi, I'm really happy with the test result, I've decided to use Golang for some reasons and I would like to
ponctuate them here:

- I have low experience with Golang yeat I've got excited with the oportunity to create a project using it.
- I knew and already created some Golang API, but only in personal projects or using larger tools (gorilla mux, Kong).
- To me, I thinks it's really cool to create something without depending a lot of third parties libaries and to use type enforcing
languages for backend development.

With that in mind, I'm finishing the core solution (Basic API + Kubernetes) within ~5 hours, but I think I would like
to enhance the system a little more. First I've created a basic recursive fibonacci resolution. Then I enhanced it adding 
memory, tring a little bit of Dynamic programming.

### Day two

The in memory resolution performs well, but since you mentioned that scalability and high calculus capability would be
important points of the resolution I've decied to include `Redis` as a fast in-memory storage, so diferent executions
could use previous calculations in order to give responses for higher N's with no problems, and also to enhance the
response time.

This implementation took ~ 2 hours from where I've stopped in the last day.


### Day three

To wrap up my project I've just refactored some repetitive parts off, replacing for better encapsulation functions,
also enhanced data types and this documentation!

This implementation took ~1 hour.

