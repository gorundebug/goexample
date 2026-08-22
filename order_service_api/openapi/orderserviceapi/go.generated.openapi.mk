.PHONY: all gen

all: gen

gen:
	@mkdir -p "$(MODULE_DIR)/pkg/generated/openapi/orderserviceapi"; \
	$(OAPI_CODEGEN) -generate "std-http,models" \
		-package orderserviceapi \
		-o "$(MODULE_DIR)/pkg/generated/openapi/orderserviceapi/orderserviceapi_http.openapi.go" \
		-import-mapping=./processorder/processorder.yaml:github.com/gorundebug/order_service_api/pkg/generated/openapi/orderserviceapi/processorder \
		orderserviceapi.generated.yaml; \
	cp "orderserviceapi.generated.yaml" "$(MODULE_DIR)/pkg/generated/openapi/orderserviceapi/openapi.yaml"; \
	mkdir -p "$(MODULE_DIR)/pkg/generated/openapi/orderserviceapi/processorder" && \
	cp "processorder/processorder.yaml" "$(MODULE_DIR)/pkg/generated/openapi/orderserviceapi/processorder/processorder.yaml"; \