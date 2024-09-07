
proto: 
	protoc --proto_path=person --go_out=./person --go_opt=paths=source_relative --go-grpc_out=./person --go-grpc_opt=paths=source_relative person.proto

run:
	go run server/main.go

install: 
	go mod tidy