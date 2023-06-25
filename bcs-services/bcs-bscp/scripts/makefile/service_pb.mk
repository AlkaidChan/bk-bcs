# Currently, Makefile is only compiled for proto files under pkg/protocol/xx_service/.
# If you adjust the file directory, Makefile will have a problem.

PROTO=$(wildcard ./*.proto)
VERSION=$(shell protoc --version)
DIR_NAME=$(notdir $(CURDIR))

OBJ:=$(patsubst %.proto, %.pb.go, $(PROTO))
GRPC:=$(patsubst %.proto, %_grpc.pb.go, $(PROTO))
GW:=$(patsubst %.proto, %.pb.gw.go, $(PROTO))

all:
    ifeq ("$(VERSION)","libprotoc 22.0")
		@protoc --proto_path=. --proto_path=../../../../ --proto_path=../../../pkg/thirdparty/protobuf/ --go_opt=paths=source_relative --go_out=. --go-grpc_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:. --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative --grpc-gateway_out=. *.proto
    else
		@echo "make pb failed, protoc version not 22.0"
		exit 1
    endif

clean:
	@rm -f $(OBJ) $(GW) $(GRPC)

py:
	mkdir -p ../../../pkg/protocol/python/$(DIR_NAME)
	@echo $(DIR_NAME)
	python3 -m grpc_tools.protoc --proto_path=. --proto_path=../../../../ --proto_path=../../../pkg/thirdparty/protobuf/ --python_out=../../../pkg/protocol/python/$(DIR_NAME) --pyi_out=../../../pkg/protocol/python/$(DIR_NAME) --grpc_python_out=../../../pkg/protocol/python/$(DIR_NAME) *.proto
