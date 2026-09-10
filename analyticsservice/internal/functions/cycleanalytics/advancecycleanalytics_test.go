package cycleanalytics

import (
	"context"
	"testing"

	"github.com/gorundebug/analyticsservice/internal/types"
	"github.com/gorundebug/servicelib/runtime"
	"github.com/stretchr/testify/assert"
)

// Increment the cycle counter while preserving the analytics event identity.
func TestAdvanceCycleAnalytics_Map(t *testing.T) {
	f := &AdvanceCycleAnalytics{}
	var collected []*types.AnalyticsEvent
	out := runtime.CollectFunc[*types.AnalyticsEvent](func(_ context.Context, v *types.AnalyticsEvent) {
		collected = append(collected, v)
	})
	value := &types.AnalyticsEvent{Key: "cycle", Value: 1, Kind: "cycle"}
	f.Map(context.Background(), nil, value, out)
	if assert.Len(t, collected, 1) {
		assert.Equal(t, &types.AnalyticsEvent{Key: "cycle", Value: 2, Kind: "cycle"}, collected[0])
	}
	assert.Equal(t, 1, value.Value)
}
