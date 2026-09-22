.PHONY: all gen

all: gen

gen:
	@set -e; mkdir -p "$(MODULE_DIR)/pkg/generated/proto/inventoryserviceapi/processorderitem"; \
	find . -maxdepth 1 -name "*.proto" | while read -r file; do \
		PATH="$(dir $(PROTOC)):$$PATH" $(PROTOC) -I="$(MODULE_DIR)" --go_out="$(MODULE_DIR)/pkg/generated" --go_opt=paths=source_relative "$(CURDIR)/$${file#./}"; \
	done;