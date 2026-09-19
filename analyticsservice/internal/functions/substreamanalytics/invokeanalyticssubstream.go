package substreamanalytics

import (
	"context"

	"github.com/gorundebug/analyticsservice/internal/types"

	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/environment"
	"github.com/gorundebug/servicelib/transformation"
)

var _ transformation.MapFunction[*types.AnalyticsEvent, *types.AnalyticsResult] = (*InvokeAnalyticsSubstream)(nil)

// InvokeAnalyticsSubstream
type InvokeAnalyticsSubstream struct {
	substream runtime.SubStream[*types.AnalyticsEvent, *types.AnalyticsResult]
}

func NewInvokeAnalyticsSubstream(substream runtime.SubStream[*types.AnalyticsEvent, *types.AnalyticsResult]) *InvokeAnalyticsSubstream {
	return &InvokeAnalyticsSubstream{substream: substream}
}

func (f *InvokeAnalyticsSubstream) Map(ctx context.Context, _ runtime.Stream, value *types.AnalyticsEvent, out runtime.Collect[*types.AnalyticsResult]) {
	if err := f.substream.Consume(ctx, value, runtime.SubStreamCollectorFunc[*types.AnalyticsResult](func(resultCtx context.Context, result *types.AnalyticsResult) bool {
		out.Out(resultCtx, result)
		return true
	})); err != nil {
		panic(err)
	}
}

// MakeInvokeAnalyticsSubstream is instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeInvokeAnalyticsSubstream(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.MapStreamConfig) (*InvokeAnalyticsSubstream, error) {
	return &InvokeAnalyticsSubstream{}, nil
}
