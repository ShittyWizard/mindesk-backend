#!/usr/bin/env bash

bash stop-docker-test-env.sh
docker build -t mindesk .
docker-compose up