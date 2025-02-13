proto:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	rm -rf definitions/*
	@protoc --proto_path=proto --go_out=definitions --go_opt=paths=source_relative \
    --go-grpc_out=definitions --go-grpc_opt=paths=source_relative \
    proto/client/*.proto proto/runner/*.proto

.PHONY: proto