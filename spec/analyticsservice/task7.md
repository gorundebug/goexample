# Task 7/17: `JoinedAnalyticsSink`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `custom-sink` |
| File | `analyticsservice/internal/functions/endpoint/joinedanalyticssink.go` |
| Service | `Analytics Service` |


## Behaviour

Validate and record the result of the two-way analytics join.




## Stream types
- Input: `AnalyticsResult` — `analyticsservice/internal/types/analyticsresult.go`
- Output: `AnalyticsResult` — `analyticsservice/internal/types/analyticsresult.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/endpoint/joinedanalyticssink.go` and preserve its generated contract
- [ ] Inspect input type `AnalyticsResult` in `analyticsservice/internal/types/analyticsresult.go`
- [ ] Inspect output type `AnalyticsResult` in `analyticsservice/internal/types/analyticsresult.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task7.md — JoinedAnalyticsSink — Go — done`