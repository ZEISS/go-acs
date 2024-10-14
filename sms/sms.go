package sms

import (
	"context"

	"github.com/zeiss/carry"
)

const DefaultVersion = "2021-03-07"

// Request is the request for sending an SMS.
type Request struct {
	// From is the phone number of the sender formatted as E.164 format.
	From string `json:"from"`
	// SmsRecipients is the recipients of the SMS request.
	SMSRecipients []SMSRecipients `json:"smsRecipients"`
	// Message is the message of the SMS request.
	Message string `json:"message"`
	// SmsSendOptions is the options for sending the SMS request.
	SMSSendOptions SMSSendOptions `json:"smsSendOptions"`
}

// SMSRecipients is the recipients of the SMS request.
type SMSRecipients struct {
	// To is the phone number of the recipient.
	To string `json:"to"`
	// RepeatabilityRequestID is the ID of the request.
	RepeatabilityRequestID string `json:"repeatabilityRequestId"`
	// RepeatabilityFirstSent is the time the request was first sent.
	RepeatabilityFirstSent string `json:"repeatabilityFirstSent"`
}

// SMSSendOptions is the options for sending the SMS request.
type SMSSendOptions struct {
	// EnableDeliveryReport is whether to enable delivery reports.
	EnableDeliveryReport bool `json:"enableDeliveryReport"`
	// Tag is the tag for the request.
	Tag string `json:"tag"`
}

// Response is the response for sending an SMS.
type Response struct {
	Value []SMSSendResponseItem `json:"value"`
}

// SMSSendResponseItem is the response item for sending an SMS.
type SMSSendResponseItem struct {
	// ErrorMessage is the error message of the response.
	ErrorMessage string `json:"errorMessage"`
	// HttpStatusCode is the status code of the response.
	HttpStatusCode int `json:"httpStatusCode"`
	// MessageID is the ID of the message.
	MessageID string `json:"messageId,omitempty"`
	// To is the phone number of the recipient.
	To string `json:"to"`
	// Successful is whether the request was successful.
	Successful bool `json:"successful"`
	// RepeatabilityResult is the result of the request.
	RepeatabilityResult RepeatabilityResult `json:"repeatabilityResult"`
}

// RepeatabilityResult is the result of a repeatability request.
type RepeatabilityResult string

const (
	// RepeatabilityResultAccepted is the result of a successful request.
	RepeatabilityResultAccepted RepeatabilityResult = "accepted"
	// RepeatabilityResultRejected is the result of a failed request.
	RepeatabilityResultRejected RepeatabilityResult = "rejected"
)

// Error is the error for the SMS API.
type Error struct{}

// Service is the service for the SMS API.
type Service struct {
	client *carry.Client
}

// NewService returns a new SmsService
func NewService(c *carry.Client) *Service {
	return &Service{c}
}

// SendSMS sends an SMS message.
func (s *Service) SendSMS(ctx context.Context, request *Request) (*Response, error) {
	result := &Response{}

	_, err := s.client.New().Post("/sms").ReceiveSuccess(ctx, result)
	if err != nil {
		return nil, err
	}

	return result, err
}
