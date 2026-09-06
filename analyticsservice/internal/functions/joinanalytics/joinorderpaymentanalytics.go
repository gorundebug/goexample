package joinanalytics

import (
	"context"

	"github.com/gorundebug/analyticsservice/internal/types"

	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/environment"
	"github.com/gorundebug/servicelib/transformation"
)

var _ transformation.JoinFunction[string, *types.AnalyticsEvent, *types.AnalyticsEvent, *types.AnalyticsResult] = (*JoinOrderPaymentAnalytics)(nil)

// JoinOrderPaymentAnalytics
type JoinOrderPaymentAnalytics struct{}

func (f *JoinOrderPaymentAnalytics) Join(ctx context.Context, _ runtime.Stream, key string, leftValue []*types.AnalyticsEvent, rightValue []*types.AnalyticsEvent, out runtime.Collect[*types.AnalyticsResult]) bool {
	if len(leftValue) == 0 || len(rightValue) == 0 {
		return false
	}
	out.Out(ctx, &types.AnalyticsResult{Key: key, Total: leftValue[0].Value + rightValue[0].Value, Kind: "join"})
	return true
}

// MakeJoinOrderPaymentAnalytics is instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeJoinOrderPaymentAnalytics(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.JoinStreamConfig) (*JoinOrderPaymentAnalytics, error) {
	return &JoinOrderPaymentAnalytics{}, nil
}
