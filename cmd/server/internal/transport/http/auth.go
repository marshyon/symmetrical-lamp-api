package http

import (
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

/*

this is fine and all but there is no mechanism to create a token

see the following docs on the same library where we can create a token

it would be typical to create a token in a route handler for auth that returns the token
and accepts a username and password if you are using basic auth

if we are not and are using a token already provided by a third party, then we would
just need to accept the token via the bearer token header and validate it against the third party
using their api

in the case of firebase, they have an api that will validate the token for us, as I used this
in a recent project in python and have used the api in node and go also, so this should
be a goer

failing that, the user / password could still be used to auth against firebase and then
create a token within the app and return that to the client but this would requre again
the creation of an auth token endpoint for the user to go to and provide their credentials
and then return the token

for browser based firebase clients, this is handled by the firebase sdk and the user is
redirected to the firebase auth page and then redirected back to the app with the token
in the url

this is by far more agreeable tom me than handling the user / password in the app with all
the security implications that entails

https://pkg.go.dev/github.com/dgrijalva/jwt-go#example-NewWithClaims-CustomClaimsType

mySigningKey := []byte("AllYourBase")

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

// Create the Claims
claims := MyCustomClaims{
	"bar",
	jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	},
}

token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
ss, err := token.SignedString(mySigningKey)
fmt.Printf("%v %v", ss, err)


*/

func JWTAuth(
	original func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		// Bearer token-string
		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		if validateToken(authHeaderParts[1]) {
			original(w, r)
		} else {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
	}
}

func validateToken(accessToken string) bool {
	var mySigningKey = []byte("missionimpossible")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not validate auth token")
		}

		return mySigningKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}
