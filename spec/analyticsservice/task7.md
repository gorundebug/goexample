# Task 7/26: `AnalyticsPaymentsSource`

> Rules: [`spec/rules.md`](../rules.md)

Paths and commands below are relative to the generated project root.

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `custom-source` |
| File | `analyticsservice/internal/functions/endpoint/analyticspaymentssource.go` |
| Service | `Analytics Service` |


## Behaviour

Produce a deterministic payment analytics event for the canonical join examples.




## Stream types
- Emits into graph: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`
- Error: `error`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/endpoint/analyticspaymentssource.go` and preserve its generated contract
- [ ] Inspect emitted message type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task7.md — AnalyticsPaymentsSource — Go — done`