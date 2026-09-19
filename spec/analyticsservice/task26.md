# Task 26/26: `InvokeAnalyticsSubstream`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `map` |
| File | `analyticsservice/internal/functions/substreamanalytics/invokeanalyticssubstream.go` |
| Test | `analyticsservice/internal/functions/substreamanalytics/invokeanalyticssubstream_test.go` |
| Service | `Analytics Service` |


## Behaviour

Invoke the service-local analytics SubStream and emit its returned result.




## Stream types
- Input: `AnalyticsEvent` — `analyticsservice/internal/types/analyticsevent.go`
- Output: `AnalyticsResult` — `analyticsservice/internal/types/analyticsresult.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/substreamanalytics/invokeanalyticssubstream.go` and preserve its generated contract
- [ ] Inspect input type `AnalyticsEvent` in `analyticsservice/internal/types/analyticsevent.go`
- [ ] Inspect output type `AnalyticsResult` in `analyticsservice/internal/types/analyticsresult.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `analyticsservice/internal/functions/substreamanalytics/invokeanalyticssubstream_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task26.md — InvokeAnalyticsSubstream — Go — done`