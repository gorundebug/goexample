# Task 8/8: `SoftDeadline`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `delay` |
| File | `orderservice/internal/functions/order/softdeadline.go` |
| Test | `orderservice/internal/functions/order/softdeadline_test.go` |
| Service | `Order Service` |


## Behaviour

Trigger the timeout branch shortly before the request deadline, leaving the configured duration to assemble a response.
When no request deadline exists, use the configured duration itself. Never wait past an existing deadline.





## Stream types
- Input: `Order` — `orderservice/internal/types/order.go`
- Output: `Order` — `orderservice/internal/types/order.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `orderservice/internal/functions/order/softdeadline.go` and preserve its generated contract
- [ ] Inspect input type `Order` in `orderservice/internal/types/order.go`
- [ ] Inspect output type `Order` in `orderservice/internal/types/order.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `orderservice/internal/functions/order/softdeadline_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] orderservice/task8.md — SoftDeadline — Go — done`