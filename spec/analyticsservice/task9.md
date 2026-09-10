# Task 9/22: `CycleAnalyticsInputSource`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `custom-source` |
| File | `analyticsservice/internal/functions/endpoint/cycleanalyticsinputsource.go` |
| Service | `Analytics Service` |


## Behaviour

Produce one deterministic analytics event that exercises the finite feedback cycle.




## Stream types
- Input: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/endpoint/cycleanalyticsinputsource.go` and preserve its generated contract
- [ ] Inspect input type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task9.md — CycleAnalyticsInputSource — Go — done`