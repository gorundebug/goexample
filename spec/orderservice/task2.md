# Task 2/8: `ProcessOrderItemSink`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `grpc-sink` |
| File | `orderservice/internal/functions/endpoint/processorderitemsink.go` |
| Service | `Order Service` |


## Behaviour

Reserve inventory for one order item using its order ID, item ID, SKU, and quantity.
Return the available quantity, reservation outcome, and status. The caller combines this response with the original identity, requested quantity, and unit price.
If the inventory call fails, the caller returns a non-reserved PROCESSING_ERROR result with the failure message.




## External contract

| Field | Value |
|-------|-------|
| Format | `proto` |
| Source | `inventory_service_api/proto/inventoryserviceapi/processorderitem/processorderitem.proto` |
| Request | `ProcessOrderItemRequest` |
| Response | `ProcessOrderItemResponse` |


## Stream types
- Input: `OrderItem` — `model/pkg/types/orderitem.go`
- Output: `OrderItemResult` — `model/pkg/types/orderitemresult.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `orderservice/internal/functions/endpoint/processorderitemsink.go` and preserve its generated contract
- [ ] Read `inventory_service_api/proto/inventoryserviceapi/processorderitem/processorderitem.proto`; change the source contract rather than generated bindings
- [ ] Inspect input type `OrderItem` in `model/pkg/types/orderitem.go`
- [ ] Inspect output type `OrderItemResult` in `model/pkg/types/orderitemresult.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Verify the endpoint/result lifecycle, including completion and error paths
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] orderservice/task2.md — ProcessOrderItemSink — Go — done`