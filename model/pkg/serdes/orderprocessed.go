package serdes

import (
	"encoding/json"
	"fmt"

	"github.com/gorundebug/model/pkg/types"
)

type OrderProcessedSerde struct{}

func (s *OrderProcessedSerde) IsStub() bool {
	return false
}

func (s *OrderProcessedSerde) SerializeObj(value interface{}, b []byte) ([]byte, error) {
	v, ok := value.(*types.OrderProcessed)
	if !ok {
		return nil, fmt.Errorf("value is not *types.OrderProcessed")
	}
	return s.Serialize(v, b)
}

func (s *OrderProcessedSerde) DeserializeObj(data []byte) (interface{}, error) {
	return s.Deserialize(data)
}

func (s *OrderProcessedSerde) Serialize(value *types.OrderProcessed, b []byte) ([]byte, error) {
	return json.Marshal(value)
}

func (s *OrderProcessedSerde) Deserialize(data []byte) (*types.OrderProcessed, error) {
	value := &types.OrderProcessed{}
	if err := json.Unmarshal(data, value); err != nil {
		return nil, err
	}
	return value, nil
}
