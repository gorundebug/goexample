package substreamanalytics

import (
	"context"
	"testing"

	"github.com/gorundebug/analyticsservice/internal/types"
	"github.com/gorundebug/servicelib/runtime"
	"github.com/stretchr/testify/assert"
)

// Invoke the service-local analytics SubStream and emit its returned result.

func TestInvokeAnalyticsSubstream_Map(t *testing.T) {
	t.Skip("not yet implemented") // TODO: remove when implementation is ready
	f := &InvokeAnalyticsSubstream{}
	var collected []*types.AnalyticsResult
	out := runtime.CollectFunc[*types.AnalyticsResult](func(_ context.Context, v *types.AnalyticsResult) {
		collected = append(collected, v)
	})
	var value *types.AnalyticsEvent
	// TODO: populate value with meaningful test data
	f.Map(context.Background(), nil, value, out)
	assert.NotEmpty(t, collected)
}
