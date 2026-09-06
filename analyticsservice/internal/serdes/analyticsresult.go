package serdes

import (
	"fmt"

	"github.com/gorundebug/analyticsservice/internal/types"
)

type AnalyticsResultSerde struct{}

func (s *AnalyticsResultSerde) IsStub() bool {
	return false
}

func (s *AnalyticsResultSerde) SerializeObj(value interface{}, b []byte) ([]byte, error) {
	v, ok := value.(*types.AnalyticsResult)
	if !ok {
		return nil, fmt.Errorf("value is not *types.AnalyticsResult")
	}
	return s.Serialize(v, b)
}

func (s *AnalyticsResultSerde) DeserializeObj(data []byte) (interface{}, error) {
	return s.Deserialize(data)
}

func (s *AnalyticsResultSerde) Serialize(value *types.AnalyticsResult, b []byte) ([]byte, error) {
	// TODO: Need to be implemented
	return nil, fmt.Errorf("serialize method for the 'AnalyticsResultSerde' class is not implemented")
}

func (s *AnalyticsResultSerde) Deserialize(data []byte) (*types.AnalyticsResult, error) {
	// TODO: Need to be implemented
	return nil, fmt.Errorf("deserialize method for the 'AnalyticsResultSerde' class is not implemented")
}
