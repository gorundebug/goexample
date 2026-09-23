# Task 3/3: `GetInventoryItemError`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `map` |
| File | `inventoryservice/internal/functions/inventoryItem/getinventoryitemerror.go` |
| Test | `inventoryservice/internal/functions/inventoryItem/getinventoryitemerror_test.go` |
| Service | `Inventory Service` |


## Behaviour

When inventory processing fails, return an OUT_OF_STOCK result with no available quantity.
Preserve the order and item identity and requested quantity, and record the failure.





## Stream types
- Input: `InventoryFailure` — `inventoryservice/internal/types/inventoryfailure.go`
- Output: `OrderItemResult` — `model_go/pkg/types/orderitemresult.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `inventoryservice/internal/functions/inventoryItem/getinventoryitemerror.go` and preserve its generated contract
- [ ] Inspect input type `InventoryFailure` in `inventoryservice/internal/types/inventoryfailure.go`
- [ ] Inspect output type `OrderItemResult` in `model_go/pkg/types/orderitemresult.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `inventoryservice/internal/functions/inventoryItem/getinventoryitemerror_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] inventoryservice/task3.md — GetInventoryItemError — Go — done`