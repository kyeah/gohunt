Gohunt
========

A golang client library for the official Product Hunt API.

## Usage

### Gohunt Client

Interaction with the Product Hunt API is facilitated by the Gohunt Client. The client can be generated in three ways:

1. Client-Only Authentication by OAuth2
```go
client, err := gohunt.NewOAuthClient(clientID, clientSecret)
```

2. User-Authentication by Developer Token
```go
client, err := gohunt.NewUserClient(phToken)
```

3. User-Authentication by OAuth2
```go
func HandleLogin() {
   err := gohunt.RequestUserOAuthCode(clientID, redirectUrl, state)
}

func HandleRedirect(data) {
   client, err := gohunt.NewUserOAuthClient(clientID, clientSecret, redirectUrl, data.code)
}
```

### Client Interface
