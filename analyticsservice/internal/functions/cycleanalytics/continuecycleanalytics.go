package cycleanalytics

import (
	"context"

	"github.com/gorundebug/analyticsservice/internal/types"

	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/environment"
	"github.com/gorundebug/servicelib/transformation"
)

var _ transformation.FilterFunction[*types.AnalyticsEvent] = (*ContinueCycleAnalytics)(nil)

// ContinueCycleAnalytics
type ContinueCycleAnalytics struct{}

func (f *ContinueCycleAnalytics) Filter(_ context.Context, _ runtime.Stream, value *types.AnalyticsEvent) bool {
	return value.Value < 3
}

// MakeContinueCycleAnalytics is instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeContinueCycleAnalytics(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.FilterStreamConfig) (*ContinueCycleAnalytics, error) {
	return &ContinueCycleAnalytics{}, nil
}
