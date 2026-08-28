package cron

import (
	"context"

	"github.com/gorundebug/servicelib/datasource"
	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/environment"
)

var _ runtime.ScheduleEndpointFunction[string] = (*AnalyticsScheduleSource)(nil)

func MakeEndpointConsumerAnalyticsScheduleSource[T, R, E any](
	stream runtime.TypedInputStream[T, R, E],
	function runtime.ScheduleEndpointFunction[T],
) (runtime.Consumer[T], error) {
	return datasource.GocronEndpointConsumer(stream, function)
}

// AnalyticsScheduleSource converts a scheduler trigger into zero or more values for the
// existing typed input stream.
type AnalyticsScheduleSource struct{}

// // Create an analytics job message identifying the local scheduled firing.
func (f *AnalyticsScheduleSource) OnTrigger(
	ctx context.Context,
	trigger runtime.ScheduleTrigger,
	out runtime.Collect[string],
) {
	out.Out(ctx, "analytics:"+trigger.ScheduleID+":"+trigger.TriggerID)
}

// MakeAnalyticsScheduleSource constructs the endpoint function once during service startup.
func MakeAnalyticsScheduleSource(
	_ context.Context,
	_ environment.ServiceEnvironment,
	_ *runtimecfg.CronEndpointConfig,
) (*AnalyticsScheduleSource, error) {
	return &AnalyticsScheduleSource{}, nil
}
