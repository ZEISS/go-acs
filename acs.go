package acs

import (
	"net/http"

	"github.com/zeiss/go-acs/client"
	"github.com/zeiss/go-acs/sms"
)

// Client is the client for the ACS API.
type Client struct {
	c   *client.Client
	SMS *sms.Service
}

// New creates a new Client.
func New(endpointURL, key string, c *http.Client) *Client {
	base := client.New(endpointURL, key, c)

	return &Client{
		c:   base,
		SMS: sms.NewService(base),
	}
}
