package calls

// CallConnectedData is the data for call connected.
type CallConnectedData struct {
	Version          string `json:"version"`
	CallConnectionID string `json:"callConnectionId"`
	ServerCallID     string `json:"serverCallId"`
	CorrelationID    string `json:"correlationId"`
	PublicEventType  string `json:"publicEventType"`
}
