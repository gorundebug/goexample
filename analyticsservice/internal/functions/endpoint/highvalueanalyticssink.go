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

// highValueAnalyticsSinkHandler is a type alias for the EndpointHandler generic instantiation used throughout this file.
type highValueAnalyticsSinkHandler = localsink.EndpointHandler[HighValueAnalyticsSinkHandlerState, *types.AnalyticsResult, error]

var _ highValueAnalyticsSinkHandler = (*HighValueAnalyticsSink)(nil)

func MakeEndpointConsumerHighValueAnalyticsSink(stream runtime.TypedSinkStream[*types.AnalyticsResult, error], handler highValueAnalyticsSinkHandler) (runtime.Consumer[*types.AnalyticsResult], error) {
	return localsink.MakeCustomEndpointConsumer[HighValueAnalyticsSinkHandlerState, *types.AnalyticsResult, error](stream, handler)
}

// HighValueAnalyticsSinkHandlerState holds per-stream state created by BeginRequest for each logical stream.
// Enables safe concurrent processing — no synchronization needed between streams.
// Add fields here to carry data across BeginRequest → ConsumeMessage → EndRequest.
type HighValueAnalyticsSinkHandlerState struct {
}

// HighValueAnalyticsSink
type HighValueAnalyticsSink struct{}

// GetStreamID groups messages into logical streams (one BeginRequest/EndRequest per stream ID).
// Messages with the same ID share a HighValueAnalyticsSinkHandlerState instance; return "" to route all messages to one stream.
func (ep *HighValueAnalyticsSink) GetStreamID(_ context.Context, _ *types.AnalyticsResult) string {
	//TODO: return grouping key from *types.AnalyticsResult (e.g. tenant/session ID), or ""
	return ""
}

// BeginRequest is called once per stream (per unique GetStreamID), before any ConsumeMessage.
// Does NOT return an error — initialise HighValueAnalyticsSinkHandlerState and attach outgoing metadata to ctx if needed.
// // Validate and record analytics results routed to the high-value Case branch.
func (ep *HighValueAnalyticsSink) BeginRequest(ctx context.Context, _ runtime.Stream) (context.Context, HighValueAnalyticsSinkHandlerState) {
	//TODO: initialise HighValueAnalyticsSinkHandlerState, set up any per-stream resources
	return ctx, HighValueAnalyticsSinkHandlerState{}
}

// ConsumeMessage processes one value *types.AnalyticsResult from the stream.
// MUST:
//   - Perform the sink operation (write to DB, call API, etc.)
//   - Map result to error and push downstream via resultStream.Out(ctx, r)
//   - Leave resultStream unused if this is a terminal sink with no downstream consumers
//
// Return non-nil error to abort; EndRequest is called with that error.
func (ep *HighValueAnalyticsSink) ConsumeMessage(_ context.Context, _ runtime.Stream, _ HighValueAnalyticsSinkHandlerState, value *types.AnalyticsResult, resultStream runtime.Collect[error]) error {
	if value.Key != "high-value" || value.Total != 60 || value.Kind != "multi" {
		return fmt.Errorf("high-value analytics: unexpected result key=%q total=%d kind=%q", value.Key, value.Total, value.Kind)
	}
	return nil
}

// EndRequest finalises the stream after all messages are processed (or on error).
// err is the first non-nil error from ConsumeMessage; nil on the happy path.
// Does NOT return an error — flush/commit here; release resources.
func (ep *HighValueAnalyticsSink) EndRequest(_ context.Context, _ runtime.Stream, err error, _ HighValueAnalyticsSinkHandlerState) {
	//TODO: flush/commit buffered state, release resources; err != nil means ConsumeMessage failed
}

// MakeHighValueAnalyticsSink implements the handler for the HighValueAnalyticsSink local sink endpoint.
// It receives messages from the stream and writes them to a local destination.
// Instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeHighValueAnalyticsSink(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.CustomEndpointConfig) (*HighValueAnalyticsSink, error) {
	return &HighValueAnalyticsSink{}, nil
}
