# Task 5/8: `ProcessOrder`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `http-source` |
| File | `orderservice/internal/functions/processorder.go` |
| Service | `Order Service` |


## Behaviour

Accept orders with at least one item and positive quantities; reject malformed or invalid requests as client errors.
Reuse X-Request-ID when supplied, otherwise generate an order ID. Preserve customer, item, price, and X-Trace data, and apply the configured timeout of five seconds by default.
Return one response per order. When all items finish, use CONFIRMED only if every item was reserved; otherwise use PARTIALLY_CONFIRMED. If the deadline wins, return TIMED_OUT with the item results received so far.
Calculate the total from processed item prices, falling back to the submitted total when no item result arrived, and include individual item failures in the response.




## External contract

| Field | Value |
|-------|-------|
| Format | `openapi` |
| Source | `order_service_api/openapi/orderserviceapi/processorder/processorder.yaml` |
| Request | `ProcessOrderRequest` |
| Response | `ProcessOrderResponse` |


## Stream types
- Input: `Order` — `orderservice/internal/types/order.go`
- Output: `OrderState` — `orderservice/internal/types/orderstate.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `orderservice/internal/functions/processorder.go` and preserve its generated contract
- [ ] Read `order_service_api/openapi/orderserviceapi/processorder/processorder.yaml`; change the source contract rather than generated bindings
- [ ] Inspect input type `Order` in `orderservice/internal/types/order.go`
- [ ] Inspect output type `OrderState` in `orderservice/internal/types/orderstate.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Verify the endpoint/result lifecycle, including completion and error paths
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] orderservice/task5.md — ProcessOrder — Go — done`