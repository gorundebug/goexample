package automation

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gorundebug/servicelib/runtime"
)

// Preserve the result returned through the on-demand Activity endpoint.

func TestObserveActivityResult_Map(t *testing.T) {
	f := &ObserveActivityResult{}
	var collected []string
	out := runtime.CollectFunc[string](func(_ context.Context, v string) {
		collected = append(collected, v)
	})
	f.Map(context.Background(), nil, "activity:processed:job-1", out)
	assert.Equal(t, []string{"activity:processed:job-1"}, collected)
}
