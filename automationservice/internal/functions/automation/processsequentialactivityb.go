package automation

import (
	"context"

	"github.com/gorundebug/servicelib/runtime"

	"github.com/gorundebug/servicelib/runtime/environment"
	"github.com/gorundebug/servicelib/transformation"
)

var _ transformation.MapFunction[string, string] = (*ProcessSequentialActivityB)(nil)

// ProcessSequentialActivityB
type ProcessSequentialActivityB struct{}

func (f *ProcessSequentialActivityB) Map(ctx context.Context, _ runtime.Stream, value string, out runtime.Collect[string]) {
	out.Out(ctx, "sequential:b:"+value)
}

// MakeProcessSequentialActivityB is instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeProcessSequentialActivityB(ctx context.Context, env environment.ServiceEnvironment) (*ProcessSequentialActivityB, error) {
	return &ProcessSequentialActivityB{}, nil
}
