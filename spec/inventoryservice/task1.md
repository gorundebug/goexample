# Task 1/2: `ProcessOrderItem`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Kind | `grpc-source` |
| File | `inventoryservice/internal/functions/processorderitem.go` |
| Service | `Inventory Service` |
| Visibility | private |

---

## Description

Reserve inventory for one order item using its order ID, item ID, SKU, and quantity.
Return the available quantity, reservation outcome, and status. The caller combines this response with the original identity, requested quantity, and unit price.
If the inventory call fails, the caller returns a non-reserved PROCESSING_ERROR result with the failure message.

---

## API contract (ReqT / ResR)

Kind: **proto**

| Field | Value |
|-------|-------|
| File | `inventory_service_api/proto/inventoryserviceapi/processorderitem/processorderitem.proto` |
| Import | `github.com/gorundebug/inventory_service_api/pkg/generated/proto/inventoryserviceapi/processorderitem` |
| ReqT | `ProcessOrderItemRequest` |
| ResR | `ProcessOrderItemResponse` |

> **The contract file may contain empty message/schema stubs with no fields.**
> Before implementing this function, open `inventory_service_api/proto/inventoryserviceapi/processorderitem/processorderitem.proto` and verify that
> `ProcessOrderItemRequest` and `ProcessOrderItemResponse` have all
> necessary fields. If they are empty, add the fields and regenerate:
> ```
> cd example/inventory_service_api/proto/inventoryserviceapi/processorderitem && make gen
> ```
> Do NOT edit the generated Go types in `github.com/gorundebug/inventory_service_api/pkg/generated/proto/inventoryserviceapi/processorderitem` directly.

---

## Internal stream types (T / R / K)
### Input type (T)

| Field | Value |
|-------|-------|
| Name | `OrderItem` |
| Kind | `native` |
| File | `model/pkg/types/orderitem.go` |
| Import | `github.com/gorundebug/model/pkg/types` |
### Output type (R)

| Field | Value |
|-------|-------|
| Name | `OrderItemResult` |
| Kind | `native` |
| File | `model/pkg/types/orderitemresult.go` |
| Import | `github.com/gorundebug/model/pkg/types` |

---

## Implementation checklist

- [ ] Read [`spec/rules.md`](rules.md)
- [ ] Open `inventoryservice/internal/functions/processorderitem.go`
- [ ] Read contract: `inventory_service_api/proto/inventoryserviceapi/processorderitem/processorderitem.proto`
- [ ] Understand input type `OrderItem` in `model/pkg/types/orderitem.go`
- [ ] If `OrderItem` struct body is empty, add its fields from the type description before implementing
- [ ] Understand output type `OrderItemResult` in `model/pkg/types/orderitemresult.go`
- [ ] If `OrderItemResult` struct body is empty, add its fields from the type description before implementing
- [ ] Check that `ProcessOrderItemRequest` and `ProcessOrderItemResponse` have all necessary fields; if empty — extend `inventory_service_api/proto/inventoryserviceapi/processorderitem/processorderitem.proto` and run `make gen` in `inventory_service_api/proto/inventoryserviceapi/processorderitem`
- [ ] Implement the function body (replace `//TODO` stub)
- [ ] Run `make test` — all tests must pass
- [ ] Verify the endpoint works and append to `spec/progress.md` — replace the hint below with a **real call containing actual field values**:
  ````
  - [x] task1.md — ProcessOrderItem — done
  Example call (with real field values, not a template):
  ```
  grpcurl -plaintext \
  -H "x-trace: 1" \
  -d '{}' \
  localhost:9202 processorderitem.InventoryServiceApi/ProcessOrderItem
  ```
  ````
