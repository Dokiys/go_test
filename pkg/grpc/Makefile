PROTO_FILES=$(shell find ./helloproto/* -name *.proto)
.PHONY: proto
proto:
	protoc --proto_path=. --descriptor_set_out=helloproto/descriptor.desc --go_out=../ $(PROTO_FILES) google/protobuf/timestamp.proto

.PHONY: gofast_proto
gofast_proto:
	protoc --proto_path=. --gofast_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:../ $(PROTO_FILES)

# GRPC
PROTO_GRPC_FILES=$(shell find ./hellogrpc/* -name *.proto)
.PHONY: proto_grpc
proto_grpc:
	protoc --go_out=. --go_opt=paths=source_relative \
 		    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
 		    $(PROTO_GRPC_FILES)

# gRPC Gateway
PROTO_GRPC_FILES=$(shell find ./gateway/* -name *.proto)
.PHONY: proto_gateway
proto_gateway:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out ./ \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt paths=source_relative \
        --grpc-gateway_opt generate_unbound_methods=true \
        $(PROTO_GRPC_FILES)

# protoc-gen-my
.PHONY: protoc-gen-my
protoc-gen-my:
	protoc --proto_path=./protoc-gen-my --my_out=version=1.1,paths=source_relative:./protoc-gen-my/gen ./protoc-gen-my/gen/hellogrpc.proto