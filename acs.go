package acs

import (
	"net/http"

	"github.com/zeiss/carry"
	"github.com/zeiss/go-acs/calls"
	"github.com/zeiss/go-acs/identities"
	"github.com/zeiss/go-acs/sms"
)

// DefaultVersion
var DefaultVersion = struct {
	APIVersion string `url:"api-version"`
}{
	APIVersion: "2024-06-15-preview",
}

// Client is the client for the ACS API.
type Client struct {
	SMS      *sms.Service
	Call     *calls.Service
	Identity *identities.Service
}

// New creates a new Client.
func New(endpointURL, key string, c *http.Client) *Client {
	base := carry.New().
		Client(c).
		Base(endpointURL).
		QueryStruct(DefaultVersion).
		SignProvider(carry.NewHMacSigner(key))

	return &Client{
		SMS:      sms.NewService(base),
		Identity: identities.NewService(base),
		Call:     calls.NewService(base),
	}
}
