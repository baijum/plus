package plus

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Token represents an OAuth token response.
type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	IDToken     string `json:"id_token"`
}

// ClaimSet represents an IDToken response.
type ClaimSet struct {
	Sub string
}

// GetTokens takes an authentication code, client ID & client secret
// and exchanges it with the OAuth endpoint for a Google API bearer
// token and a Google+ ID token
func GetTokens(code, clientID, clientSecret string) (accessToken, idToken string, err error) {
	// Exchange the authorization code for a credentials object via a POST request
	addr := "https://accounts.google.com/o/oauth2/token"
	values := url.Values{
		"Content-Type":  {"application/x-www-form-urlencoded"},
		"code":          {code},
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"redirect_uri":  {"postmessage"},
		"grant_type":    {"authorization_code"},
	}

	resp, err := http.PostForm(addr, values)
	if err != nil {
		return "", "", fmt.Errorf("exchanging code: %v", err)
	}

	defer resp.Body.Close()

	// Decode the response body into a token object
	var token Token
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return "", "", fmt.Errorf("decoding access token: %v", err)
	}

	return token.AccessToken, token.IDToken, nil
}

// DecodeIDToken takes an ID Token and decodes it to fetch the Google+ ID within
func DecodeIDToken(idToken string) (gplusID string, err error) {
	// An ID token is a cryptographically-signed JSON object encoded in base 64.
	// Normally, it is critical that you validate an ID token before you use it,
	// but since you are communicating directly with Google over an
	// intermediary-free HTTPS channel and using your Client Secret to
	// authenticate yourself to Google, you can be confident that the token you
	// receive really comes from Google and is valid. If your server passes the ID
	// token to other components of your app, it is extremely important that the
	// other components validate the token before using it.
	var set ClaimSet
	if idToken != "" {
		// Check that the padding is correct for a base64decode
		parts := strings.Split(idToken, ".")
		if len(parts) < 2 {
			return "", fmt.Errorf("malformed ID token")
		}
		// Decode the ID token
		b, err := base64Decode(parts[1])
		if err != nil {
			return "", fmt.Errorf("malformed ID token: %v", err)
		}
		err = json.Unmarshal(b, &set)
		if err != nil {
			return "", fmt.Errorf("malformed ID token: %v", err)
		}
	}
	return set.Sub, nil
}

func base64Decode(s string) ([]byte, error) {
	// add back missing padding
	log.Println("Hello")
	switch len(s) % 4 {
	case 2:
		s += "=="
	case 3:
		s += "="
	}
	return base64.URLEncoding.DecodeString(s)
}
