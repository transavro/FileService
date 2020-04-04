generate:
	protoc -I proto/ -I${GOPATH}/src --go_out=plugins=grpc:proto proto/fileps.proto
	protoc -I proto/ -I${GOPATH}/src --grpc-gateway_out=logtostderr=true:proto proto/fileps.proto
	protoc -I proto/ -I${GOPATH}/src --swagger_out=logtostderr=true:proto proto/fileps.proto
	protoc -I proto/ -I${GOPATH}/src --govalidators_out=logtostderr=true:proto proto/fileps.proto

install:
	go build . && ./UserService
