# Task 2/8: `MapToOrderProcessed`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `map` |
| File | `orderservice/internal/functions/maptoorderprocessed.go` |
| Test | `orderservice/internal/functions/maptoorderprocessed_test.go` |
| Service | `Order Service` |


## Behaviour

Create an OrderProcessed event from the final order state.
Preserve the order ID, status, and processing time. Count all item results and reserved items; for unsuccessful orders use the final status as the failure reason.





## Stream types
- Input: `OrderState` — `orderservice/internal/types/orderstate.go`
- Output: `OrderProcessed` — `model/pkg/types/orderprocessed.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `orderservice/internal/functions/maptoorderprocessed.go` and preserve its generated contract
- [ ] Inspect input type `OrderState` in `orderservice/internal/types/orderstate.go`
- [ ] Inspect output type `OrderProcessed` in `model/pkg/types/orderprocessed.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `orderservice/internal/functions/maptoorderprocessed_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] orderservice/task2.md — MapToOrderProcessed — Go — done`