#!/bin/bash

DOCKER_IMAGE="kitanoyoru/kafka-keda-example"
DOCKER_IMAGE_TAG="latest"
SCHEMA_PATHS="/usr/local/lib/schema/avro/user.avsc"
BOOTSTRAP_SERVERS="kafka:29092"
TOPIC="kita.user"
CONSUMER_GROUP_ID="user-consumer-group"

export DOCKER_IMAGE
export DOCKER_IMAGE_TAG
export SCHEMA_PATHS
export BOOTSTRAP_SERVERS
export TOPIC
export CONSUMER_GROUP_ID

docker compose -f ./deployments/docker-compose.yaml up -d

