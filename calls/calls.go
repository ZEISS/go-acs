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
	// Source *CommunicationIdentifier `json:"source"`
	// SourceCallerIdNumber is the source caller id number.
	SourceCallerIdNumber *PhonenumberIdentifier `json:"sourceCallerIdNumber,omitempty"`
	// SourceDisplayName is the source display name.
	SourceDisplayName string `json:"sourceDisplayName,omitempty"`
	// Targets is the targets.
	Targets []CommunicationIdentifier `json:"targets"`
}

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
func (s *Service) CreateCall(ctx context.Context, body *CreateCallRequest) error {
	res, err := s.client.New().Post("/calling/callConnections").BodyJSON(body).ReceiveSuccess(ctx, nil)
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}

// CallHangUp is the method for recognizing call.
func (s *Service) CallHangUp(ctx context.Context, id string) error {
	_, err := s.client.New().Delete(fmt.Sprintf("/calling/callConnections/%s", id)).ReceiveSuccess(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
