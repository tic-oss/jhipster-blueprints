package services

import (
	"encoding/json"
	"net/http"
	"com.cmi.tic/errors"
	"strings"
	oidc "github.com/coreos/go-oidc"
 "golang.org/x/oauth2"
	_ "github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
  }

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

var (
	clientId     = "web_app"
	realm        = "jhipster"
	redirectURL = "http://localhost:9001/redirect"
	state = "somestate"
	configURL = goDotEnvVariable("KEYCLOAK_URL")
	rawIDToken string
)

func Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		provider, err := oidc.NewProvider(r.Context(), configURL)
        if err != nil {
          panic(err)
        }
	    oidcConfig := &oidc.Config{
         ClientID: clientId,
        }
        verifier := provider.Verifier(oidcConfig)
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) < 1 {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(errors.UnauthorizedError())
			return
		}
		accessToken := strings.Split(authHeader, " ")[1]

		_,err = verifier.Verify(r.Context(), accessToken)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errors.BadRequestError(err.Error()))
			return
		}

		next.ServeHTTP(w, r)
	})
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Result struct {
	access_token  string `json:"access_token"`
	token_type   string `json:"token_type"`
	refresh_token string `json:"refresh_token"`
	expiry    string    `json:"expiry"`
}

func Redirect(w http.ResponseWriter, r *http.Request){
	provider, err := oidc.NewProvider(r.Context(), configURL)
	if err != nil {
		panic(err)
	  }
	oidcConfig := &oidc.Config{
		ClientID: clientId,
	   }
	verifier := provider.Verifier(oidcConfig)
	oauth2Config := oauth2.Config{
		ClientID:     clientId,
		RedirectURL:  redirectURL,
		Endpoint: provider.Endpoint(),
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}

	if r.URL.Query().Get("state") != state {
		http.Error(w, "state did not match", http.StatusBadRequest)
		return
	}

	oauth2Token, err := oauth2Config.Exchange(r.Context(), r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}
	idToken, err := verifier.Verify(r.Context(), rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	}{oauth2Token, new(json.RawMessage)}
     
	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
    http.Redirect(w,r,"/events",http.StatusFound)
}

func Login(w http.ResponseWriter, r *http.Request){
	provider, err := oidc.NewProvider(r.Context(), configURL)
	if err != nil {
		panic(err)
	  }
	oauth2Config := oauth2.Config{
		ClientID:     clientId,
		RedirectURL:  redirectURL,
		Endpoint: provider.Endpoint(),
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}

	rawAccessToken := r.Header.Get("Authorization")
        if rawAccessToken == "" {
            http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
            return
        }
		w.Write([]byte("hello world"))
}
