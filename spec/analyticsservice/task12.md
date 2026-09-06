# Task 12/17: `KeyPaymentsForJoin`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `keyBy` |
| File | `analyticsservice/internal/functions/joinanalytics/keypaymentsforjoin.go` |
| Test | `analyticsservice/internal/functions/joinanalytics/keypaymentsforjoin_test.go` |
| Service | `Analytics Service` |


## Behaviour

Key the payment analytics event by correlation key.




## Stream types
- Input: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`
- Output: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`
- Key: `AnalyticsKey`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/joinanalytics/keypaymentsforjoin.go` and preserve its generated contract
- [ ] Inspect input type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Inspect output type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `analyticsservice/internal/functions/joinanalytics/keypaymentsforjoin_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task12.md — KeyPaymentsForJoin — Go — done`