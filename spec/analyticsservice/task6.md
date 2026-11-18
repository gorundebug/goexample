# Task 6/26: `AnalyticsOrdersSource`

> Rules: [`spec/rules.md`](../rules.md)

Paths and commands below are relative to the generated project root.

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `custom-source` |
| File | `analyticsservice/internal/functions/endpoint/analyticsorderssource.go` |
| Service | `Analytics Service` |


## Behaviour

Produce a deterministic order analytics event for the canonical join examples.




## Stream types
- Emits into graph: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`
- Error: `error`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/endpoint/analyticsorderssource.go` and preserve its generated contract
- [ ] Inspect emitted message type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`

- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task6.md — AnalyticsOrdersSource — Go — done`