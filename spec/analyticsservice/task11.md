# Task 11/17: `KeyOrdersForJoin`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `keyBy` |
| File | `analyticsservice/internal/functions/joinanalytics/keyordersforjoin.go` |
| Test | `analyticsservice/internal/functions/joinanalytics/keyordersforjoin_test.go` |
| Service | `Analytics Service` |


## Behaviour

Key the order analytics event by correlation key.




## Stream types
- Input: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`
- Output: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`
- Key: `AnalyticsKey`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/joinanalytics/keyordersforjoin.go` and preserve its generated contract
- [ ] Inspect input type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Inspect output type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `analyticsservice/internal/functions/joinanalytics/keyordersforjoin_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task11.md — KeyOrdersForJoin — Go — done`