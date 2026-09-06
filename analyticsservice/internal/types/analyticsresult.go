package types

// Output of the canonical analytics joins. Fields: Key AnalyticsKey, Total int, Kind string.
type AnalyticsResult struct {
	Key   string `json:"key"`
	Total int    `json:"total"`
	Kind  string `json:"kind"`
}
