# Task 2/2: `GetInventoryItemData`

> Rules: [`spec/rules.md`](../rules.md)

| Field | Value |
|-------|-------|
| Kind | `process` |
| File | `inventoryservice/internal/functions/getinventoryitemdata.go` |
| Test | `inventoryservice/internal/functions/getinventoryitemdata_test.go` |
| Service | `Inventory Service` |
| Visibility | private |

---

## Description

Reserve the requested quantity without allowing concurrent orders to overdraw stock.
On success, return CONFIRMED with the requested quantity available. Otherwise return OUT_OF_STOCK with the current available quantity.
Preserve the order and item identity, requested quantity, and unit price.
The example starts with SKU-001: 100, SKU-002: 50, and SKU-003: 25.

---

## Internal stream types (T / R / K)
### Input type (T)

| Field | Value |
|-------|-------|
| Name | `OrderItemResult` |
| Kind | `native` |
| File | `model/pkg/types/orderitemresult.go` |
| Import | `github.com/gorundebug/model/pkg/types` |
### Output type (R)

| Field | Value |
|-------|-------|
| Name | `OrderItem` |
| Kind | `native` |
| File | `model/pkg/types/orderitem.go` |
| Import | `github.com/gorundebug/model/pkg/types` |

---

## Implementation checklist

- [ ] Read [`spec/rules.md`](rules.md)
- [ ] Open `inventoryservice/internal/functions/getinventoryitemdata.go`
- [ ] Understand input type `OrderItemResult` in `model/pkg/types/orderitemresult.go`
- [ ] If `OrderItemResult` struct body is empty, add its fields from the type description before implementing
- [ ] Understand output type `OrderItem` in `model/pkg/types/orderitem.go`
- [ ] If `OrderItem` struct body is empty, add its fields from the type description before implementing
- [ ] Implement the function body (replace `//TODO` stub)
- [ ] Open `inventoryservice/internal/functions/getinventoryitemdata_test.go` and implement the test cases
- [ ] Run `make test` — all tests must pass
- [ ] Append to `spec/progress.md`: `- [x] task2.md — GetInventoryItemData — done`
