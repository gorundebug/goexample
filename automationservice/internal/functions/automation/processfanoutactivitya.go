package automation

import (
	"context"

	"github.com/gorundebug/servicelib/runtime"

	"github.com/gorundebug/servicelib/runtime/environment"
	"github.com/gorundebug/servicelib/transformation"
)

var _ transformation.MapFunction[string, string] = (*ProcessFanoutActivityA)(nil)

// ProcessFanoutActivityA
type ProcessFanoutActivityA struct{}

func (f *ProcessFanoutActivityA) Map(ctx context.Context, _ runtime.Stream, value string, out runtime.Collect[string]) {
	out.Out(ctx, "fanout:a:"+value)
}

// MakeProcessFanoutActivityA is instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeProcessFanoutActivityA(ctx context.Context, env environment.ServiceEnvironment) (*ProcessFanoutActivityA, error) {
	return &ProcessFanoutActivityA{}, nil
}
