#!/bin/bash

version=$(go run ./cmd/rabbitping -version | awk '{ print $2 }' | awk -F= '{ print $2 }')

echo version=$version

docker build --no-cache \
    -t udhos/rabbitping:latest \
    -t udhos/rabbitping:$version \
    -f docker/Dockerfile .

echo "push: docker push udhos/rabbitping:$version; docker push udhos/rabbitping:latest"
