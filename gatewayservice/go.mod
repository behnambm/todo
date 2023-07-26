module github.com/behnambm/todo/gatewayservice

go 1.20

replace github.com/behnambm/todo/common => /home/behnam/go/src/github.com/behnambm/todo/common

require (
	github.com/behnambm/todo/common v0.0.0-00010101000000-000000000000
	github.com/labstack/echo/v4 v4.11.1
	github.com/mattn/go-sqlite3 v1.14.17
	github.com/streadway/amqp v1.1.0
	google.golang.org/grpc v1.56.2
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
)
