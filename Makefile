generate:
	protoc --proto_path=proto proto/bank.proto --go_out=proto/ --go-grpc_out=proto/
	protoc --proto_path=proto proto/authentication.proto --go_out=proto/ --go-grpc_out=proto/