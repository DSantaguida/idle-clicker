generate:
	protoc --proto_path=proto proto/${target}.proto --go_out=proto/ --go-grpc_out=proto/