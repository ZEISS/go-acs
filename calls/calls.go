package calls

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/zeiss/carry"
)

// Service is the service for call.
type Service struct {
	client *carry.Client
}

// NewService returns a new CallService
func NewService(c *carry.Client) *Service {
	return &Service{c}
}

// Opt is a type for options.
type Opt func(*resty.Request)

// CreateCallRequest is the body for creating a call.
type CreateCallRequest struct {
	// CallIntelligenceOptions is the options for call intelligence.
	CallIntelligenceOptions *CallIntelligenceOptions `json:"callIntelligenceOptions,omitempty"`
	// CallbackUri is the callback uri.
	CallbackUri string `json:"callbackUri"`
	// MediaStreaminOptions is the options for media streaming.
	MediaStreaminOptions *MediaStreamingOptions `json:"mediaStreamingOptions,omitempty"`
	// OperationContext is the operation context.
	OperationContext string `json:"operationContext,omitempty"`
	// Source is the source.
	Source *CommunicationIdentifier `json:"source"`
	// SourceCallerIdNumber is the source caller id number.
	SourceCallerIdNumber *PhonenumberIdentifier `json:"sourceCallerIdNumber,omitempty"`
	// SourceDisplayName is the source display name.
	SourceDisplayName string `json:"sourceDisplayName,omitempty"`
	// Targets is the targets.
	Targets []CommunicationIdentifier `json:"targets"`
}

// CallConnectionState is the state for call connection.
type CallConnectionState string

const (
	// CallConnectionStateConnected is the connected state.
	CallConnectionStateConnected CallConnectionState = "connected"
	// CallConnectionStateConnecting is the connecting state.
	CallConnectionStateConnecting CallConnectionState = "connecting"
	// CallConnectionStateDisconnected is the disconnected state.
	CallConnectionStateDisconnected CallConnectionState = "disconnected"
	// CallConnectionStateDisconnecting is the disconnecting state.
	CallConnectionStateDisconnecting CallConnectionState = "disconnecting"
	// CallConnectionStateTransferAccepted is the transfer accepted state.
	CallConnectionStateTransferAccepted CallConnectionState = "transferAccepted"
	// CallConnectionStateTransferring is the transferring state.
	CallConnectionStateTransferring CallConnectionState = "transferring"
	// CallConnectionStateUnknown is the unknown state.
	CallConnectionStateUnknown CallConnectionState = "unknown"
)

// CreateCallResponse is the response for creating a call.
type CreateCallResponse struct {
	// AnsweredBy is the answered by.
	AnsweredBy CommunicationIdentifier `json:"answeredBy"`
	// AnsweredFor is the answered for.
	AnsweredFor PhonenumberIdentifier `json:"answeredFor"`
	// CallConnectionId is the call connection id.
	CallConnectionId string `json:"callConnectionId"`
	// CallConnectionState is the call connection state.
	CallConnectionState CallConnectionState `json:"callConnectionState"`
	// CallbackURI is the callback uri.
	CallbackURI string `json:"callbackUri"`
	// CorrelationId is the correlation id.
	CorrelationId string `json:"correlationId"`
	// MediaStreamingSubscription is the media streaming subscription.
	MediaStreamingSubscription MediaStreamingSubscription `json:"mediaStreamingSubscription"`
	// ServerCallId is the server call id.
	ServerCallId string `json:"serverCallId"`
	// Source is the source.
	Source CommunicationIdentifier `json:"source"`
	// SourceCallerIdNumber is the source caller id number.
	SourceCallerIdNumber PhonenumberIdentifier `json:"sourceCallerIdNumber"`
	// SourceDisplayName is the source display name.
	SourceDisplayName string `json:"sourceDisplayName"`
	// Targets is the targets.
	Targets []CommunicationIdentifier `json:"targets"`
	// TranscriptionSubscription is the transcription subscription.
	TranscriptionSubscription string `json:"transcriptionSubscription"`
}

// MediaStreamingSubscription is the media streaming subscription.
type MediaStreamingSubscription struct {
	// Id is the id.
	ID string `json:"id"`
	// State is the state.
	State MediaStreamingSubscriptionState `json:"state"`
	// SubscribedContentTypes is the subscribed content types.
	SubscribedContentTypes []MediaStreamingContentType `json:"subscribedContentTypes"`
}

// TranscriptionSubscription is the transcription subscription.
type TranscriptionSubscription struct {
	// Id is the id.
	ID string `json:"id"`
	// State is the state.
	State TranscriptionSubscriptionState `json:"state"`
	// SubscribedResultTypes is the subscribed result types.
	SubscribedResultTypes []TranscriptionResultType `json:"subscribedResultTypes"`
}

// TranscriptionSubscriptionState is the state for transcription subscription.
type TranscriptionSubscriptionState string

const (
	// TranscriptionSubscriptionStateActive is the active state.
	TranscriptionSubscriptionStateActive TranscriptionSubscriptionState = "active"
	// TranscriptionSubscriptionStateDisabled is the disabled state.
	TranscriptionSubscriptionStateDisabled TranscriptionSubscriptionState = "disabled"
	// TranscriptionSubscriptionStateInactive is the inactive state.
	TranscriptionSubscriptionStateInactive TranscriptionSubscriptionState = "inactive"
)

// TranscriptionResultType is the type for transcription result.
type TranscriptionResultType struct {
	// Final is the final result.
	Final string `json:"final"`
	// Intermediate is the intermediate result.
	Intermediate string `json:"intermediate"`
}

// MediaStreamingSubscriptionState is the state for media streaming subscription.
type MediaStreamingSubscriptionState string

const (
	// MediaStreamingSubscriptionStateActive is the active state.
	MediaStreamingSubscriptionStateActive MediaStreamingSubscriptionState = "active"
	// MediaStreamingSubscriptionStateInactive is the inactive state.
	MediaStreamingSubscriptionStateInactive MediaStreamingSubscriptionState = "inactive"
	// MediaStreamingSubscriptionStateDisabled is the disabled state.
	MediaStreamingSubscriptionStateDisabled MediaStreamingSubscriptionState = "disabled"
)

// CommunicationIdentifier is a communication user identifier.
type CommunicationIdentifier struct {
	ID                string                 `json:"id,omitempty"`
	Kind              string                 `json:"kind"`
	CommunicationUser *CommunicationUser     `json:"communicationUser,omitempty"`
	PhoneNumber       *PhonenumberIdentifier `json:"phoneNumber,omitempty"`
}

// CallIntelligenceOptions is the options for call intelligence.
type CallIntelligenceOptions struct {
	CognitiveServicesEndpoint string `json:"cognitiveServicesEndpoint,omitempty"`
}

// MediaStreaminOptions is the options for media streaming.
type MediaStreamingOptions struct {
	// AudioChannelType is the type for audio channel.
	AudioChannelType MediaStreamingAudioChannelType `json:"audioChannelType"`
	// ContentType is the type for content.
	ContentType MediaStreamingContentType `json:"contentType"`
	// StartMediaStreaming is the start media streaming.
	StartMediaStreaming bool `json:"startMediaStreaming"`
	// TransportType is the type for media streaming transport.
	TransportType MediaStreamingTransportType `json:"transportType"`
	// TransportUrl is the url for media streaming transport.
	TransportUrl string `json:"transportUrl"`
}

// MediaStreamingAudioChannelType is the type for audio channel.
type MediaStreamingAudioChannelType string

const (
	// MediaStreamingAudioChannelTypeMixed is the mixed audio channel type.
	MediaStreamingAudioChannelTypeMixed MediaStreamingAudioChannelType = "mixed"
	// MediaStreamingAudioChannelTypeUnmixed
	MediaStreamingAudioChannelTypeUnmixed MediaStreamingAudioChannelType = "unmixed"
)

// MediaStreamingContentType is the type for content.
type MediaStreamingContentType string

const (
	// MediaStreamingContentTypeAudio is the audio content type.
	MediaStreamingContentTypeAudio MediaStreamingContentType = "audio"
)

// MediaStreamingTransportType is the type for media streaming transport.
type MediaStreamingTransportType string

const (
	// MediaStreamingTransportTypeWebsocket is the websocket transport type.
	MediaStreamingTransportTypeWebsocket MediaStreamingTransportType = "websocket"
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

// TranscriptionOptions is the options for transcription.
type TranscriptionOptions struct {
	// EnableInterimResults is the flag to enable interim results.
	EnableIntermediateResults bool `json:"enableIntermediateResults"`
	// Locale is the locale.
	Locale string `json:"locale"`
	// SpeechRecognitionModelEndpointId is the speech recognition model endpoint id.
	SpeechRecognitionModelEndpointId string `json:"speechRecognitionModelEndpointId"`
	// StartTranscription is the flag to start transcription.
	StartTranscription bool `json:"startTranscription"`
	// TransportType is the type for transcription transport.
	TransportType string `json:"transportType"`
	// TransportUrl is the url for transcription transport.
	TransportUrl string `json:"transportUrl"`
}

// TranscriptionTransportType is the type for transcription transport.
type TranscriptionTransportType string

const (
	// TranscriptionTransportTypeWebsocket is the websocket transport type.
	TranscriptionTransportTypeWebsocket TranscriptionTransportType = "websocket"
)

// CreateCall creates a call.
func (s *Service) CreateCall(ctx context.Context, body *CreateCallRequest) (*CreateCallResponse, error) {
	res := &CreateCallResponse{}

	_, err := s.client.New().Post("/calling/callConnections").BodyJSON(body).ReceiveSuccess(ctx, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// CallHangUp is the method for recognizing call.
func (s *Service) CallHangUp(ctx context.Context, id string) error {
	_, err := s.client.New().Delete(fmt.Sprintf("/calling/callConnections/%s", id)).ReceiveSuccess(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
