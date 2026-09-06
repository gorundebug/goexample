package serdes

import (
	"fmt"

	"github.com/gorundebug/analyticsservice/internal/types"
)

type AnalyticsEventSerde struct{}

func (s *AnalyticsEventSerde) IsStub() bool {
	return false
}

func (s *AnalyticsEventSerde) SerializeObj(value interface{}, b []byte) ([]byte, error) {
	v, ok := value.(*types.AnalyticsEvent)
	if !ok {
		return nil, fmt.Errorf("value is not *types.AnalyticsEvent")
	}
	return s.Serialize(v, b)
}

func (s *AnalyticsEventSerde) DeserializeObj(data []byte) (interface{}, error) {
	return s.Deserialize(data)
}

func (s *AnalyticsEventSerde) Serialize(value *types.AnalyticsEvent, b []byte) ([]byte, error) {
	// TODO: Need to be implemented
	return nil, fmt.Errorf("serialize method for the 'AnalyticsEventSerde' class is not implemented")
}

func (s *AnalyticsEventSerde) Deserialize(data []byte) (*types.AnalyticsEvent, error) {
	// TODO: Need to be implemented
	return nil, fmt.Errorf("deserialize method for the 'AnalyticsEventSerde' class is not implemented")
}
