package cycleanalytics

import (
	"context"

	"github.com/gorundebug/analyticsservice/internal/types"

	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/environment"
	"github.com/gorundebug/servicelib/transformation"
)

var _ transformation.MapFunction[*types.AnalyticsEvent, *types.AnalyticsEvent] = (*AdvanceCycleAnalytics)(nil)

// AdvanceCycleAnalytics
type AdvanceCycleAnalytics struct{}

func (f *AdvanceCycleAnalytics) Map(ctx context.Context, _ runtime.Stream, value *types.AnalyticsEvent, out runtime.Collect[*types.AnalyticsEvent]) {
	next := *value
	next.Value++
	out.Out(ctx, &next)
}

// MakeAdvanceCycleAnalytics is instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeAdvanceCycleAnalytics(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.MapStreamConfig) (*AdvanceCycleAnalytics, error) {
	return &AdvanceCycleAnalytics{}, nil
}
