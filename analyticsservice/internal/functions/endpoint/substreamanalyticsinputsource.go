package endpoint

import (
	"context"

	localsource "github.com/gorundebug/servicelib/datasource/localsource"
	"github.com/gorundebug/servicelib/runtime"

	"github.com/gorundebug/servicelib/runtime/environment"

	"github.com/gorundebug/analyticsservice/internal/types"
)

var _ localsource.EndpointHandler[SubstreamAnalyticsInputSourceHandlerState, *types.AnalyticsEvent, any, error] = (*SubstreamAnalyticsInputSource)(nil)

func MakeEndpointConsumerSubstreamAnalyticsInputSource(stream runtime.TypedInputStream[*types.AnalyticsEvent, any, error], handler *SubstreamAnalyticsInputSource) (runtime.Consumer[*types.AnalyticsEvent], error) {
	return localsource.MakeCustomEndpointConsumer[SubstreamAnalyticsInputSourceHandlerState, *types.AnalyticsEvent, any, error](stream, handler, handler)
}

// SubstreamAnalyticsInputSourceHandlerState holds per-message state created by BeginRequest for each incoming value.
// Enables safe concurrent processing — no synchronization needed between messages.
// Add fields here to carry data across BeginRequest → ConsumeMessage → EndRequest.
type SubstreamAnalyticsInputSourceHandlerState struct {
}

// SubstreamAnalyticsInputSource
type SubstreamAnalyticsInputSource struct {
}

// Start produces values for this custom source until the context is cancelled
// or the producer has no more values. The framework owns the lifecycle and calls
// Stop during service shutdown.
func (ep *SubstreamAnalyticsInputSource) Start(ctx context.Context, consumer runtime.Consumer[*types.AnalyticsEvent]) error {
	consumer.Consume(ctx, &types.AnalyticsEvent{Key: "substream", Value: 7, Kind: "input"})
	return nil
}

// Stop asks the custom producer to stop accepting new work.
func (ep *SubstreamAnalyticsInputSource) Stop(_ context.Context) {
	//TODO: stop producer resources
}

// Concurrency returns the maximum number of messages processed concurrently.
// Return 0 for unlimited concurrency.
func (ep *SubstreamAnalyticsInputSource) Concurrency(_ localsource.StreamContext[*types.AnalyticsEvent, any, error]) int {
	//TODO: return max concurrent goroutines, or 0 for unlimited
	return 0
}

// BeginRequest is called once per incoming value, before ConsumeMessage.
// MUST:
//   - Initialise per-value state in SubstreamAnalyticsInputSourceHandlerState
//   - Return non-nil error to reject the value — framework will NOT call
//     ConsumeMessage or EndRequest; release any acquired resources here.
//
// // Produce one deterministic analytics event that invokes the service-local SubStream example.
func (ep *SubstreamAnalyticsInputSource) BeginRequest(ctx context.Context, _ localsource.StreamContext[*types.AnalyticsEvent, any, error]) (context.Context, SubstreamAnalyticsInputSourceHandlerState, error) {
	//TODO: initialise SubstreamAnalyticsInputSourceHandlerState; return error to reject this value
	return ctx, SubstreamAnalyticsInputSourceHandlerState{}, nil
}

// ConsumeMessage is called once per incoming value *types.AnalyticsEvent.
// Typical pattern (async pipeline result):
//  1. Map value (*types.AnalyticsEvent) → pipeline message and emit via sc.Collect
//  2. Register result callback: resultCtx.SetResultCallback(id, func(...) bool {
//     push any downstream; return true
//     })
//  3. Call resultCtx.Done() to unblock the framework after emitting
//
// For a synchronous pattern: emit result directly via sc.Collect and call resultCtx.Done().
// Thread safety: GetMessageID and result callbacks may run CONCURRENTLY with
// ConsumeMessage — synchronise access to SubstreamAnalyticsInputSourceHandlerState accordingly.
// Return non-nil error to abort; EndRequest is called with that error.
func (ep *SubstreamAnalyticsInputSource) ConsumeMessage(ctx context.Context, sc localsource.StreamContext[*types.AnalyticsEvent, any, error], _ SubstreamAnalyticsInputSourceHandlerState, value *types.AnalyticsEvent, resultCtx localsource.ResultContext[SubstreamAnalyticsInputSourceHandlerState, *types.AnalyticsEvent, any, error]) error {
	sc.Collect(ctx, value)
	resultCtx.Done()
	return nil
}

// GetMessageID returns a stable correlation ID from any so the framework can
// route async pipeline results back to the correct ResultContext callback.
// Return "" if this endpoint does not use async result routing.
func (ep *SubstreamAnalyticsInputSource) GetMessageID(_ context.Context, _ localsource.StreamContext[*types.AnalyticsEvent, any, error], _ SubstreamAnalyticsInputSourceHandlerState, _ any) string {
	//TODO: return stable ID from any (e.g. request_id field), or ""
	return ""
}

// EndRequest finalises processing after Done() is signalled (or on ConsumeMessage error).
// err is the first non-nil error from any earlier stage; nil on the happy path.
// Does NOT return an error — log or record metrics here; release resources.
func (ep *SubstreamAnalyticsInputSource) EndRequest(_ context.Context, _ localsource.StreamContext[*types.AnalyticsEvent, any, error], err error, _ SubstreamAnalyticsInputSourceHandlerState) {
	//TODO: release resources, log outcome (err != nil means processing failed)
}

// MakeSubstreamAnalyticsInputSource implements the handler for the SubstreamAnalyticsInputSource local source endpoint.
// It reads data from a local source and produces messages into the stream.
// Instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeSubstreamAnalyticsInputSource(ctx context.Context, env environment.ServiceEnvironment) (*SubstreamAnalyticsInputSource, error) {
	return &SubstreamAnalyticsInputSource{}, nil
}
