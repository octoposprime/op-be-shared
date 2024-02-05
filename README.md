# op-be-shared
The Shared Libraries ( proto-buffs, tools, libs, etc. ) for the Backend Layer of OctopOSPrime

## Pre-Requirements
Grpc
```
export PATH="$PATH:$(go env GOPATH)/bin"

go get  google.golang.org/grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

make proto-generate
```