package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type GoogleAPIClaim struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}

func main() {
	now := time.Now()
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}
	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), GoogleAPIClaim{
		"https://www.googleapis.com/auth/devstorage.readonly",
		jwt.StandardClaims{
			Issuer:    "761326798069-r5mljlln1rd4lrbhg75efgigp36m78j5@developer.gserviceaccount.com",
			Audience:  "https://www.googleapis.com/oauth2/v4/token",
			ExpiresAt: now.Add(time.Hour).Unix(),
			IssuedAt:  now.Unix(),
		},
	})
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(tokenString)

}

var signBytes = []byte(`-----BEGIN PRIVATE KEY-----
...
-----END PRIVATE KEY-----`)
