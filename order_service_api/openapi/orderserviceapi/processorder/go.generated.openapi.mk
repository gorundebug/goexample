.PHONY: all gen

all: gen

gen:
	@mkdir -p "$(MODULE_DIR)/pkg/generated/openapi/orderserviceapi/processorder"; \
	$(OAPI_CODEGEN) -generate "types,skip-prune" \
		-package processorder \
		-o "$(MODULE_DIR)/pkg/generated/openapi/orderserviceapi/processorder/processorder.go" \
		processorder.yaml