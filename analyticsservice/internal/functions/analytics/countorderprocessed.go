package analytics

import (
	"context"

	"github.com/gorundebug/model_go/pkg/types"
	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/environment"
	"github.com/gorundebug/servicelib/runtime/environment/metrics"
	"github.com/gorundebug/servicelib/transformation"
)

var _ transformation.ProcessFunction[*types.OrderProcessed, *types.OrderProcessed, error] = (*CountOrderProcessed)(nil)

// CountOrderProcessed
type CountOrderProcessed struct {
	ordersTotal metrics.Int64CounterVec
}

func (f *CountOrderProcessed) Process(ctx context.Context, _ runtime.Stream, value *types.OrderProcessed, out runtime.Collect[*types.OrderProcessed], _ runtime.Collect[error]) {
	result := "unsuccessful"
	if value.Status == "CONFIRMED" {
		result = "successful"
	}
	f.ordersTotal.With(metrics.Labels{"result": result}).Inc(ctx)
	out.Out(ctx, value)
}

// MakeCountOrderProcessed is instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeCountOrderProcessed(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.ProcessStreamConfig) (*CountOrderProcessed, error) {
	counter, err := env.Metrics().Scope("analytics", nil).CounterVec(
		"orders_total", "Number of processed orders by result",
	)
	if err != nil {
		return nil, err
	}
	return &CountOrderProcessed{ordersTotal: counter}, nil
}
