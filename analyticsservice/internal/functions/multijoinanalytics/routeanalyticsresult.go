package multijoinanalytics

import (
	"context"
	"fmt"

	"github.com/gorundebug/analyticsservice/internal/types"

	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/environment"
	"github.com/gorundebug/servicelib/transformation"
)

var _ transformation.BuildSwitchFunction[*types.AnalyticsResult] = (*RouteAnalyticsResult)(nil)

// RouteAnalyticsResult
type RouteAnalyticsResult struct{}

func (f *RouteAnalyticsResult) BuildSwitch(_ runtime.Stream, whenItems []transformation.When) (func(*types.AnalyticsResult) int, error) {
	if len(whenItems) != 2 {
		return nil, fmt.Errorf("analytics result case requires exactly 2 branches, got %d", len(whenItems))
	}
	return func(value *types.AnalyticsResult) int {
		if value.Total >= 50 {
			return 0
		}
		return 1
	}, nil
}

// MakeRouteAnalyticsResult is instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeRouteAnalyticsResult(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.CaseStreamConfig) (*RouteAnalyticsResult, error) {
	return &RouteAnalyticsResult{}, nil
}
