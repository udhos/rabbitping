[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/udhos/rabbitping/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/udhos/rabbitping)](https://goreportcard.com/report/github.com/udhos/rabbitping)
[![Go Reference](https://pkg.go.dev/badge/github.com/udhos/rabbitping.svg)](https://pkg.go.dev/github.com/udhos/rabbitping)
[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/rabbitping)](https://artifacthub.io/packages/search?repo=rabbitping)
[![Docker Pulls rabbitping](https://img.shields.io/docker/pulls/udhos/rabbitping)](https://hub.docker.com/r/udhos/rabbitping)

# rabbitping

rabbitping

# Env vars

```
export AMQP_URL=amqp://guest:guest@rabbitmq:5672/
export INTERVAL=10s
export TIMEOUT=5s
export FAILURE_THRESHOLD=6
export RESTART_DEPLOY=my-miniapi # change this to the name of the deployment that must be restarted upon failure
export RESTART_NAMESPACE=default
export METRICS_ADDR=:3000
export METRICS_PATH=/metrics
export METRICS_NAMESPACE=""
export METRICS_BUCKETS_LATENCY="0.0005, 0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, .5, 1, 2.5, 5"
export HEALTH_ADDR=:8888
export HEALTH_PATH=/health
```

# Rabbitmq

Run locally:

```
docker run --rm --hostname my-rabbit --name some-rabbit --network host rabbitmq:3
```

Deploy into kubernetes:

```
kubectl create deploy rabbitmq --image=rabbitmq:3 --port 5672

kubectl expose deploy rabbitmq
```

# Docker

Docker hub:

https://hub.docker.com/r/udhos/rabbitping

Run from docker hub:

```
docker run -p 8080:8080 --rm udhos/rabbitping:0.1.0
```

Build recipe:

```
./docker/build.sh

docker push udhos/rabbitping:0.1.0
```

# Helm chart

You can use the provided helm charts to install rabbitping in your Kubernetes cluster.

See https://udhos.github.io/rabbitping/

## Lint

    helm lint ./charts/rabbitping --values charts/rabbitping/values.yaml

## Debug

    helm template ./charts/rabbitping --values charts/rabbitping/values.yaml --debug

## Render at server

    helm install my-rabbitping ./charts/rabbitping --values charts/rabbitping/values.yaml --dry-run

## Install

    helm install my-rabbitping ./charts/rabbitping --values charts/rabbitping/values.yaml

    helm list -A
