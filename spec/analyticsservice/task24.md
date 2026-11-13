# Task 24/26: `RouteAnalyticsResult`

> Rules: [`spec/rules.md`](../rules.md)

Paths and commands below are relative to the generated project root.

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

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/multijoinanalytics/routeanalyticsresult.go` and preserve its generated contract
- [ ] Inspect input type `AnalyticsResult` in `analyticsservice/internal/types/analyticsresult.go`
- [ ] Implement `BuildSwitch` and its value-only selector; neither receives a per-message `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `analyticsservice/internal/functions/multijoinanalytics/routeanalyticsresult_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task24.md — RouteAnalyticsResult — Go — done`