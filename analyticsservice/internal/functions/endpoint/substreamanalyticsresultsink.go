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

// substreamAnalyticsResultSinkHandler is a type alias for the EndpointHandler generic instantiation used throughout this file.
type substreamAnalyticsResultSinkHandler = localsink.EndpointHandler[SubstreamAnalyticsResultSinkHandlerState, *types.AnalyticsResult, error]

var _ substreamAnalyticsResultSinkHandler = (*SubstreamAnalyticsResultSink)(nil)

func MakeEndpointConsumerSubstreamAnalyticsResultSink(stream runtime.TypedSinkStream[*types.AnalyticsResult, error], handler substreamAnalyticsResultSinkHandler) (runtime.Consumer[*types.AnalyticsResult], error) {
	return localsink.MakeCustomEndpointConsumer[SubstreamAnalyticsResultSinkHandlerState, *types.AnalyticsResult, error](stream, handler)
}

// SubstreamAnalyticsResultSinkHandlerState holds per-stream state created by BeginRequest for each logical stream.
// Enables safe concurrent processing — no synchronization needed between streams.
// Add fields here to carry data across BeginRequest → ConsumeMessage → EndRequest.
type SubstreamAnalyticsResultSinkHandlerState struct {
}

// SubstreamAnalyticsResultSink
type SubstreamAnalyticsResultSink struct{}

// GetStreamID groups messages into logical streams (one BeginRequest/EndRequest per stream ID).
// Messages with the same ID share a SubstreamAnalyticsResultSinkHandlerState instance; return "" to route all messages to one stream.
func (ep *SubstreamAnalyticsResultSink) GetStreamID(_ context.Context, _ *types.AnalyticsResult) string {
	//TODO: return grouping key from *types.AnalyticsResult (e.g. tenant/session ID), or ""
	return ""
}

// BeginRequest is called once per stream (per unique GetStreamID), before any ConsumeMessage.
// Does NOT return an error — initialise SubstreamAnalyticsResultSinkHandlerState and attach outgoing metadata to ctx if needed.
// // Validate and record the result returned by the service-local SubStream example.
func (ep *SubstreamAnalyticsResultSink) BeginRequest(ctx context.Context, _ runtime.Stream) (context.Context, SubstreamAnalyticsResultSinkHandlerState) {
	//TODO: initialise SubstreamAnalyticsResultSinkHandlerState, set up any per-stream resources
	return ctx, SubstreamAnalyticsResultSinkHandlerState{}
}

// ConsumeMessage processes one value *types.AnalyticsResult from the stream.
// MUST:
//   - Perform the sink operation (write to DB, call API, etc.)
//   - Map result to error and push downstream via resultStream.Out(ctx, r)
//   - Leave resultStream unused if this is a terminal sink with no downstream consumers
//
// Return non-nil error to abort; EndRequest is called with that error.
func (ep *SubstreamAnalyticsResultSink) ConsumeMessage(_ context.Context, _ runtime.Stream, _ SubstreamAnalyticsResultSinkHandlerState, value *types.AnalyticsResult, resultStream runtime.Collect[error]) error {
	if value.Key != "substream" || value.Total != 14 || value.Kind != "substream" {
		return fmt.Errorf("substream analytics: unexpected result key=%q total=%d kind=%q", value.Key, value.Total, value.Kind)
	}
	return nil
}

// EndRequest finalises the stream after all messages are processed (or on error).
// err is the first non-nil error from ConsumeMessage; nil on the happy path.
// Does NOT return an error — flush/commit here; release resources.
func (ep *SubstreamAnalyticsResultSink) EndRequest(_ context.Context, _ runtime.Stream, err error, _ SubstreamAnalyticsResultSinkHandlerState) {
	//TODO: flush/commit buffered state, release resources; err != nil means ConsumeMessage failed
}

// MakeSubstreamAnalyticsResultSink implements the handler for the SubstreamAnalyticsResultSink local sink endpoint.
// It receives messages from the stream and writes them to a local destination.
// Instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeSubstreamAnalyticsResultSink(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.CustomEndpointConfig) (*SubstreamAnalyticsResultSink, error) {
	return &SubstreamAnalyticsResultSink{}, nil
}
