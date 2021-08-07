JS_OUT_DIR = "../web/src/proto"

.PHONY: proto
proto:
	mkdir -p internal/proto web/src/proto
	cd proto; protoc \
		--go_out=../internal/proto/ \
		--go_opt=paths=source_relative \
		--go-grpc_out=../internal/proto/ \
		--go-grpc_opt=paths=source_relative \
		--plugin="protoc-gen-ts=../web/node_modules/.bin/protoc-gen-ts" \
		--js_out="import_style=commonjs,binary:${JS_OUT_DIR}" \
		--ts_out="service=grpc-web:${JS_OUT_DIR}" \
		base.proto
