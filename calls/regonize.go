package calls

import (
	"context"
	"fmt"
)

// CallRecognizeRequest is the body for recognizing call.
type CallRecognizeRequest struct {
	RecognizeInputType          RecognizeInputType `json:"recognizeInputType"`
	RecognizeOptions            *RecognizeOptions  `json:"recognizeOptions,omitempty"`
	InterruptCallMediaOperation bool               `json:"interruptCallMediaOperation,omitempty"`
	OperationCallbackUri        string             `json:"operationCallbackUri"`
	OperationContext            string             `json:"operatonContext"`
	PlayPrompt                  PlaySource         `json:"playPrompt"`
	PlayPrompots                []PlaySource       `json:"playPrompts"`
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

// RecognizeOptions is the options for recognizing call.
type RecognizeOptions struct {
	// Choices is the list of choices for recognizing call.
	Choices []Choice `json:"choices"`
	// DtmfOptions is the options for recognizing DTMF.
	DtmfOptions *DtmfOptions `json:"dtmfOptions,omitempty"`
	// InitialSilenceTimeoutInSeconds is the initial silence timeout in seconds.
	InitialSilenceTimeoutInSeconds int `json:"initialSilenceTimeoutInSeconds"`
	// InterruptPrompt is the prompt for interrupting.
	InterruptPrompt bool `json:"interruptPrompt"`
	// SpeechLanguage is the language for speech.
	SpeechLanguage string `json:"speechLanguage"`
	// SpeechOptions is the options for recognizing speech.
	SpeechOptions *SpeechOptions `json:"speechOptions,omitempty"`
	// SpeechRecognitionModelEndpointId is the endpoint ID for speech recognition model.
	SpeechRecognitionModelEndpointId string `json:"speechRecognitionModelEndpointId"`
	// TargetParticipant is the target participant for recognizing.
	TargetParticipant *CommunicationIdentifier `json:"targetParticipant"`
}

// Choice is the choice for recognizing call.
type Choice struct {
	// Label is the label for choice.
	Label string `json:"label"`
	// Phrase is the phrase for choice.
	Phrases []string `json:"phrases"`
	// Tone is the tone for choice.
	Tone Tone `json:"tone"`
}

// Tone is the tone for choice.
type Tone string

const (
	// ToneA is the a tone for choice.
	ToneA Tone = "a"
	// ToneB is the b tone for choice.
	ToneB Tone = "b"
	// ToneC is the c tone for choice.
	ToneC Tone = "c"
	// ToneD is the d tone for choice.
	ToneD Tone = "d"
	// ToneEight is the eight tone for choice.
	ToneEight Tone = "eight"
	// ToneFive is the five tone for choice.
	ToneFive Tone = "five"
	// ToneFour is the four tone for choice.
	ToneFour Tone = "four"
	// ToneNine is the nine tone for choice.
	ToneNine Tone = "nine"
	// ToneOne is the one tone for choice.
	ToneOne Tone = "one"
	// ToneSeven is the seven tone for choice.
	ToneSeven Tone = "seven"
	// ToneSix is the six tone for choice.
	ToneSix Tone = "six"
	// ToneStar is the star tone for choice.
	ToneStar Tone = "star"
	// ToneThree is the three tone for choice.
	ToneThree Tone = "three"
	// ToneTwo is the two tone for choice.
	ToneTwo Tone = "two"
	// ToneZero is the zero tone for choice.
	ToneZero Tone = "zero"
	// TonePound is the pound tone for choice.
	TonePound Tone = "pound"
)

// DtmfOptions is the options for recognizing DTMF.
type DtmfOptions struct {
	// InterDigitTimeoutInSeconds is the inter-digit timeout in seconds.
	InterDigitTimeoutInSeconds int `json:"interDigitTimeoutInSeconds"`
	// MaxTonesToCollect is the max tones to collect.
	MaxTonesToCollect int `json:"maxTonesToCollect"`
	// StopTones is the stop tones.
	StopTones []Tone `json:"stopTones"`
}

// SpeechOptions is the options for recognizing speech.
type SpeechOptions struct {
	// EndSilenceTimeoutInMs is the end silence timeout in milliseconds.
	EndSilenceTimeoutInMs int `json:"endSilenceTimeoutInMs"`
}

// CallMediaRecognize is used to recognize the call.
func (s *Service) CallMediaRecognize(ctx context.Context, id string, key string, body *CallRecognizeRequest) error {
	_, err := s.client.New().Post(fmt.Sprintf("/calling/callConnections/%s:recognize", id)).BodyJSON(body).ReceiveSuccess(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
