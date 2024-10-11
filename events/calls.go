package events

// MicrosoftCommunicationCallConnected is the data type of the event.
// This parses the data of the Microsoft.Communication.CallConnected event.
type MicrosoftCommunicationCallConnected struct {
	// Version is the version of the event.
	Version string `json:"version"`
	// CallConnectionID is the ID of the call connection.
	CallConnectionID string `json:"callConnectionId"`
	// ServerCallID is the ID of the server call.
	ServerCallID string `json:"serverCallId"`
	// CorrelationID is the ID of the correlation.
	CorrelationID string `json:"correlationId"`
	// PublicEventType is the type of the event.
	PublicEventType string `json:"publicEventType"`
}

// MicrosoftCommunicationParticipantsUpdated is the data type of the event.
// This parses the data of the Microsoft.Communication.ParticipantsUpdate event.
type MicrosoftCommunicationParticipantsUpdated struct {
	// Participants is the list of participants.
	Participants []Participant `json:"participants"`
	// SequenceNumber is the sequence number.
	SequenceNumber int `json:"sequenceNumber"`
	// Version is the version of the event.
	Version string `json:"version"`
	// CallConnectionID is the ID of the call connection.
	CallConnectionID string `json:"callConnectionId"`
	// ServerCallID is the ID of the server call.
	ServerCallID string `json:"serverCallId"`
	// CorrelationID is the ID of the correlation.
	CorrelationID string `json:"correlationId"`
	// PublicEventType is the type of the event.
	PublicEventType string `json:"publicEventType"`
}

// Participant is the participant.
type Participant struct {
	// Identifier is the identifier of the participant.
	Identifier CommunicationIdentifier `json:"identifier"`
	// IsMuted is the flag of the participant.
	IsMuted bool `json:"isMuted"`
	// IsOnHold is the flag of the participant.
	IsOnHold bool `json:"isOnHold"`
}

// CommunicationIdentifier is a communication user identifier.
type CommunicationIdentifier struct {
	ID                string                      `json:"id,omitempty"`
	Kind              CommunicationIdentifierKind `json:"kind"`
	CommunicationUser *CommunicationUser          `json:"communicationUser,omitempty"`
	PhoneNumber       *PhonenumberIdentifier      `json:"phoneNumber,omitempty"`
}

// CommunicationIdentifierKind is the kind of the communication identifier.
type CommunicationIdentifierKind string

const (
	// CommunicationIdentifierKindCommunicationUser is the communication user kind.
	CommunicationIdentifierKindCommunicationUser CommunicationIdentifierKind = "communicationUser"
	// CommunicationIdentifierKindPhoneNumber is the phone number kind.
	CommunicationIdentifierKindPhoneNumber CommunicationIdentifierKind = "phoneNumber"
)

// PhonenumberIdentifier is the phone number identifier.
type PhonenumberIdentifier struct {
	// ID is the id
	ID string `json:"id"`
	// Value is the value.
	Value string `json:"value"`
}

// CommunicationUser is a communication user.
type CommunicationUser struct {
	ID string `json:"id"`
}

// MicrosoftCommunicationRecognizeCompleted is the data type of the event.
// This parses the data of the Microsoft.Communication.RecognizeCompleted event.
type MicrosoftCommunicationRecognizeCompleted struct {
	// RecognitionType is the type of recognition.
	RecognitionType RecognizeInputType `json:"recognitionType"`
	// ChoiceResult is the result of choice.
	ChoiceResult *ChoiceResult `json:"choiceResult,omitempty"`
	// Version is the version of the event.
	Version string `json:"version"`
	// ResultInformation is the information of the result.
	ResultInformation *ResultInformation `json:"resultInformation,omitempty"`
	// CallConnectionID is the ID of the call connection.
	CallConnectionID string `json:"callConnectionId"`
	// ServerCallID is the ID of the server call.
	ServerCallID string `json:"serverCallId"`
	// CorrelationID is the ID of the correlation.
	CorrelationID string `json:"correlationId"`
	// PublicEventType is the type of the event.
	PublicEventType string `json:"publicEventType"`
}

// RecognizeInputType is the type of input for recognizing call.
type RecognizeInputType string

const (
	// RecognizeInputTypeChoices is the type of input for recognizing call.
	RecognizeInputTypeChoices RecognizeInputType = "choices"
	// RecognizeInputTypeDtmf is the type of input for recognizing call.
	RecognizeInputTypeDtmf RecognizeInputType = "dtmf"
	// RecognizeInputTypeSpeech is the type of input for recognizing call.
	RecognizeInputTypeSpeech RecognizeInputType = "speech"
	// RecognizeInputTypeSpeechOrDtmf is the type of input for recognizing call.
	RecognizeInputTypeSpeechOrDtmf RecognizeInputType = "speechOrDtmf"
)

// ChoiceResult is the result for choice.
type ChoiceResult struct {
	// Label is the label of the choice.
	Label string `json:"label"`
}

// ResultInformation is the information for result.
type ResultInformation struct {
	// Code is the code of the result.
	Code int `json:"code"`
	// SubCode is the sub code of the result.
	SubCode int `json:"subCode"`
	// Message is the message of the result.
	Message string `json:"message"`
}
