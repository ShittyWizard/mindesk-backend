#!/usr/bin/env bash

# run mongo
docker run --name local-mongodb -p 27017:27017 -d mongo