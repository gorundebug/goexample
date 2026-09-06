package multijoinanalytics

import (
	"context"

	"github.com/gorundebug/analyticsservice/internal/types"

	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/environment"
	"github.com/gorundebug/servicelib/transformation"
)

var _ transformation.MultiJoinFunction[string, any, *types.AnalyticsResult] = (*MultiJoinAnalyticsEvents)(nil)

// MultiJoinAnalyticsEvents
type MultiJoinAnalyticsEvents struct{}

func (f *MultiJoinAnalyticsEvents) MultiJoin(ctx context.Context, _ runtime.Stream, key string, values [][]interface{}, out runtime.Collect[*types.AnalyticsResult]) bool {
	if len(values) != 3 || len(values[0]) == 0 || len(values[1]) == 0 || len(values[2]) == 0 {
		return false
	}
	order, orderOK := values[0][0].(*types.AnalyticsEvent)
	payment, paymentOK := values[1][0].(*types.AnalyticsEvent)
	shipment, shipmentOK := values[2][0].(*types.AnalyticsEvent)
	if !orderOK || !paymentOK || !shipmentOK {
		return false
	}
	out.Out(ctx, &types.AnalyticsResult{Key: key, Total: order.Value + payment.Value + shipment.Value, Kind: "multi"})
	return true
}

// MakeMultiJoinAnalyticsEvents is instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeMultiJoinAnalyticsEvents(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.MultiJoinStreamConfig) (*MultiJoinAnalyticsEvents, error) {
	return &MultiJoinAnalyticsEvents{}, nil
}
