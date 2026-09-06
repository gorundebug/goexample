package types

// Input for the canonical analytics joins. Fields: Key AnalyticsKey, Value int, Kind string.
type AnalyticsEvent struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
	Kind  string `json:"kind"`
}
