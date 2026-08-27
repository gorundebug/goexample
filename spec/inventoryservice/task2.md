# Task 2/2: `GetInventoryItemData`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Language | `Go` |
| Kind | `process` |
| File | `inventoryservice/internal/functions/inventoryItem/getinventoryitemdata.go` |
| Test | `inventoryservice/internal/functions/inventoryItem/getinventoryitemdata_test.go` |
| Service | `Inventory Service` |


## Behaviour

Reserve the requested quantity without allowing concurrent orders to overdraw stock.
On success, return CONFIRMED with the requested quantity available. Otherwise return OUT_OF_STOCK with the current available quantity.
Preserve the order and item identity, requested quantity, and unit price.
The example starts with SKU-001: 100, SKU-002: 50, and SKU-003: 25.





## Stream types
- Input: `OrderItem` — `model/pkg/types/orderitem.go`
- Output: `OrderItemResult` — `model/pkg/types/orderitemresult.go`

## Checklist

- [ ] Read [`spec/rules.md`](../rules.md), especially the `Go` section
- [ ] Open `inventoryservice/internal/functions/inventoryItem/getinventoryitemdata.go` and preserve its generated contract
- [ ] Inspect input type `OrderItem` in `model/pkg/types/orderitem.go`
- [ ] Inspect output type `OrderItemResult` in `model/pkg/types/orderitemresult.go`
- [ ] Implement the Go function and propagate the received `context.Context`
- [ ] Run `make test`
- [ ] Implement meaningful assertions in `inventoryservice/internal/functions/inventoryItem/getinventoryitemdata_test.go`
- [ ] Re-read this checklist
- [ ] Append to `spec/progress.md`: `- [x] inventoryservice/task2.md — GetInventoryItemData — Go — done`