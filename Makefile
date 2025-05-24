PROTO_DIR := proto
OUT_DIR   := gen

PROTO_FILES := $(shell find $(PROTO_DIR) -name '*.proto')

GO_OPTS   := paths=source_relative
GRPC_OPTS := paths=source_relative,require_unimplemented_servers=false

.PHONY: proto clean tools

proto:
	@echo "🛠  Generating Go + gRPC code from .proto files…"
	@mkdir -p $(OUT_DIR)
	@protoc \
	  -I=$(PROTO_DIR) \
	  --go_out=$(OUT_DIR)   --go_opt=$(GO_OPTS) \
	  --go-grpc_out=$(OUT_DIR) --go-grpc_opt=$(GRPC_OPTS) \
	  $(PROTO_FILES)
	@echo "✅  Done."

clean:
	@echo "🧹  Cleaning generated stubs…"
	@rm -rf $(OUT_DIR)
	@echo "✅  Clean."

tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

