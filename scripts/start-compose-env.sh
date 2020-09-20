#!/usr/bin/env bash

sudo docker-compose down
sudo sh ./stop-docker-test-env.sh
sudo sh ./clear-docker-images.sh
sudo docker-compose up