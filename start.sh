#!/bin/sh

make -C ./userservice build
make -C ./authservice build
make -C ./todoservice build
make -C ./gatewayservice build

docker compose down
docker compose up --build
