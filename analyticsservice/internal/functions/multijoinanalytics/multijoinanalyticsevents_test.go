package multijoinanalytics

import (
	"context"
	"testing"

	"github.com/gorundebug/analyticsservice/internal/types"
	"github.com/gorundebug/servicelib/runtime"
	"github.com/stretchr/testify/assert"
)

// Combine matching order, payment, and shipment analytics events.

func TestMultiJoinAnalyticsEvents_MultiJoin(t *testing.T) {
	t.Skip("not yet implemented") // TODO: remove when implementation is ready
	f := &MultiJoinAnalyticsEvents{}
	var collected []*types.AnalyticsResult
	out := runtime.CollectFunc[*types.AnalyticsResult](func(_ context.Context, v *types.AnalyticsResult) {
		collected = append(collected, v)
	})
	var key string
	// TODO: populate key and values ([][]interface{} — one inner slice per joined stream)
	result := f.MultiJoin(context.Background(), nil, key, nil, out)
	_ = result
	assert.NotEmpty(t, collected)
}
