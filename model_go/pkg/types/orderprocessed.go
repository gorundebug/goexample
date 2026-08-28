package types

import "time"

// Final order-processing event. Fields: OrderID string, Status string, ProcessedAt time.Time, TotalItems int,
// ConfirmedItems int, FailureReason string.
type OrderProcessed struct {
	OrderID        string    `json:"order_id"`
	Status         string    `json:"status"`
	ProcessedAt    time.Time `json:"processed_at"`
	TotalItems     int       `json:"total_items"`
	ConfirmedItems int       `json:"confirmed_items"`
	FailureReason  string    `json:"failure_reason,omitempty"`
}
