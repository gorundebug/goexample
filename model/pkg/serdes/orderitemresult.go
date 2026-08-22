package serdes

import (
	"encoding/json"
	"fmt"

	"github.com/gorundebug/model/pkg/types"
)

type OrderItemResultSerde struct{}

func (s *OrderItemResultSerde) IsStub() bool {
	return false
}

func (s *OrderItemResultSerde) SerializeObj(value interface{}, b []byte) ([]byte, error) {
	v, ok := value.(*types.OrderItemResult)
	if !ok {
		return nil, fmt.Errorf("value is not *types.OrderItemResult")
	}
	return s.Serialize(v, b)
}

func (s *OrderItemResultSerde) DeserializeObj(data []byte) (interface{}, error) {
	return s.Deserialize(data)
}

func (s *OrderItemResultSerde) Serialize(value *types.OrderItemResult, _ []byte) ([]byte, error) {
	return json.Marshal(value)
}

func (s *OrderItemResultSerde) Deserialize(data []byte) (*types.OrderItemResult, error) {
	var v types.OrderItemResult
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}
