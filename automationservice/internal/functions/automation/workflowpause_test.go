package automation

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
)

// Use the same Delay contract backed by the Temporal Workflow timer.

func TestWorkflowPause_Duration(t *testing.T) {
	f := &WorkflowPause{}
	stream := &automationDelayStream{cfg: &runtimecfg.DelayStreamConfig{Duration: 40}}
	result := f.Duration(context.Background(), stream, "job")
	assert.Equal(t, 40*time.Millisecond, result)
}

func TestWorkflowPause_DelayError(t *testing.T) {
	f := &WorkflowPause{}
	f.DelayError(context.Background(), nil, "job", context.DeadlineExceeded, nil)
}
