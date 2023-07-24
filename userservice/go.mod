module github.com/behnambm/todo/userservice

go 1.20

replace github.com/behnambm/todo/common => /home/behnam/go/src/github.com/behnambm/todo/common

require (
	github.com/behnambm/todo/common v0.0.0-00010101000000-000000000000
	github.com/mattn/go-sqlite3 v1.14.17
	github.com/streadway/amqp v1.1.0
	google.golang.org/grpc v1.56.2
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230720185612-659f7aaaa771 // indirect
)
