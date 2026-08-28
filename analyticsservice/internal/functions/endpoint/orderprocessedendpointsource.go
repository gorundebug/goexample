package endpoint

import (
	"context"
	"encoding/json"

	datasourcekafka "github.com/gorundebug/servicelib/datasource/kafka"
	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/environment"

	"github.com/gorundebug/model_go/pkg/types"
)

// orderProcessedEndpointHandler is a type alias for the EndpointHandler generic instantiation used throughout this file.
type orderProcessedEndpointHandler = datasourcekafka.EndpointHandler[OrderProcessedEndpointSourceHandlerState, *types.OrderProcessed, *types.OrderProcessed, error]

var _ orderProcessedEndpointHandler = (*OrderProcessedEndpointSource)(nil)

func MakeEndpointConsumerOrderProcessedEndpointSource(stream runtime.TypedInputStream[*types.OrderProcessed, *types.OrderProcessed, error], handler orderProcessedEndpointHandler) (runtime.Consumer[*types.OrderProcessed], error) {
	return datasourcekafka.MakeSaramaKafkaEndpointConsumer[OrderProcessedEndpointSourceHandlerState, *types.OrderProcessed, *types.OrderProcessed, error](stream, handler)
}

// OrderProcessedEndpointSourceHandlerState holds per-message state created by BeginRequest for each Kafka message.
// Enables safe concurrent processing — no synchronization needed between messages.
// Add fields here to carry data across BeginRequest → ConsumeMessage → EndRequest.
type OrderProcessedEndpointSourceHandlerState struct {
}

// OrderProcessedEndpointSource
type OrderProcessedEndpointSource struct{}

// Concurrency returns the maximum number of messages processed concurrently.
// Return 0 for unlimited concurrency (bounded only by partition count).
func (ep *OrderProcessedEndpointSource) Concurrency(_ datasourcekafka.StreamContext[*types.OrderProcessed, *types.OrderProcessed, error]) int {
	return 0
}

// BeginRequest is called once per Kafka message, before ConsumeMessage.
// MUST:
//   - Extract trace/correlation data from msgData.Headers and store in OrderProcessedEndpointSourceHandlerState
//   - Return non-nil error to reject the message — framework will NOT call
//     ConsumeMessage or EndRequest; release any acquired resources here.
//
// // Decode each OrderProcessed event and pass it to the analytics pipeline. Mark the Kafka message processed only after
// the pipeline has handled it successfully.
func (ep *OrderProcessedEndpointSource) BeginRequest(ctx context.Context, _ datasourcekafka.StreamContext[*types.OrderProcessed, *types.OrderProcessed, error]) (context.Context, OrderProcessedEndpointSourceHandlerState, error) {
	return ctx, OrderProcessedEndpointSourceHandlerState{}, nil
}

// ConsumeMessage is called once per Kafka message.
// Typical pattern (async pipeline result):
//  1. Deserialize msgData.Value ([]byte) → *types.OrderProcessed and emit via sc.Collect
//  2. Register result callback: resultCtx.SetResultCallback(id, func(...) bool {
//     push result downstream; return true
//     })
//  3. Call resultCtx.Done() to signal end of synchronous processing
//
// Explicit offset control:
//   - msgData.MarkMessage() — mark offset for auto-commit (default Sarama behaviour)
//   - msgData.Commit()      — commit offset immediately (manual commit mode)
//
// Thread safety: GetMessageID and result callbacks may run CONCURRENTLY with
// ConsumeMessage — synchronise access to OrderProcessedEndpointSourceHandlerState accordingly.
// Return non-nil error to abort; EndRequest is called with that error.
func (ep *OrderProcessedEndpointSource) ConsumeMessage(ctx context.Context, sc datasourcekafka.StreamContext[*types.OrderProcessed, *types.OrderProcessed, error], _ OrderProcessedEndpointSourceHandlerState, msgData datasourcekafka.ConsumerMessage, resultCtx datasourcekafka.ResultContext[OrderProcessedEndpointSourceHandlerState, *types.OrderProcessed, *types.OrderProcessed, error]) error {
	value := &types.OrderProcessed{}
	if err := json.Unmarshal(msgData.Value, value); err != nil {
		return err
	}
	resultCtx.SetResultCallback(value.OrderID, func(_ context.Context, _ datasourcekafka.StreamContext[*types.OrderProcessed, *types.OrderProcessed, error], _ OrderProcessedEndpointSourceHandlerState, _ *types.OrderProcessed) bool {
		msgData.MarkMessage("")
		resultCtx.Done()
		return true
	})
	sc.Collect(ctx, value)
	return nil
}

// GetMessageID returns a stable correlation ID from *types.OrderProcessed so the framework can
// route async pipeline results back to the correct ResultContext callback.
// Return "" if this endpoint does not use async result routing.
func (ep *OrderProcessedEndpointSource) GetMessageID(_ context.Context, _ datasourcekafka.StreamContext[*types.OrderProcessed, *types.OrderProcessed, error], _ OrderProcessedEndpointSourceHandlerState, value *types.OrderProcessed) string {
	return value.OrderID
}

// EndRequest finalises processing after Done() is signalled (or on ConsumeMessage error).
// err is the first non-nil error from any earlier stage; nil on the happy path.
// Does NOT return an error — log or record metrics here; release resources.
func (ep *OrderProcessedEndpointSource) EndRequest(_ context.Context, _ datasourcekafka.StreamContext[*types.OrderProcessed, *types.OrderProcessed, error], err error, _ OrderProcessedEndpointSourceHandlerState) {
}

// MakeOrderProcessedEndpointSource implements the handler for the OrderProcessed Kafka source endpoint.
// It consumes messages from a Kafka topic and produces them into the stream.
// Instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeOrderProcessedEndpointSource(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.KafkaEndpointConfig) (*OrderProcessedEndpointSource, error) {
	return &OrderProcessedEndpointSource{}, nil
}
