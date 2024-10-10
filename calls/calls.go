package calls

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/zeiss/go-acs/client"
)

// Service is the service for call.
type Service struct {
	client *client.Client
}

// NewService returns a new CallService
func NewService(c *client.Client) *Service {
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

// WithAuthToken sets the auth token.
func WithAuthToken(token string) Opt {
	return func(r *resty.Request) {
		r.SetAuthToken(token)
	}
}

// CreateCall creates a call.
func (s *Service) CreateCall(ctx context.Context, key string, req *CreateCallRequest, opts ...client.Opt) error {
	return s.client.Post(ctx, key, "/calling/callConnections", "api-version=2024-06-15-preview", req, nil, opts...)
}
