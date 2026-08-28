package types

// Inventory reservation result for a single order item. Fields: OrderID string (parent order correlation), ItemID
// string, SKU string, RequestedQty int, AvailableQty int (actual stock at reservation time), Reserved bool, Status
// string (CONFIRMED / OUT_OF_STOCK / PROCESSING_ERROR), Error string (transport error text, empty on success).
type OrderItemResult struct {
	OrderID      string
	ItemID       string
	SKU          string
	RequestedQty int
	AvailableQty int
	Reserved     bool
	Status       string
	UnitPrice    float64
	Error        string
}
