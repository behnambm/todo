#!/bin/sh

make -C ./userservice build
make -C ./authservice build
make -C ./todoservice build

docker compose down
docker compose up --build
