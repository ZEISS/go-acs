package acs

import (
	"net/http"

	"github.com/zeiss/go-acs/calls"
	"github.com/zeiss/go-acs/client"
	"github.com/zeiss/go-acs/identities"
	"github.com/zeiss/go-acs/sms"
)

// Client is the client for the ACS API.
type Client struct {
	SMS      *sms.Service
	Call     *calls.Service
	Identity *identities.Service
}

// New creates a new Client.
func New(endpointURL, key string, c *http.Client) *Client {
	base := client.New().
		Client(c).
		Base(endpointURL).
		QueryStruct(client.DefaultVersion).
		SignProvider(client.NewHMacSigner(key))

	return &Client{
		SMS:      sms.NewService(base),
		Identity: identities.NewService(base),
		Call:     calls.NewService(base),
	}
}
