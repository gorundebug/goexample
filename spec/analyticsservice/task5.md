# Task 5/22: `ContinueCycleAnalytics`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `filter` |
| File | `analyticsservice/internal/functions/cycleanalytics/continuecycleanalytics.go` |
| Test | `analyticsservice/internal/functions/cycleanalytics/continuecycleanalytics_test.go` |
| Service | `Analytics Service` |


## Behaviour

Keep intermediate analytics events whose cycle counter is below three.




## Stream types
- Input: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`
- Output: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/cycleanalytics/continuecycleanalytics.go` and preserve its generated contract
- [ ] Inspect input type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Inspect output type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `analyticsservice/internal/functions/cycleanalytics/continuecycleanalytics_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task5.md — ContinueCycleAnalytics — Go — done`