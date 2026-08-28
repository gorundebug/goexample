package serdes

import (
	"encoding/json"
	"fmt"

	"github.com/gorundebug/model_go/pkg/types"
)

type OrderItemSerde struct{}

func (s *OrderItemSerde) IsStub() bool {
	return false
}

func (s *OrderItemSerde) SerializeObj(value interface{}, b []byte) ([]byte, error) {
	v, ok := value.(*types.OrderItem)
	if !ok {
		return nil, fmt.Errorf("value is not *types.OrderItem")
	}
	return s.Serialize(v, b)
}

func (s *OrderItemSerde) DeserializeObj(data []byte) (interface{}, error) {
	return s.Deserialize(data)
}

func (s *OrderItemSerde) Serialize(value *types.OrderItem, _ []byte) ([]byte, error) {
	return json.Marshal(value)
}

func (s *OrderItemSerde) Deserialize(data []byte) (*types.OrderItem, error) {
	var value types.OrderItem
	if err := json.Unmarshal(data, &value); err != nil {
		return nil, err
	}
	return &value, nil
}
