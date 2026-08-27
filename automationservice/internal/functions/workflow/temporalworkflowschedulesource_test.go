package workflow

import (
	"context"
	"testing"

	"github.com/gorundebug/servicelib/runtime"
	"github.com/stretchr/testify/require"
)

func TestTemporalWorkflowScheduleSourceOnTrigger(t *testing.T) {
	function := &TemporalWorkflowScheduleSource{}
	var collected []string
	out := runtime.CollectFunc[string](func(_ context.Context, value string) {
		collected = append(collected, value)
	})
	trigger := runtime.ScheduleTrigger{ScheduleID: "workflow-schedule", TriggerID: "trigger-2"}

	function.OnTrigger(context.Background(), trigger, out)

	require.Equal(t, []string{"scheduled-workflow:workflow-schedule:trigger-2"}, collected)
}
