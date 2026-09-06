# Task 6/17: `HighValueAnalyticsSink`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `custom-sink` |
| File | `analyticsservice/internal/functions/endpoint/highvalueanalyticssink.go` |
| Service | `Analytics Service` |


## Behaviour

Validate and record analytics results routed to the high-value Case branch.




## Stream types
- Input: `AnalyticsResult` — `analyticsservice/internal/types/analyticsresult.go`
- Output: `AnalyticsResult` — `analyticsservice/internal/types/analyticsresult.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/endpoint/highvalueanalyticssink.go` and preserve its generated contract
- [ ] Inspect input type `AnalyticsResult` in `analyticsservice/internal/types/analyticsresult.go`
- [ ] Inspect output type `AnalyticsResult` in `analyticsservice/internal/types/analyticsresult.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task6.md — HighValueAnalyticsSink — Go — done`