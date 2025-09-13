package authen

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

// Authenticator is an OAuth2 authenticator
type Authenticator struct {
	oauth2.Config
	*oidc.Provider
}

func New() (*Authenticator, error) {
	endpoint := fmt.Sprintf("https://%s/", os.Getenv("AUTH0_DOMAIN"))
	provider, err := oidc.NewProvider(context.Background(), endpoint)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("AUTH0_CALLBACK_URL"),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return &Authenticator{
		Config:   conf,
		Provider: provider,
	}, nil
}

// ParseIDToken extracts and verifies the ID Token from the OAuth2 token.
func (a *Authenticator) ParseIDToken(ctx context.Context, authToken *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := authToken.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("auth token does not contain an id token")
	}

	conf := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(conf).Verify(ctx, rawIDToken)
}
