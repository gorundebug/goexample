package joinanalytics

import (
	"context"
	"testing"

	"github.com/gorundebug/analyticsservice/internal/types"
	"github.com/gorundebug/servicelib/runtime"
	"github.com/stretchr/testify/assert"
)

// Join matching order and payment analytics events and emit their combined total.

func TestJoinOrderPaymentAnalytics_Join(t *testing.T) {
	t.Skip("not yet implemented") // TODO: remove when implementation is ready
	f := &JoinOrderPaymentAnalytics{}
	var collected []*types.AnalyticsResult
	out := runtime.CollectFunc[*types.AnalyticsResult](func(_ context.Context, v *types.AnalyticsResult) {
		collected = append(collected, v)
	})
	var key string
	var leftValues []*types.AnalyticsEvent
	var rightValues []*types.AnalyticsEvent
	// TODO: populate key, leftValues, rightValues with meaningful test data
	result := f.Join(context.Background(), nil, key, leftValues, rightValues, out)
	_ = result
	assert.NotEmpty(t, collected)
}
