.PHONY: proto
proto:
	mkdir -p back/proto
	cd proto; protoc --go_out=../back/proto/ --go_opt=paths=source_relative --go-grpc_out=../back/proto/ --go-grpc_opt=paths=source_relative base.proto
