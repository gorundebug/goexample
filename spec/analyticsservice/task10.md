# Task 10/17: `JoinOrderPaymentAnalytics`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `join` |
| File | `analyticsservice/internal/functions/joinanalytics/joinorderpaymentanalytics.go` |
| Test | `analyticsservice/internal/functions/joinanalytics/joinorderpaymentanalytics_test.go` |
| Service | `Analytics Service` |


## Behaviour

Join matching order and payment analytics events and emit their combined total.




## Stream types
- Input: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`
- Output: `AnalyticsResult` — `analyticsservice/internal/types/analyticsresult.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/joinanalytics/joinorderpaymentanalytics.go` and preserve its generated contract
- [ ] Inspect input type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Inspect output type `AnalyticsResult` in `analyticsservice/internal/types/analyticsresult.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `analyticsservice/internal/functions/joinanalytics/joinorderpaymentanalytics_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task10.md — JoinOrderPaymentAnalytics — Go — done`