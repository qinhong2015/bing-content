package bingcontent

import (
	"context"
	"golang.org/x/oauth2"
)

type Session struct {
	AccountId              string
	CustomerId             string
	DeveloperToken         string
	Username               string
	Password               string
	withRequireLiveConnect bool
	HTTPClient             HttpClient
	TokenSource            oauth2.TokenSource
}

type SessionConfig struct {
	OAuth2Config   *oauth2.Config
	OAuth2Token    *oauth2.Token
	AccountId      string
	CustomerId     string
	DeveloperToken string
	HTTPClient     HttpClient
}

func NewSession(config SessionConfig) *Session {
	tokenSource := config.OAuth2Config.TokenSource(context.TODO(), config.OAuth2Token)

	return &Session{
		AccountId:      config.AccountId,
		CustomerId:     config.CustomerId,
		DeveloperToken: config.DeveloperToken,
		HTTPClient:     config.HTTPClient,
		TokenSource:    tokenSource,
	}
}
