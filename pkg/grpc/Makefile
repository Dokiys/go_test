PROTO_FILES=$(shell find ./helloproto/* -name *.proto)
proto:
	protoc --proto_path=. --descriptor_set_out=helloproto/descriptor.desc --go_out=../ $(PROTO_FILES) google/protobuf/timestamp.proto

gofast_proto:
	protoc --proto_path=. --gofast_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:../ $(PROTO_FILES)

# GRPC
PROTO_GRPC_FILES=$(shell find ./hellogrpc/* -name *.proto)
proto_grpc:
	protoc --go_out=. --go_opt=paths=source_relative \
 		    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
 		    $(PROTO_GRPC_FILES)

#proto:
#	protoc --proto_path=. \
#           --gogofaster_out=plugins=grpc2,paths=source_relative:. \
#           --go-http_out=paths=source_relative:. \
#           $(PROTO_FILES)