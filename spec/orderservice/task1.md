# Task 1/8: `MapOrderItemResultToOrderState`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `map` |
| File | `orderservice/internal/functions/maporderitemresulttoorderstate.go` |
| Test | `orderservice/internal/functions/maporderitemresulttoorderstate_test.go` |
| Service | `Order Service` |


## Behaviour

Produce an order result containing one inventory result and preserving its order ID.
Mark it CONFIRMED when the item was reserved; otherwise mark it PARTIALLY_CONFIRMED.
Record the time when this result is produced.





## Stream types
- Input: `OrderItemResult` — `model/pkg/types/orderitemresult.go`
- Output: `OrderState` — `orderservice/internal/types/orderstate.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `orderservice/internal/functions/maporderitemresulttoorderstate.go` and preserve its generated contract
- [ ] Inspect input type `OrderItemResult` in `model/pkg/types/orderitemresult.go`
- [ ] Inspect output type `OrderState` in `orderservice/internal/types/orderstate.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `orderservice/internal/functions/maporderitemresulttoorderstate_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] orderservice/task1.md — MapOrderItemResultToOrderState — Go — done`