package substreamanalytics

import (
	"context"

	"github.com/gorundebug/analyticsservice/internal/types"

	"github.com/gorundebug/servicelib/runtime"

	"github.com/gorundebug/servicelib/runtime/environment"
	"github.com/gorundebug/servicelib/transformation"
)

var _ transformation.MapFunction[*types.AnalyticsEvent, *types.AnalyticsResult] = (*BuildSubstreamAnalyticsResult)(nil)

// BuildSubstreamAnalyticsResult
type BuildSubstreamAnalyticsResult struct{}

func (f *BuildSubstreamAnalyticsResult) Map(ctx context.Context, _ runtime.Stream, value *types.AnalyticsEvent, out runtime.Collect[*types.AnalyticsResult]) {
	out.Out(ctx, &types.AnalyticsResult{
		Key:   value.Key,
		Total: value.Value * 2,
		Kind:  "substream",
	})
}

// MakeBuildSubstreamAnalyticsResult is instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeBuildSubstreamAnalyticsResult(ctx context.Context, env environment.ServiceEnvironment) (*BuildSubstreamAnalyticsResult, error) {
	return &BuildSubstreamAnalyticsResult{}, nil
}
