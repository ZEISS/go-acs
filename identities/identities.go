package identities

import (
	"context"
	"time"

	"github.com/zeiss/go-acs/client"
)

// Service is the service for identity.
type Service struct {
	client *client.Client
}

// CreateIdentityRequestBody is the request body for creating an identity.
type CreateIdentityRequestBody struct {
	// Identity is the identity of the request.
	CreateTokenWithScopes []CommunicationIdentityTokenScope `json:"createTokenWithScopes"`
	// ExpiresInMinutes is the expiration time of the request.
	ExpiresInMinutes int `json:"expiresInMinutes"`
}

// CommunicationIdentityTokenScope is the scope of the token.
type CommunicationIdentityTokenScope string

// CommunicationIdentityTokenScopeChat is the chat scope.
const CommunicationIdentityTokenScopeChat CommunicationIdentityTokenScope = "chat"

// CommunicationIdentityTokenScopeChatJoin is the chat join scope.
const CommunicationIdentityTokenScopeChatJoin CommunicationIdentityTokenScope = "chat.join"

// CommunicationIdentityTokenScopeChatJoinLimited is the chat join limited scope.
const CommunicationIdentityTokenScopeChatJoinLimited CommunicationIdentityTokenScope = "chat.join.limited"

// CommunicationIdentityTokenScopeVoip is the voip scope.
const CommunicationIdentityTokenScopeVoip CommunicationIdentityTokenScope = "voip"

// CommunicationIdentityTokenScopeVoipJoin is the voip join scope.
const CommunicationIdentityTokenScopeVoipJoin CommunicationIdentityTokenScope = "voip.join"

// NewService returns a new SmsService
func NewService(c *client.Client) *Service {
	return &Service{c}
}

// CommunicationIdentityAccessToken is the access token of the identity.
type CommunicationIdentityAccessToken struct {
	ExpiresOn time.Time `json:"expiresOn"`
	Token     string    `json:"token"`
}

// CommunicationIdentity is the identity of the request.
type CommunicationIdentity struct {
	ID string `json:"id"`
}

// CommunicationIdentityAccessTokenResult is the result of the access token.
type CommunicationIdentityAccessTokenResult struct {
	AccessToken CommunicationIdentityAccessToken `json:"accessToken"`
	Identity    CommunicationIdentity            `json:"identity"`
}

// CreateIdentity creates an identity.
func (s *Service) CreateIdentity(ctx context.Context, key string, body *CreateIdentityRequestBody) (*CommunicationIdentityAccessTokenResult, error) {
	res := &CommunicationIdentityAccessTokenResult{}

	err := s.client.Post(ctx, key, "/identities", "api-version=2021-03-07", body, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
