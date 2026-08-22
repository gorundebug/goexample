package types

// A single line item within an order, emitted independently into the pipeline for parallel inventory checks. Fields:
// OrderID string (parent order correlation), ItemID string, SKU string, Quantity int, UnitPrice float64.
type OrderItem struct {
	OrderID   string
	ItemID    string
	SKU       string
	Quantity  int
	UnitPrice float64
}
