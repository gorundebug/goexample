.PHONY: all gen

all: gen

gen:
	@mkdir -p "$(MODULE_DIR)/pkg/generated/proto/inventoryserviceapi/processorderitem"; \
	find . -maxdepth 1 -name "*.proto" | while read -r file; do \
		PATH="$(dir $(PROTOC)):$$PATH" $(PROTOC) -I=. -I="$(MODULE_DIR)" --go_out="$(MODULE_DIR)/pkg/generated/proto/inventoryserviceapi/processorderitem" --go_opt=paths=source_relative $$file; \
	done;