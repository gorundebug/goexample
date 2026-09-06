package joinanalytics

import (
	"context"
	"testing"

	"github.com/gorundebug/analyticsservice/internal/types"
	"github.com/gorundebug/servicelib/runtime"
	"github.com/gorundebug/servicelib/runtime/datastruct"
	"github.com/stretchr/testify/assert"
)

// Key the payment analytics event by correlation key.

func TestKeyPaymentsForJoin_KeyBy(t *testing.T) {
	t.Skip("not yet implemented") // TODO: remove when implementation is ready
	f := &KeyPaymentsForJoin{}
	var collected []datastruct.KeyValue[string, *types.AnalyticsEvent]
	out := runtime.CollectFunc[datastruct.KeyValue[string, *types.AnalyticsEvent]](func(_ context.Context, v datastruct.KeyValue[string, *types.AnalyticsEvent]) {
		collected = append(collected, v)
	})
	var value *types.AnalyticsEvent
	// TODO: populate value with meaningful test data
	f.KeyBy(context.Background(), nil, value, out)
	assert.NotEmpty(t, collected)
}
