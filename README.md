# go sample api

This project is used for creating a sample REST API using go and deploy it in local kubernetes cluster.

## Table of Contents
1. [Prerequisites](#prereq)
1. [Project Creation](#creation)
1. [Usage](#usage)

<a name="prereq"></a>
## Prerequisites

- [go](https://golang.org/) version 1.14+
- [docker](https://www.docker.com/products/docker-desktop)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
  

<a name="creation"></a>
## Project Creation

1) From your working directory clone the project from the repository.

2)Open the project in VSCode and the below command in terminal

```shell
docker build -t gosampleapi081320
```

3) Run the below kubectl commands to create the POD in the cluster

```shell
kubectl create -f "deployment.yaml"
```

4) Ensure the POD is running state by running the below commands

```shell
kubectl get pods
kubectl describe pod [podname]
kubectl logs [podname]
```
Note: Copy the POD name which will be used at a later point

5) Run the below command for the port forward so that incoming requests from localhost routs to the container running inside the POD

```shell
kubectl port-forward [podname] 9000:9000
```

6) Open the browser and hit the below urls

```shell
http://localhost:9000/
http://localhost:9000/articles/

```
