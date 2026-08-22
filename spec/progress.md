# Implementation Progress

Track completed tasks here. After implementing each task, append the **exact block** specified in that task's checklist — the format differs by task type:

**Pure transformation tasks** (Map, FlatMap, Filter, KeyBy, Join, Process, Delay):
```
- [x] taskN.md — FunctionName — done
```

**Endpoint tasks** (http-source, grpc-source) — include the example call:
````
- [x] taskN.md — FunctionName — done
  Example call:
  ```
  <command shown in the task checklist>
  ```
````

> **IMPORTANT:** Always copy the progress entry format from the task's own checklist, not from this file. Each task specifies exactly what to append.

---

<!-- completed tasks below -->
- [x] task2.md — ProcessOrderItems — done
- [x] task3.md — ProcessOrderItem — done
- [x] task4.md — MapOrderItemResultToOrderState — done
- [x] task5.md — SoftDeadline — done
- [x] task6.md — MapToOrderState — done
- [x] task1.md (orderservice) — ProcessOrder — done
  Example call (with real field values, not a template):
  ```
  curl -X POST http://localhost:9091/v1/processorder \
  -H "Content-Type: application/json" \
  -H "X-Request-ID: req-001" \
  -H "X-Trace: 1" \
  -d '{"customer_id":"cust-1","items":[{"item_id":"item-1","sku":"SKU-001","quantity":2,"unit_price":10.0}]}'
  ```
- [x] task1.md (inventoryservice) — ProcessOrderItem — done
  Example call (with real field values, not a template):
  ```
  grpcurl -plaintext \
  -H "x-trace: 1" \
  -d '{"order_id":"order-001","item_id":"item-1","sku":"SKU-001","quantity":2}' \
  localhost:9202 processorderitem.InventoryServiceApi/ProcessOrderItem
  ```
- [x] task2.md (inventoryservice) — GetInventoryItemData — done