# Task 3/3: `OrderProcessedEndpointSource`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `kafka-source` |
| File | `analyticsservice/internal/functions/endpoint/orderprocessedendpointsource.go` |
| Service | `Analytics Service` |


## Behaviour

Exchange OrderProcessed events keyed by order ID.
Producers include the final status, processing time, total and confirmed item counts, and a failure reason for unsuccessful orders.
Consumers decode the event and mark its Kafka message processed only after the pipeline handles it successfully.





## Stream types
- Input: `OrderProcessed` — `model_go/pkg/types/orderprocessed.go`
- Output: `OrderProcessed` — `model_go/pkg/types/orderprocessed.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/endpoint/orderprocessedendpointsource.go` and preserve its generated contract
- [ ] Inspect input type `OrderProcessed` in `model_go/pkg/types/orderprocessed.go`
- [ ] Inspect output type `OrderProcessed` in `model_go/pkg/types/orderprocessed.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task3.md — OrderProcessedEndpointSource — Go — done`