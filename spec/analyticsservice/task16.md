# Task 16/17: `MultiJoinAnalyticsEvents`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `multiJoin` |
| File | `analyticsservice/internal/functions/multijoinanalytics/multijoinanalyticsevents.go` |
| Test | `analyticsservice/internal/functions/multijoinanalytics/multijoinanalyticsevents_test.go` |
| Service | `Analytics Service` |


## Behaviour

Combine matching order, payment, and shipment analytics events.




## Stream types
- Input: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`
- Output: `AnalyticsResult` — `analyticsservice/internal/types/analyticsresult.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/multijoinanalytics/multijoinanalyticsevents.go` and preserve its generated contract
- [ ] Inspect input type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Inspect output type `AnalyticsResult` in `analyticsservice/internal/types/analyticsresult.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `analyticsservice/internal/functions/multijoinanalytics/multijoinanalyticsevents_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task16.md — MultiJoinAnalyticsEvents — Go — done`