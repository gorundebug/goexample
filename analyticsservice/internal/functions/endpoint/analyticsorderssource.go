package endpoint

import (
	"context"

	localsource "github.com/gorundebug/servicelib/datasource/localsource"
	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/environment"

	"github.com/gorundebug/analyticsservice/internal/types"
)

var _ localsource.EndpointHandler[AnalyticsOrdersSourceHandlerState, *types.AnalyticsEvent, any, error] = (*AnalyticsOrdersSource)(nil)

func MakeEndpointConsumerAnalyticsOrdersSource(stream runtime.TypedInputStream[*types.AnalyticsEvent, any, error], handler *AnalyticsOrdersSource) (runtime.Consumer[*types.AnalyticsEvent], error) {
	return localsource.MakeCustomEndpointConsumer[AnalyticsOrdersSourceHandlerState, *types.AnalyticsEvent, any, error](stream, handler, handler)
}

// AnalyticsOrdersSourceHandlerState holds per-message state created by BeginRequest for each incoming value.
// Enables safe concurrent processing — no synchronization needed between messages.
// Add fields here to carry data across BeginRequest → ConsumeMessage → EndRequest.
type AnalyticsOrdersSourceHandlerState struct {
}

// AnalyticsOrdersSource
type AnalyticsOrdersSource struct {
}

func (ep *AnalyticsOrdersSource) Start(ctx context.Context, consumer runtime.Consumer[*types.AnalyticsEvent]) error {
	consumer.Consume(ctx, &types.AnalyticsEvent{Key: "high-value", Value: 10, Kind: "order"})
	consumer.Consume(ctx, &types.AnalyticsEvent{Key: "standard", Value: 1, Kind: "order"})
	return nil
}

func (ep *AnalyticsOrdersSource) Stop(context.Context) {}

// Concurrency returns the maximum number of messages processed concurrently.
// Return 0 for unlimited concurrency.
func (ep *AnalyticsOrdersSource) Concurrency(_ localsource.StreamContext[*types.AnalyticsEvent, any, error]) int {
	//TODO: return max concurrent goroutines, or 0 for unlimited
	return 0
}

// BeginRequest is called once per incoming value, before ConsumeMessage.
// MUST:
//   - Initialise per-value state in AnalyticsOrdersSourceHandlerState
//   - Return non-nil error to reject the value — framework will NOT call
//     ConsumeMessage or EndRequest; release any acquired resources here.
//
// // Produce a deterministic order analytics event for the canonical join examples.
func (ep *AnalyticsOrdersSource) BeginRequest(ctx context.Context, _ localsource.StreamContext[*types.AnalyticsEvent, any, error]) (context.Context, AnalyticsOrdersSourceHandlerState, error) {
	//TODO: initialise AnalyticsOrdersSourceHandlerState; return error to reject this value
	return ctx, AnalyticsOrdersSourceHandlerState{}, nil
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
// ConsumeMessage — synchronise access to AnalyticsOrdersSourceHandlerState accordingly.
// Return non-nil error to abort; EndRequest is called with that error.
func (ep *AnalyticsOrdersSource) ConsumeMessage(ctx context.Context, sc localsource.StreamContext[*types.AnalyticsEvent, any, error], _ AnalyticsOrdersSourceHandlerState, value *types.AnalyticsEvent, resultCtx localsource.ResultContext[AnalyticsOrdersSourceHandlerState, *types.AnalyticsEvent, any, error]) error {
	sc.Collect(ctx, value)
	resultCtx.Done()
	return nil
}

// GetMessageID returns a stable correlation ID from any so the framework can
// route async pipeline results back to the correct ResultContext callback.
// Return "" if this endpoint does not use async result routing.
func (ep *AnalyticsOrdersSource) GetMessageID(_ context.Context, _ localsource.StreamContext[*types.AnalyticsEvent, any, error], _ AnalyticsOrdersSourceHandlerState, _ any) string {
	//TODO: return stable ID from any (e.g. request_id field), or ""
	return ""
}

// EndRequest finalises processing after Done() is signalled (or on ConsumeMessage error).
// err is the first non-nil error from any earlier stage; nil on the happy path.
// Does NOT return an error — log or record metrics here; release resources.
func (ep *AnalyticsOrdersSource) EndRequest(_ context.Context, _ localsource.StreamContext[*types.AnalyticsEvent, any, error], err error, _ AnalyticsOrdersSourceHandlerState) {
	//TODO: release resources, log outcome (err != nil means processing failed)
}

// MakeAnalyticsOrdersSource implements the handler for the AnalyticsOrdersSource local source endpoint.
// It reads data from a local source and produces messages into the stream.
// Instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeAnalyticsOrdersSource(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.CustomEndpointConfig) (*AnalyticsOrdersSource, error) {
	return &AnalyticsOrdersSource{}, nil
}
