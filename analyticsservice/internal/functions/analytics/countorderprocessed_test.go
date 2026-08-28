package analytics

import (
	"context"
	"testing"

	"github.com/gorundebug/model_go/pkg/types"
	"github.com/gorundebug/servicelib/runtime"
	"github.com/stretchr/testify/assert"
)

// Count successful and unsuccessful orders independently, then return the event unchanged.

func TestCountOrderProcessed_Process(t *testing.T) {
	t.Skip("not yet implemented") // TODO: remove when implementation is ready
	f := &CountOrderProcessed{}
	var results []*types.OrderProcessed
	out := runtime.CollectFunc[*types.OrderProcessed](func(_ context.Context, v *types.OrderProcessed) {
		results = append(results, v)
	})
	var errs []error
	rout := runtime.CollectFunc[error](func(_ context.Context, v error) {
		errs = append(errs, v)
	})
	var value *types.OrderProcessed
	// TODO: populate value with meaningful test data
	f.Process(context.Background(), nil, value, out, rout)
	assert.Empty(t, errs)
	assert.NotEmpty(t, results)
}
