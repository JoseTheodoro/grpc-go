## grpc-go

##### install protoc 
https://grpc.io/docs/protoc-installation/

##### install proto-go-gen
``go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
  && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest``

##### update path environment
``export PATH="$PATH:$HOME/.local/bin:$(go env GOPATH)/bin"``

##### run command in your terminal
``make install``

``make proto``

``make run``
