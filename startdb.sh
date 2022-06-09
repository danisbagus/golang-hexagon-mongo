#!/bin/bash

docker-compose up -d

sleep 5

docker exec mongors1 /scripts/rs-init.sh