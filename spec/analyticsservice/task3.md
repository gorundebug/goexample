# Task 3/22: `AdvanceCycleAnalytics`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `map` |
| File | `analyticsservice/internal/functions/cycleanalytics/advancecycleanalytics.go` |
| Test | `analyticsservice/internal/functions/cycleanalytics/advancecycleanalytics_test.go` |
| Service | `Analytics Service` |


## Behaviour

Increment the cycle counter while preserving the analytics event identity.




## Stream types
- Input: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`
- Output: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/cycleanalytics/advancecycleanalytics.go` and preserve its generated contract
- [ ] Inspect input type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Inspect output type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `analyticsservice/internal/functions/cycleanalytics/advancecycleanalytics_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task3.md — AdvanceCycleAnalytics — Go — done`