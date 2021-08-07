JS_OUT_DIR = "../front/src/proto"

.PHONY: proto
proto:
	mkdir -p back/proto front/src/proto
	cd proto; protoc \
		--go_out=../back/proto/ \
		--go_opt=paths=source_relative \
		--go-grpc_out=../back/proto/ \
		--go-grpc_opt=paths=source_relative \
		--plugin="protoc-gen-ts=../front/node_modules/.bin/protoc-gen-ts" \
		--js_out="import_style=commonjs,binary:${JS_OUT_DIR}" \
		--ts_out="service=grpc-web:${JS_OUT_DIR}" \
		base.proto
