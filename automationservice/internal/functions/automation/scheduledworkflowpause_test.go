package automation

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
)

// Use the official Temporal Workflow timer for a scheduled Workflow.

func TestScheduledWorkflowPause_Duration(t *testing.T) {
	f := &ScheduledWorkflowPause{}
	stream := &automationDelayStream{cfg: &runtimecfg.DelayStreamConfig{Duration: 35}}
	result := f.Duration(context.Background(), stream, "job")
	assert.Equal(t, 35*time.Millisecond, result)
}

func TestScheduledWorkflowPause_DelayError(t *testing.T) {
	f := &ScheduledWorkflowPause{}
	f.DelayError(context.Background(), nil, "job", context.DeadlineExceeded, nil)
}
