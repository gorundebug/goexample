package functions

import (
	"context"

	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/environment"
	"github.com/gorundebug/servicelib/transformation"
)

var _ transformation.MapFunction[string, string] = (*ProcessDurableJob)(nil)

// ProcessDurableJob
type ProcessDurableJob struct{}

func (f *ProcessDurableJob) Map(ctx context.Context, _ runtime.Stream, value string, out runtime.Collect[string]) {
	out.Out(ctx, "processed:"+value)
	if err := runtime.DurableCallSuccess(ctx); err != nil {
		panic(err)
	}
}

// MakeProcessDurableJob is instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeProcessDurableJob(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.MapStreamConfig) (*ProcessDurableJob, error) {
	return &ProcessDurableJob{}, nil
}
