# Task 8/26: `AnalyticsShipmentsSource`

> Rules: [`spec/rules.md`](../rules.md)

Paths and commands below are relative to the generated project root.

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `custom-source` |
| File | `analyticsservice/internal/functions/endpoint/analyticsshipmentssource.go` |
| Service | `Analytics Service` |


## Behaviour

Produce a deterministic shipment analytics event for the canonical multi-way join example.




## Stream types
- Emits into graph: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`
- Error: `error`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/endpoint/analyticsshipmentssource.go` and preserve its generated contract
- [ ] Inspect emitted message type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task8.md — AnalyticsShipmentsSource — Go — done`