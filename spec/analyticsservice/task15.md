# Task 15/26: `SubstreamAnalyticsInputSource`

> Rules: [`spec/rules.md`](../rules.md)

Paths and commands below are relative to the generated project root.

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `custom-source` |
| File | `analyticsservice/internal/functions/endpoint/substreamanalyticsinputsource.go` |
| Service | `Analytics Service` |


## Behaviour

Produce one deterministic analytics event that invokes the service-local SubStream example.




## Stream types
- Emits into graph: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`
- Error: `error`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/endpoint/substreamanalyticsinputsource.go` and preserve its generated contract
- [ ] Inspect emitted message type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task15.md — SubstreamAnalyticsInputSource — Go — done`