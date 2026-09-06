# Task 17/17: `RouteAnalyticsResult`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `case` |
| File | `analyticsservice/internal/functions/multijoinanalytics/routeanalyticsresult.go` |
| Test | `analyticsservice/internal/functions/multijoinanalytics/routeanalyticsresult_test.go` |
| Service | `Analytics Service` |


## Behaviour

Route high-value analytics results to the first branch and all others to the second branch.




## Stream types
- Input: `AnalyticsResult` — `analyticsservice/internal/types/analyticsresult.go`
- Output: `AnalyticsResult` — `analyticsservice/internal/types/analyticsresult.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/multijoinanalytics/routeanalyticsresult.go` and preserve its generated contract
- [ ] Inspect input type `AnalyticsResult` in `analyticsservice/internal/types/analyticsresult.go`
- [ ] Inspect output type `AnalyticsResult` in `analyticsservice/internal/types/analyticsresult.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `analyticsservice/internal/functions/multijoinanalytics/routeanalyticsresult_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task17.md — RouteAnalyticsResult — Go — done`