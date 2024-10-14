package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zeiss/go-acs"
	"github.com/zeiss/go-acs/identities"
)

var (
	endpointURL string = ""
	key         string = ""
)

func main() {
	client := http.Client{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	acsClient := acs.New(endpointURL, key, &client)

	res, err := acsClient.Identity.CreateIdentity(ctx, &identities.CreateIdentityRequestBody{
		CreateTokenWithScopes: []identities.CommunicationIdentityTokenScope{identities.CommunicationIdentityTokenScopeVoip},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
