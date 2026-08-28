# Task 2/3: `AnalyticsScheduleSource`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `schedule-source` |
| File | `analyticsservice/internal/functions/cron/analyticsschedulesource.go` |
| Service | `Analytics Service` |


## Behaviour

Create an analytics job message identifying the local scheduled firing.





## Stream types
- Input: `AutomationJob`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/cron/analyticsschedulesource.go` and preserve its generated contract
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task2.md — AnalyticsScheduleSource — Go — done`