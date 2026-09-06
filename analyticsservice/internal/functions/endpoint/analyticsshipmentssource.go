package endpoint

import (
	"context"

	localsource "github.com/gorundebug/servicelib/datasource/localsource"
	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/environment"

	"github.com/gorundebug/analyticsservice/internal/types"
)

var _ localsource.EndpointHandler[AnalyticsShipmentsSourceHandlerState, *types.AnalyticsEvent, any, error] = (*AnalyticsShipmentsSource)(nil)

func MakeEndpointConsumerAnalyticsShipmentsSource(stream runtime.TypedInputStream[*types.AnalyticsEvent, any, error], handler *AnalyticsShipmentsSource) (runtime.Consumer[*types.AnalyticsEvent], error) {
	return localsource.MakeCustomEndpointConsumer[AnalyticsShipmentsSourceHandlerState, *types.AnalyticsEvent, any, error](stream, handler, handler)
}

// AnalyticsShipmentsSourceHandlerState holds per-message state created by BeginRequest for each incoming value.
// Enables safe concurrent processing — no synchronization needed between messages.
// Add fields here to carry data across BeginRequest → ConsumeMessage → EndRequest.
type AnalyticsShipmentsSourceHandlerState struct {
}

// AnalyticsShipmentsSource
type AnalyticsShipmentsSource struct {
}

func (ep *AnalyticsShipmentsSource) Start(ctx context.Context, consumer runtime.Consumer[*types.AnalyticsEvent]) error {
	consumer.Consume(ctx, &types.AnalyticsEvent{Key: "high-value", Value: 30, Kind: "shipment"})
	consumer.Consume(ctx, &types.AnalyticsEvent{Key: "standard", Value: 3, Kind: "shipment"})
	return nil
}

func (ep *AnalyticsShipmentsSource) Stop(context.Context) {}

// Concurrency returns the maximum number of messages processed concurrently.
// Return 0 for unlimited concurrency.
func (ep *AnalyticsShipmentsSource) Concurrency(_ localsource.StreamContext[*types.AnalyticsEvent, any, error]) int {
	//TODO: return max concurrent goroutines, or 0 for unlimited
	return 0
}

// BeginRequest is called once per incoming value, before ConsumeMessage.
// MUST:
//   - Initialise per-value state in AnalyticsShipmentsSourceHandlerState
//   - Return non-nil error to reject the value — framework will NOT call
//     ConsumeMessage or EndRequest; release any acquired resources here.
//
// // Produce a deterministic shipment analytics event for the canonical multi-way join example.
func (ep *AnalyticsShipmentsSource) BeginRequest(ctx context.Context, _ localsource.StreamContext[*types.AnalyticsEvent, any, error]) (context.Context, AnalyticsShipmentsSourceHandlerState, error) {
	//TODO: initialise AnalyticsShipmentsSourceHandlerState; return error to reject this value
	return ctx, AnalyticsShipmentsSourceHandlerState{}, nil
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
// ConsumeMessage — synchronise access to AnalyticsShipmentsSourceHandlerState accordingly.
// Return non-nil error to abort; EndRequest is called with that error.
func (ep *AnalyticsShipmentsSource) ConsumeMessage(ctx context.Context, sc localsource.StreamContext[*types.AnalyticsEvent, any, error], _ AnalyticsShipmentsSourceHandlerState, value *types.AnalyticsEvent, resultCtx localsource.ResultContext[AnalyticsShipmentsSourceHandlerState, *types.AnalyticsEvent, any, error]) error {
	sc.Collect(ctx, value)
	resultCtx.Done()
	return nil
}

// GetMessageID returns a stable correlation ID from any so the framework can
// route async pipeline results back to the correct ResultContext callback.
// Return "" if this endpoint does not use async result routing.
func (ep *AnalyticsShipmentsSource) GetMessageID(_ context.Context, _ localsource.StreamContext[*types.AnalyticsEvent, any, error], _ AnalyticsShipmentsSourceHandlerState, _ any) string {
	//TODO: return stable ID from any (e.g. request_id field), or ""
	return ""
}

// EndRequest finalises processing after Done() is signalled (or on ConsumeMessage error).
// err is the first non-nil error from any earlier stage; nil on the happy path.
// Does NOT return an error — log or record metrics here; release resources.
func (ep *AnalyticsShipmentsSource) EndRequest(_ context.Context, _ localsource.StreamContext[*types.AnalyticsEvent, any, error], err error, _ AnalyticsShipmentsSourceHandlerState) {
	//TODO: release resources, log outcome (err != nil means processing failed)
}

// MakeAnalyticsShipmentsSource implements the handler for the AnalyticsShipmentsSource local source endpoint.
// It reads data from a local source and produces messages into the stream.
// Instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeAnalyticsShipmentsSource(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.CustomEndpointConfig) (*AnalyticsShipmentsSource, error) {
	return &AnalyticsShipmentsSource{}, nil
}
