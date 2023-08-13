Prerequisites:
1. Go plugins:
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    update path
    
    https://grpc.io/docs/languages/go/quickstart/#prerequisites

2. Protocol Buffer Compiler version 3
    follow instructions on https://grpc.io/docs/protoc-installation/ or https://github.com/protocolbuffers/protobuf/releases



Run Server:
$ go run greeter_server/main.go

Run Client:
$ go run greeter_client/main.go

CRUD Contracts:
1. update helloworld/helloworld.proto file
2. run protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto 
3. implement your changes on server and client programs