# Task 1/2: `CountOrderProcessed`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `process` |
| File | `analyticsservice/internal/functions/analytics/countorderprocessed.go` |
| Test | `analyticsservice/internal/functions/analytics/countorderprocessed_test.go` |
| Service | `Analytics Service` |


## Behaviour

Count successful and unsuccessful orders independently, then return the event unchanged.





## Stream types
- Input: `OrderProcessed` — `model/pkg/types/orderprocessed.go`
- Output: `OrderProcessed` — `model/pkg/types/orderprocessed.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/analytics/countorderprocessed.go` and preserve its generated contract
- [ ] Inspect input type `OrderProcessed` in `model/pkg/types/orderprocessed.go`
- [ ] Inspect output type `OrderProcessed` in `model/pkg/types/orderprocessed.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `analyticsservice/internal/functions/analytics/countorderprocessed_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task1.md — CountOrderProcessed — Go — done`