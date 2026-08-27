package workflow

import (
	"context"
	"fmt"

	"github.com/gorundebug/servicelib/datasource"
	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/environment"
)

var _ runtime.ScheduleEndpointFunction[string] = (*TemporalWorkflowScheduleSource)(nil)

func MakeEndpointConsumerTemporalWorkflowScheduleSource[T, R, E any](
	stream runtime.TypedInputStream[T, R, E],
	function runtime.ScheduleEndpointFunction[T],
) (runtime.Consumer[T], error) {
	return datasource.TemporalScheduleEndpointConsumer(stream, function)
}

// TemporalWorkflowScheduleSource converts a scheduler trigger into zero or more values for the
// existing typed input stream.
type TemporalWorkflowScheduleSource struct{}

// // Create a Workflow job message identifying the durable scheduled firing.
func (f *TemporalWorkflowScheduleSource) OnTrigger(
	ctx context.Context,
	trigger runtime.ScheduleTrigger,
	out runtime.Collect[string],
) {
	out.Out(ctx, fmt.Sprintf("scheduled-workflow:%s:%s", trigger.ScheduleID, trigger.TriggerID))
}

// MakeTemporalWorkflowScheduleSource constructs the endpoint function once during service startup.
func MakeTemporalWorkflowScheduleSource(
	_ context.Context,
	_ environment.ServiceEnvironment,
	_ *runtimecfg.TemporalEndpointConfig,
) (*TemporalWorkflowScheduleSource, error) {
	return &TemporalWorkflowScheduleSource{}, nil
}
