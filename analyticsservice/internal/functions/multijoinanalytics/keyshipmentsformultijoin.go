package multijoinanalytics

import (
	"context"

	"github.com/gorundebug/analyticsservice/internal/types"

	"github.com/gorundebug/servicelib/runtime"
	runtimecfg "github.com/gorundebug/servicelib/runtime/config"
	"github.com/gorundebug/servicelib/runtime/datastruct"
	"github.com/gorundebug/servicelib/runtime/environment"
	"github.com/gorundebug/servicelib/transformation"
)

var _ transformation.KeyByFunction[*types.AnalyticsEvent, string, *types.AnalyticsEvent] = (*KeyShipmentsForMultiJoin)(nil)

// KeyShipmentsForMultiJoin
type KeyShipmentsForMultiJoin struct{}

func (f *KeyShipmentsForMultiJoin) KeyBy(ctx context.Context, _ runtime.Stream, value *types.AnalyticsEvent, out runtime.Collect[datastruct.KeyValue[string, *types.AnalyticsEvent]]) {
	out.Out(ctx, datastruct.KeyValue[string, *types.AnalyticsEvent]{Key: value.Key, Value: value})
}

// MakeKeyShipmentsForMultiJoin is instantiated once at application startup via its maker function.
// Fields of this struct are not protected by any synchronization — do not use
// shared mutable state here without external synchronization.
func MakeKeyShipmentsForMultiJoin(ctx context.Context, env environment.ServiceEnvironment, cfg *runtimecfg.KeyByStreamConfig) (*KeyShipmentsForMultiJoin, error) {
	return &KeyShipmentsForMultiJoin{}, nil
}
