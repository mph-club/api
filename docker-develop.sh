#!/bin/bash

# This script starts the app and uses realize to
# reload local tasks, theres no build

set -e

eval `aws ecr get-login --region us-east-1 --no-include-email`
docker-compose rm -s -f
docker-compose up -d --build --force-recreate
