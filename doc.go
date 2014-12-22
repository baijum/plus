/*
Package plus provides utilities for Google+ Sign-In (server-side apps)

Examples:

  accessToken, idToken, err := plus.GetTokens(code, clientID, clientSecret)
  if err != nil {
      log.Fatal("Error getting tokens: ", err)
  }

   gplusID, err := plus.DecodeIDToken(idToken)
   if err != nil {
      log.Fatal("Error decoding ID token: ", err)
   }
*/
package plus
