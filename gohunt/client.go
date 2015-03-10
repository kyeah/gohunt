// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"bytes"
	"encoding/json"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
)

type Client struct {
	AuthToken     *Token
	Authorization string
}

type Token struct {
	AccessToken string  `json:"access_token"`
	TokenType   string  `json:"token_type"`
	Expiry      float32 `json:"expires_in"`
	Scope       string  `json:"scope"`
}

// User-Authenticated Client with Developer Token
func NewUserClient(accessToken string) *Client {
	tok := &Token{
		AccessToken: accessToken,
		TokenType:   "bearer",
	}

	return GenAuthClient(tok)
}

// Request Access Grant Code and send to redirectUrl
func RequestUserOAuthCode(w http.ResponseWriter, r *http.Request, clientID string, redirectUrl string, state string) {
	var (
		host   = "api.producthunt.com"
		base   = "https://" + host
		config = &oauth2.Config{
			ClientID:    clientID,
			Scopes:      []string{"public", "private"},
			RedirectURL: redirectUrl,
			Endpoint: oauth2.Endpoint{
				AuthURL:  base + "/v1/oauth/authorize",
				TokenURL: base + "/v1/oauth/token",
			},
		}
	)

	reqUrl := config.AuthCodeURL(state, oauth2.AccessTypeOnline)
	http.Redirect(w, r, reqUrl, http.StatusFound)
}

// OAuth2 User-Authenticated Client with Access Grant Code
func NewUserOAuthClient(clientID string, clientSecret string, redirectUrl string, code string) (*Client, error) {
	var (
		host   = "api.producthunt.com"
		base   = "https://" + host
		config = &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       []string{"public", "private"},
			RedirectURL:  redirectUrl,
			Endpoint: oauth2.Endpoint{
				AuthURL:  base + "/v1/oauth/authorize",
				TokenURL: base + "/v1/oauth/token",
			},
		}
	)

	otok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}

	tok := &Token{
		AccessToken: otok.AccessToken,
		TokenType:   otok.TokenType,
	}

	return GenAuthClient(tok), nil
}

// OAuth2 Client-Only Authentication
func NewOAuthClient(clientID string, clientSecret string) (*Client, error) {
	var (
		host     = "api.producthunt.com"
		base     = "https://" + host
		tokenURL = base + "/v1/oauth/token"
		req      = Request{
			url:    tokenURL,
			action: "POST",
			values: &url.Values{
				"grant_type":    {"client_credentials"},
				"client_id":     {clientID},
				"client_secret": {clientSecret},
			},
		}
	)

	response, err := req.getResponse()
	if err != nil {
		return nil, err
	}

	tok := &Token{}
	err = json.NewDecoder(response).Decode(tok)
	if err != nil {
		return nil, err
	}

	return GenAuthClient(tok), nil
}

func GenAuthClient(tok *Token) *Client {
	var buffer bytes.Buffer
	buffer.WriteString("Bearer ")
	buffer.WriteString(tok.AccessToken)

	return &Client{
		AuthToken:     tok,
		Authorization: buffer.String(),
	}
}

func (c *Client) sendRequest(req *Request) (*bytes.Buffer, error) {
	return req.getAuthResponse(c.Authorization)
}
