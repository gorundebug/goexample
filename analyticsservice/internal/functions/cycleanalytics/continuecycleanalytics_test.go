package cycleanalytics

import (
	"context"
	"testing"

	"github.com/gorundebug/analyticsservice/internal/types"
	"github.com/stretchr/testify/assert"
)

// Keep intermediate analytics events whose cycle counter is below three.
func TestContinueCycleAnalytics_Filter(t *testing.T) {
	f := &ContinueCycleAnalytics{}
	assert.True(t, f.Filter(context.Background(), nil, &types.AnalyticsEvent{Value: 2}))
	assert.False(t, f.Filter(context.Background(), nil, &types.AnalyticsEvent{Value: 3}))
}
