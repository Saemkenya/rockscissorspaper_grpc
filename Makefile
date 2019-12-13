M = $(shell printf "\033[34;1m▶\033[0m")

genProto: dep ; $(info $(M) Generating the protobuff codebase ...)
	protoc --proto_path=. --go_out=plugins=grpc:. proto/rps.proto

.PHONY: dep
