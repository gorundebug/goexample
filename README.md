# Example

## Services

- [Analytics Service](https://github.com/gorundebug/analyticsservice/blob/main/README.md)
- [Inventory Service](https://github.com/gorundebug/inventoryservice/blob/main/README.md)
- [Order Service](https://github.com/gorundebug/orderservice/blob/main/README.md)

## Modules

- [inventory_service_api](https://github.com/gorundebug/inventory_service_api/blob/main/README.md)
- [model](https://github.com/gorundebug/model/blob/main/README.md)
- [order_service_api](https://github.com/gorundebug/order_service_api/blob/main/README.md)

## Local development

```bash
make tools        # install buf, golangci-lint, act (once)
make build        # build all services
make run          # run all services locally
make test         # run tests
make lint         # run linter
```

## Local CI with act

Runs GitHub Actions locally in Docker — uses `go.work` to resolve local modules,
Athens to cache public packages. No external registry needed during development.

### First-time setup (once)

```bash
make docker-up      # start Athens, Prometheus, Grafana and all services
```

### Run CI

```bash
make act            # run CI for all services
```

### Publish modules and services to GitHub

Each module and service lives in its own GitHub repository. Push all at once or individually:

```bash
make git-push                    # push everything (requires: make tools && gh auth login)
make git-push-inventory_service_api      # push module inventory_service_api → github.com/gorundebug/inventory_service_api
make git-push-model      # push module model → github.com/gorundebug/model
make git-push-order_service_api      # push module order_service_api → github.com/gorundebug/order_service_api
make git-push-inventoryservice   # push service Inventory Service → github.com/gorundebug/inventoryservice
make git-push-orderservice   # push service Order Service → github.com/gorundebug/orderservice
```

Prerequisites: run `make tools` once to install `gh` CLI, then authenticate if needed:

```bash
./tools/gh auth login   # select GitHub.com → HTTPS → token
```

The token needs the **`repo`** scope (required to create private repositories).


## Docker

```bash
make docker-up      # build images and start all services
make docker-up RUNTIME_IMAGE=1 # start minimal runtime images without sources or build tools
make docker-down    # stop all services
make docker-restart # restart all services
make docker-clean   # stop services and remove all volumes (wipes persistent data)
make debug-inventoryservice  # rebuild and restart Inventory Service in debug mode (Delve on port 2345)
make debug-orderservice  # rebuild and restart Order Service in debug mode (Delve on port 2346)
```

`RUNTIME_IMAGE=1` selects the final multi-stage runtime target. Benchmark and
profiling tools select it automatically. The default remains the development
layout with debugger/source mounts where the language supports them.

## Optional order analytics through Kafka

The shared `orderProcessed` Kafka endpoint is disabled in Order Service by
default, so it completes without creating a producer. To publish order results
to Redpanda, set this in `orderservice/config/overrides.yaml`:

```yaml
endpoints:
  orderProcessed:
    enabled: true
```

`analyticsservice` consumes the `order-processed` topic and counts successful
and unsuccessful orders. The compose workspace includes its Redpanda broker.

## Call the Order Service

After `make docker-up`, submit an order that can be reserved from the initial
inventory:

```bash
curl --fail-with-body \
  -X POST http://localhost:9091/v1/processorder \
  -H 'Content-Type: application/json' \
  -d '{
    "customer_id": "customer-1",
    "items": [
      {
        "item_id": "item-1",
        "sku": "SKU-001",
        "quantity": 2,
        "unit_price": 799.0
      }
    ]
  }'
```

The response has order status `CONFIRMED`; its item has `reserved: true` and
status `CONFIRMED`. Initial inventory is `SKU-001: 100`, `SKU-002: 50`, and
`SKU-003: 25`. Successful requests reduce that inventory until the Inventory
Service is restarted.
