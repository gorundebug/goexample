# Task 13/26: `OrderProcessedEndpointSource`

> Rules: [`spec/rules.md`](../rules.md)

Paths and commands below are relative to the generated project root.

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
- Emits into graph: `OrderProcessed` — `model_go/pkg/types/orderprocessed.go`
- Receives graph result: `OrderProcessed` — `model_go/pkg/types/orderprocessed.go`
- Error: `error`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `analyticsservice/internal/functions/endpoint/orderprocessedendpointsource.go` and preserve its generated contract
- [ ] Inspect emitted message type `OrderProcessed` in `model_go/pkg/types/orderprocessed.go`
- [ ] Inspect graph result type `OrderProcessed` in `model_go/pkg/types/orderprocessed.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] analyticsservice/task13.md — OrderProcessedEndpointSource — Go — done`