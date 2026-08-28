# Task 7/8: `ProcessOrderItems`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `flatMap` |
| File | `orderservice/internal/functions/order/processorderitems.go` |
| Test | `orderservice/internal/functions/order/processorderitems_test.go` |
| Service | `Order Service` |


## Behaviour

Emit every order item independently for inventory processing.
Preserve each item's data and assign the parent order ID.





## Stream types
- Input: `Order` — `orderservice/internal/types/order.go`
- Output: `OrderItem` — `model_go/pkg/types/orderitem.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `orderservice/internal/functions/order/processorderitems.go` and preserve its generated contract
- [ ] Inspect input type `Order` in `orderservice/internal/types/order.go`
- [ ] Inspect output type `OrderItem` in `model_go/pkg/types/orderitem.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `orderservice/internal/functions/order/processorderitems_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] orderservice/task7.md — ProcessOrderItems — Go — done`