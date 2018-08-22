#!/bin/bash

# This script starts:
# 1. Postgres ("the intake tables")
# 2. golang container

set -e

cd docker

echo "Building MPH Parent"
cd ..
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ~/Projects/mphclub-golang-api/bin/mphclub-golang-api
cd docker

# Copy the golang binary into the docker folder
cp ../bin/mphclub-golang-api go-app/mphclub-golang-api

# Copy static swagger index.html into the docker folder
cp ../swagger/redoc-static.html go-app/swagger/redoc-static.html

#eval `aws ecr get-login --region us-east-1 --no-include-email`
docker-compose rm -s -f
docker-compose up -d --build --force-recreate

#sleep 10
