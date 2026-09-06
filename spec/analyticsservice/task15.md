# Task 15/17: `KeyShipmentsForMultiJoin`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `keyBy` |
| File | `analyticsservice/internal/functions/multijoinanalytics/keyshipmentsformultijoin.go` |
| Test | `analyticsservice/internal/functions/multijoinanalytics/keyshipmentsformultijoin_test.go` |
| Service | `Analytics Service` |


## Behaviour

Key the shipment analytics event for the multi-way join.




## Stream types
- Input: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`
- Output: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`
- Key: `AnalyticsKey`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/multijoinanalytics/keyshipmentsformultijoin.go` and preserve its generated contract
- [ ] Inspect input type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Inspect output type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `analyticsservice/internal/functions/multijoinanalytics/keyshipmentsformultijoin_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task15.md — KeyShipmentsForMultiJoin — Go — done`