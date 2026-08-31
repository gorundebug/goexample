#!/usr/bin/env bash
# clone.generated.sh — restore separately published Go services and modules.
# Other languages keep every component in the project repository.
#
# Usage: bash clone.generated.sh

set -euo pipefail

clone_if_missing() {
    local dir="$1"
    local repo="$2"
    local revision="$3"
    if [ -d "$dir" ]; then
        echo "  skip  $dir (already present)"
    else
        echo "  clone $repo@$revision → $dir"
        git clone --branch "$revision" --depth 1 "$repo" "$dir"
    fi
}

echo "==> Cloning services..."
clone_if_missing "analyticsservice" "https://github.com/gorundebug/analyticsservice.git" "v0.2.34"
clone_if_missing "automationservice" "https://github.com/gorundebug/automationservice.git" "v0.2.34"
clone_if_missing "inventoryservice" "https://github.com/gorundebug/inventoryservice.git" "v0.2.34"
clone_if_missing "orderservice" "https://github.com/gorundebug/orderservice.git" "v0.2.34"

echo "==> Cloning modules..."
clone_if_missing "inventory_service_api" "https://github.com/gorundebug/inventory_service_api.git" "v0.2.34"
clone_if_missing "model_go" "https://github.com/gorundebug/model_go.git" "v0.2.34"
clone_if_missing "order_service_api" "https://github.com/gorundebug/order_service_api.git" "v0.2.34"

echo "==> Done."