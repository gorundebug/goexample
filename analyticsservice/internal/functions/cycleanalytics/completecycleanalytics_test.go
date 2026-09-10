package cycleanalytics

import (
	"context"
	"testing"

	"github.com/gorundebug/analyticsservice/internal/types"
	"github.com/stretchr/testify/assert"
)

// Keep the terminal analytics event once its cycle counter reaches three.
func TestCompleteCycleAnalytics_Filter(t *testing.T) {
	f := &CompleteCycleAnalytics{}
	assert.False(t, f.Filter(context.Background(), nil, &types.AnalyticsEvent{Value: 2}))
	assert.True(t, f.Filter(context.Background(), nil, &types.AnalyticsEvent{Value: 3}))
}
