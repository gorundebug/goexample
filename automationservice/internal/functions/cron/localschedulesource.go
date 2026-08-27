package cron

import (
	"context"
	"fmt"

	"github.com/gorundebug/servicelib/datasource"
	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/environment"
)

var _ runtime.ScheduleEndpointFunction[string] = (*LocalScheduleSource)(nil)

func MakeEndpointConsumerLocalScheduleSource[T, R, E any](
	stream runtime.TypedInputStream[T, R, E],
	function runtime.ScheduleEndpointFunction[T],
) (runtime.Consumer[T], error) {
	return datasource.GocronEndpointConsumer(stream, function)
}

// LocalScheduleSource converts a scheduler trigger into zero or more values for the
// existing typed input stream.
type LocalScheduleSource struct{}

func (f *LocalScheduleSource) OnTrigger(
	ctx context.Context,
	trigger runtime.ScheduleTrigger,
	out runtime.Collect[string],
) {
	out.Out(ctx, fmt.Sprintf("local:%s:%s", trigger.ScheduleID, trigger.TriggerID))
}

// MakeLocalScheduleSource constructs the endpoint function once during service startup.
func MakeLocalScheduleSource(
	_ context.Context,
	_ environment.ServiceEnvironment,
	_ *runtimecfg.CronEndpointConfig,
) (*LocalScheduleSource, error) {
	return &LocalScheduleSource{}, nil
}
