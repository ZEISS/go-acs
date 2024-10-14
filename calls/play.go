package calls

import (
	"context"
	"fmt"
)

// CallMediaPlayRequest is the body for playing media.
type CallMediaPlayRequest struct {
	PlaySources                 []PlaySource              `json:"playSources"`
	InterruptCallMediaOperation bool                      `json:"interruptCallMediaOperation,omitempty"`
	OperationCallbackUri        string                    `json:"operationCallbackUri,omitempty"`
	OperationContext            string                    `json:"operationContext,omitempty"`
	PlayOptions                 *PlayOptions              `json:"playOptions,omitempty"`
	PlayTo                      []CommunicationIdentifier `json:"playTo,omitempty"`
}

// PlayOptions is the options for playing media.
type PlayOptions struct {
	// Loop is the loop for playing media.
	Loop bool `json:"loop"`
}

// PlaySource is the source for playing media.
type PlaySource struct {
	Kind              PlaySourceType `json:"kind"`
	PlaySourceCacheID string         `json:"playSourceCacheId,omitempty"`
	File              *FileSource    `json:"file,omitempty"`
	SSMLSource        *SSMLSource    `json:"ssml,omitempty"`
	TextSource        *TextSource    `json:"text,omitempty"`
}

// FileSource is the source for playing file.
type FileSource struct {
	// URI is the URI for file.
	URI string `json:"uri"`
}

// SSMLSource is the source for playing SSML.
type SSMLSource struct {
	// CustomVoiceEndpointID is the custom voice endpoint ID.
	CustomVoiceEndpointID string `json:"customVoiceEndpointId,omitempty"`
	// SSMLText is the SSML text.
	SSMLText string `json:"ssmlText"`
}

// TextSource is the source for playing text.
type TextSource struct {
	// CustomVoiceEndpointID is the custom voice endpoint ID.
	CustomVoiceEndpointID string `json:"customVoiceEndpointId,omitempty"`
	// SourceLocale is the source locale.
	SourceLocale string `json:"sourceLocale,omitempty"`
	// Text is the text.
	Text string `json:"text"`
	// VoiceKind is the voice kind.
	VoiceKind VoiceKind `json:"voiceKind"`
	// VoiceName is the voice name.
	VoiceName string `json:"voiceName"`
}

// VoiceKind is the kind for voice.
type VoiceKind string

const (
	// VoiceKindMale is a male voice.
	VoiceKindMale VoiceKind = "male"
	// VoiceKind
	VoiceKindFemale VoiceKind = "female"
)

// PlaySourceType is the type for play source.
type PlaySourceType string

const (
	// PlaySourceTypeFile is the file type for play source.
	PlaySourceTypeFile PlaySourceType = "file"
	// PlaySourceTypeSSML is the SSML type for play source.
	PlaySourceTypeSSML PlaySourceType = "ssml"
	// PlaySourceTypeText is the text type for play source.
	PlaySourceTypeText PlaySourceType = "text"
)

// CallMediaPlay is the call media play.
func (s *Service) CallMediaPlay(ctx context.Context, id string, body *CallMediaPlayRequest) error {
	_, err := s.client.New().Post(fmt.Sprintf("/calling/callConnections/%s:play", id)).BodyJSON(body).ReceiveSuccess(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
