package cycleanalytics

import (
	"context"

	"github.com/gorundebug/analyticsservice/internal/types"

	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/environment"
	"github.com/gorundebug/servicelib/transformation"
)

var _ transformation.FilterFunction[*types.AnalyticsEvent] = (*CompleteCycleAnalytics)(nil)

// CompleteCycleAnalytics
type CompleteCycleAnalytics struct{}

func (f *CompleteCycleAnalytics) Filter(_ context.Context, _ runtime.Stream, value *types.AnalyticsEvent) bool {
	return value.Value >= 3
}

// MakeCompleteCycleAnalytics is instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeCompleteCycleAnalytics(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.FilterStreamConfig) (*CompleteCycleAnalytics, error) {
	return &CompleteCycleAnalytics{}, nil
}
