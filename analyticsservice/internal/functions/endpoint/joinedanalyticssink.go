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

// joinedAnalyticsSinkHandler is a type alias for the EndpointHandler generic instantiation used throughout this file.
type joinedAnalyticsSinkHandler = localsink.EndpointHandler[JoinedAnalyticsSinkHandlerState, *types.AnalyticsResult, error]

var _ joinedAnalyticsSinkHandler = (*JoinedAnalyticsSink)(nil)

func MakeEndpointConsumerJoinedAnalyticsSink(stream runtime.TypedSinkStream[*types.AnalyticsResult, error], handler joinedAnalyticsSinkHandler) (runtime.Consumer[*types.AnalyticsResult], error) {
	return localsink.MakeCustomEndpointConsumer[JoinedAnalyticsSinkHandlerState, *types.AnalyticsResult, error](stream, handler)
}

// JoinedAnalyticsSinkHandlerState holds per-stream state created by BeginRequest for each logical stream.
// Enables safe concurrent processing — no synchronization needed between streams.
// Add fields here to carry data across BeginRequest → ConsumeMessage → EndRequest.
type JoinedAnalyticsSinkHandlerState struct {
}

// JoinedAnalyticsSink
type JoinedAnalyticsSink struct{}

// GetStreamID groups messages into logical streams (one BeginRequest/EndRequest per stream ID).
// Messages with the same ID share a JoinedAnalyticsSinkHandlerState instance; return "" to route all messages to one stream.
func (ep *JoinedAnalyticsSink) GetStreamID(_ context.Context, _ *types.AnalyticsResult) string {
	//TODO: return grouping key from *types.AnalyticsResult (e.g. tenant/session ID), or ""
	return ""
}

// BeginRequest is called once per stream (per unique GetStreamID), before any ConsumeMessage.
// Does NOT return an error — initialise JoinedAnalyticsSinkHandlerState and attach outgoing metadata to ctx if needed.
// // Validate and record the result of the two-way analytics join.
func (ep *JoinedAnalyticsSink) BeginRequest(ctx context.Context, _ runtime.Stream) (context.Context, JoinedAnalyticsSinkHandlerState) {
	//TODO: initialise JoinedAnalyticsSinkHandlerState, set up any per-stream resources
	return ctx, JoinedAnalyticsSinkHandlerState{}
}

// ConsumeMessage processes one value *types.AnalyticsResult from the stream.
// MUST:
//   - Perform the sink operation (write to DB, call API, etc.)
//   - Map result to error and push downstream via resultStream.Out(ctx, r)
//   - Leave resultStream unused if this is a terminal sink with no downstream consumers
//
// Return non-nil error to abort; EndRequest is called with that error.
func (ep *JoinedAnalyticsSink) ConsumeMessage(_ context.Context, _ runtime.Stream, _ JoinedAnalyticsSinkHandlerState, value *types.AnalyticsResult, resultStream runtime.Collect[error]) error {
	if value.Kind != "join" {
		return fmt.Errorf("joined analytics: unexpected kind %q", value.Kind)
	}
	expected := map[string]int{"high-value": 30, "standard": 3}
	want, ok := expected[value.Key]
	if !ok || value.Total != want {
		return fmt.Errorf("joined analytics: unexpected result key=%q total=%d", value.Key, value.Total)
	}
	return nil
}

// EndRequest finalises the stream after all messages are processed (or on error).
// err is the first non-nil error from ConsumeMessage; nil on the happy path.
// Does NOT return an error — flush/commit here; release resources.
func (ep *JoinedAnalyticsSink) EndRequest(_ context.Context, _ runtime.Stream, err error, _ JoinedAnalyticsSinkHandlerState) {
	//TODO: flush/commit buffered state, release resources; err != nil means ConsumeMessage failed
}

// MakeJoinedAnalyticsSink implements the handler for the JoinedAnalyticsSink local sink endpoint.
// It receives messages from the stream and writes them to a local destination.
// Instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeJoinedAnalyticsSink(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.CustomEndpointConfig) (*JoinedAnalyticsSink, error) {
	return &JoinedAnalyticsSink{}, nil
}
