package calls

// CallConnectedData is the data for call connected.
type CallConnectedData struct {
	Version          string `json:"version"`
	CallConnectionID string `json:"callConnectionId"`
	ServerCallID     string `json:"serverCallId"`
	CorrelationID    string `json:"correlationId"`
	PublicEventType  string `json:"publicEventType"`
}

type ParticipantsUpdatedData struct {
	Participants     []Participant `json:"participants"`
	SequenceNumber   int           `json:"sequenceNumber"`
	Version          string        `json:"version"`
	CallConnectionID string        `json:"callConnectionId"`
	ServerCallID     string        `json:"serverCallId"`
	CorrelationID    string        `json:"correlationId"`
	PublicEventType  string        `json:"publicEventType"`
}

// Participant is the participant.
type Participant struct {
	Identifier CommunicationIdentifier `json:"identifier"`
	IsMuted    bool                    `json:"isMuted"`
	IsOnHold   bool                    `json:"isOnHold"`
}

type CallRecognizeCompletedData struct {
	RecognitionType   RecognizeInputType `json:"recognitionType"`
	ChoiceResult      *ChoiceResult      `json:"choiceResult,omitempty"`
	Version           string             `json:"version"`
	ResultInformation *ResultInformation `json:"resultInformation,omitempty"`
	CallConnectionID  string             `json:"callConnectionId"`
	ServerCallID      string             `json:"serverCallId"`
	CorrelationID     string             `json:"correlationId"`
	PublicEventType   string             `json:"publicEventType"`
}

// ChoiceResult is the result for choice.
type ChoiceResult struct {
	Label string `json:"label"`
}

// ResultInformation is the information for result.
type ResultInformation struct {
	Code    int    `json:"code"`
	SubCode int    `json:"subCode"`
	Message string `json:"message"`
}
