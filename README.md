Golang utilities for Google+ Sign-In (server-side apps)
=======================================================

[![GoDoc](https://godoc.org/github.com/baijum/plus?status.svg)](https://godoc.org/github.com/baijum/plus)

Package plus provides utilities for Google+ Sign-In (server-side apps)

Example:
```
accessToken, idToken, err := plus.GetTokens(code, clientID, clientSecret)
if err != nil {
	log.Fatal("Error getting tokens: ", err)
}

gplusID, err := plus.DecodeIDToken(idToken)
if err != nil {
	log.Fatal("Error decoding ID token: ", err)
}
```

The code base is derived from [Google+ quickstart
example](https://github.com/googleplus/gplus-quickstart-go)
