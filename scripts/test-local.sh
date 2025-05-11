#!/bin/bash

set -e

docker run --rm \
    -d \
    --name sweetrpg-db-mongodb-test \
    -p 27017:27017 \
    -e MONGO_INITDB_ROOT_USERNAME=admin \
    -e MONGO_INITDB_ROOT_PASSWORD=admin \
    -e MONGO_INITDB_DATABASE=sweetrpg-db \
    mongodb/mongodb-community-server:8.0.8-ubi8
sleep 10

export TEST_DB_URI="mongodb://admin:admin@localhost:27017/?authSource=admin"

cleanup() {
    echo "Cleaning up..."
    docker stop sweetrpg-db-mongodb-test
}
trap cleanup EXIT

go test -v ./... \
    -timeout 30m \
    -coverprofile=coverage.out \
    -covermode=atomic
