.PHONY: all gen

all: gen

gen:
	@mkdir -p "$(MODULE_DIR)/pkg/generated/proto/inventoryserviceapi"; \
	mkdir -p "$(MODULE_DIR)/../descriptors"; \
	find . -maxdepth 1 -name "*.proto" | while read -r file; do \
		PATH="$(dir $(PROTOC)):$$PATH" $(PROTOC) -I=. -I="$(MODULE_DIR)" --go_out="$(MODULE_DIR)/pkg/generated/proto/inventoryserviceapi" --go_opt=paths=source_relative --go-grpc_out="$(MODULE_DIR)/pkg/generated/proto/inventoryserviceapi" --go-grpc_opt=paths=source_relative $$file; \
		PATH="$(dir $(PROTOC)):$$PATH" $(PROTOC) -I=. -I="$(MODULE_DIR)" --descriptor_set_out="$(MODULE_DIR)/../descriptors/inventoryserviceapi.protoset" --include_imports $$file; \
	done;