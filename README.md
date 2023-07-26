# Todo app

This app will provide todo functionalities in HTTP API format


### API doc

- `Postman online doc:` [Link](https://documenter.getpostman.com/view/7250373/2s946pYoJy)
- `Postman collection:` provided in the root of the project, it can be used after being imported to Postman


### Services 

- `Gateway:` This service acts as the entry point for the network, accepting and processing HTTP API requests. 
It utilizes `gRPC` or `Event` mechanisms for various operations.
- `User:` The User service is responsible for managing user-related functionalities. It offers support for both `gRPC` and `Event` protocols for different operations.
- `Auth:` This service is used to handle authentication and token management which supports `gRPC` 
- `Todo:` The Todo service is designed for managing todos and their associated items. It provides support for both `gRPC` and `Event` protocols.


### Run 

First you need to clone the repo and pull these images:

```bash
git clone https://github.com/behnambm/todo.git
```
```bash
docker pull rabbitmq:3.12-management
```
```bash
docker pull golang:1.20
```

then run this script to run the project(it uses `docker compose`):

```bash
chmod +x ./start.sh
```
```bash
./start.sh
```



### Test

Unfortunately because of the limited time there are no unit tests but for `gateway` service some tests are provided that
only can be used **after all the services are up and running** and the test will be made against real services.

To run the gateway service tests, navigate to the `gatewayservice` directory and set the environment variable for the listen port (ensure it matches the one defined in the docker-compose file):
```bash
cd gatewayservice
```

```bash
export HTTP_LISTEN_URL=http://localhost:2020
```

```bash
go test server/httpserver/server_test.go
```


