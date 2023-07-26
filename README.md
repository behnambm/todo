# Todo app

This app will provide todo functionalities in HTTP API format


### API doc

- `Postman online doc: ` [Link](https://documenter.getpostman.com/view/7250373/2s946pYoJy)
- `Postman collection: ` provided in the root of the project, it can be used after being imported to Postman


### Services 

- `Gateway: ` This service is used to provide HTTP API and it's used as the entry point of the network. 
It accepts the requests and uses `gRPC` or `Event` for different operations.
- `User: ` This service is used to manage the users and it supports both `gRPC` and `Event` for different operations
- `Auth: ` This service is used to handle authentication and token management which suppports `gRPC` 
- `Todo: ` This service is used for managing the todos and items and it supports both `gRPC` and `Event` 


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

```bash
cd gatewayservice
```

If you change this env the listen port for the `gateway` service should be changed in docker compose file accordingly.
```bash
export HTTP_LISTEN_URL=http://localhost:2020
```

```bash
go test server/httpserver/server_test.go
```


