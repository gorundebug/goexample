package automation

import (
	"context"

	"github.com/gorundebug/servicelib/runtime"

	"github.com/gorundebug/servicelib/runtime/environment"
	"github.com/gorundebug/servicelib/transformation"
)

var _ transformation.MapFunction[string, string] = (*ObserveFanoutActivityB)(nil)

// ObserveFanoutActivityB
type ObserveFanoutActivityB struct{}

func (f *ObserveFanoutActivityB) Map(ctx context.Context, _ runtime.Stream, value string, out runtime.Collect[string]) {
	out.Out(ctx, value)
}

// MakeObserveFanoutActivityB is instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeObserveFanoutActivityB(ctx context.Context, env environment.ServiceEnvironment) (*ObserveFanoutActivityB, error) {
	return &ObserveFanoutActivityB{}, nil
}
