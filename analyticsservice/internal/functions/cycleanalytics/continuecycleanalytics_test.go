package cycleanalytics

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gorundebug/analyticsservice/internal/types"
)

// Keep intermediate analytics events whose cycle counter is below three.
func TestContinueCycleAnalytics_Filter(t *testing.T) {
	f := &ContinueCycleAnalytics{}
	assert.True(t, f.Filter(context.Background(), nil, &types.AnalyticsEvent{Value: 2}))
	assert.False(t, f.Filter(context.Background(), nil, &types.AnalyticsEvent{Value: 3}))
}
