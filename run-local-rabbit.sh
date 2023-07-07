#!/bin/bash

docker run --rm --hostname my-rabbit --name some-rabbit --network host rabbitmq:3
