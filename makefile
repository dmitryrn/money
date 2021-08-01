.PHONY: proto
proto:
	cd proto; protoc --go_out=. --go_opt=paths=source_relative base.proto
