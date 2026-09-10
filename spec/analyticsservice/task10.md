# Task 10/22: `CycleAnalyticsResultSink`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `custom-sink` |
| File | `analyticsservice/internal/functions/endpoint/cycleanalyticsresultsink.go` |
| Service | `Analytics Service` |


## Behaviour

Validate the terminal event emitted after three passes through the feedback cycle.




## Stream types
- Input: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`
- Output: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/endpoint/cycleanalyticsresultsink.go` and preserve its generated contract
- [ ] Inspect input type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Inspect output type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task10.md — CycleAnalyticsResultSink — Go — done`