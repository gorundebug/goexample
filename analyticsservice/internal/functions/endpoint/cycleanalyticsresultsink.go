package endpoint

import (
	"context"
	"fmt"

	localsink "github.com/gorundebug/servicelib/datasink/localsink"
	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/environment"

	"github.com/gorundebug/analyticsservice/internal/types"
)

// cycleAnalyticsResultSinkHandler is a type alias for the EndpointHandler generic instantiation used throughout this file.
type cycleAnalyticsResultSinkHandler = localsink.EndpointHandler[CycleAnalyticsResultSinkHandlerState, *types.AnalyticsEvent, error]

var _ cycleAnalyticsResultSinkHandler = (*CycleAnalyticsResultSink)(nil)

func MakeEndpointConsumerCycleAnalyticsResultSink(stream runtime.TypedSinkStream[*types.AnalyticsEvent, error], handler cycleAnalyticsResultSinkHandler) (runtime.Consumer[*types.AnalyticsEvent], error) {
	return localsink.MakeCustomEndpointConsumer[CycleAnalyticsResultSinkHandlerState, *types.AnalyticsEvent, error](stream, handler)
}

// CycleAnalyticsResultSinkHandlerState holds per-stream state created by BeginRequest for each logical stream.
// Enables safe concurrent processing — no synchronization needed between streams.
// Add fields here to carry data across BeginRequest → ConsumeMessage → EndRequest.
type CycleAnalyticsResultSinkHandlerState struct {
}

// CycleAnalyticsResultSink
type CycleAnalyticsResultSink struct{}

// GetStreamID groups messages into logical streams (one BeginRequest/EndRequest per stream ID).
// Messages with the same ID share a CycleAnalyticsResultSinkHandlerState instance; return "" to route all messages to one stream.
func (ep *CycleAnalyticsResultSink) GetStreamID(_ context.Context, _ *types.AnalyticsEvent) string {
	//TODO: return grouping key from *types.AnalyticsEvent (e.g. tenant/session ID), or ""
	return ""
}

// BeginRequest is called once per stream (per unique GetStreamID), before any ConsumeMessage.
// Does NOT return an error — initialise CycleAnalyticsResultSinkHandlerState and attach outgoing metadata to ctx if needed.
// Validate the terminal event emitted after three passes through the feedback cycle.
func (ep *CycleAnalyticsResultSink) BeginRequest(ctx context.Context, _ runtime.Stream) (context.Context, CycleAnalyticsResultSinkHandlerState) {
	//TODO: initialise CycleAnalyticsResultSinkHandlerState, set up any per-stream resources
	return ctx, CycleAnalyticsResultSinkHandlerState{}
}

// ConsumeMessage processes one value *types.AnalyticsEvent from the stream.
// MUST:
//   - Perform the sink operation (write to DB, call API, etc.)
//   - Map result to error and push downstream via resultStream.Out(ctx, r)
//   - Leave resultStream unused if this is a terminal sink with no downstream consumers
//
// Return non-nil error to abort; EndRequest is called with that error.
func (ep *CycleAnalyticsResultSink) ConsumeMessage(_ context.Context, _ runtime.Stream, _ CycleAnalyticsResultSinkHandlerState, value *types.AnalyticsEvent, resultStream runtime.Collect[error]) error {
	if value.Key != "cycle" || value.Kind != "cycle" || value.Value != 3 {
		return fmt.Errorf("cycle analytics: unexpected result key=%q kind=%q value=%d", value.Key, value.Kind, value.Value)
	}
	return nil
}

// EndRequest finalises the stream after all messages are processed (or on error).
// err is the first non-nil error from ConsumeMessage; nil on the happy path.
// Does NOT return an error — flush/commit here; release resources.
func (ep *CycleAnalyticsResultSink) EndRequest(_ context.Context, _ runtime.Stream, err error, _ CycleAnalyticsResultSinkHandlerState) {
	//TODO: flush/commit buffered state, release resources; err != nil means ConsumeMessage failed
}

// MakeCycleAnalyticsResultSink implements the handler for the CycleAnalyticsResultSink local sink endpoint.
// It receives messages from the stream and writes them to a local destination.
// Instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeCycleAnalyticsResultSink(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.CustomEndpointConfig) (*CycleAnalyticsResultSink, error) {
	return &CycleAnalyticsResultSink{}, nil
}
