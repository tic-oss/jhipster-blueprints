package services

import (
	"encoding/json"
	"net/http"
	// "context"
	"<%= packageName %>/errors"
	"strings"
//  "crypto/tls"
	oidc "github.com/coreos/go-oidc"
 "golang.org/x/oauth2"
	_ "github.com/gorilla/mux"
	"fmt"
)


type LoginResponse struct {
	AccessToken string `json:"access_token"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

var (
	clientId     = "web_app"
	// clientSecret = "cUXhytlJXcm7rV0SwijqDSC3d73uVx0B"
	realm        = "jhipster"
	redirectURL = "http://localhost:<%= serverPort %>/redirect"
	state = "somestate"
	configURL = "http://localhost:9080/realms/jhipster"
	rawIDToken string
)

// var resp 
// var provider oidc.Provider
// var client gocloak.GoCloak

func InitializeOauthServer() {
	// fmt.Print(provider.Endpoint())
	// oauth2Config := oauth2.Config{
	// 	ClientID:     clientId,
	// 	ClientSecret: clientSecret,
	// 	RedirectURL:  redirectURL,
	// 	Endpoint: provider.Endpoint(),
	// 	Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	// }
}

func Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		provider, err := oidc.NewProvider(r.Context(), configURL)
        if err != nil {
          panic(err)
        }
		fmt.Print(provider.Endpoint())
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
	    // token, _ := client.LoginClient(r.Context(), clientId,clientSecret, realm)
        // fmt.Println(token)
		accessToken := strings.Split(authHeader, " ")[1]
		// fmt.Println(accessToken);
		
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
	//   fmt.Print(provider.Endpoint())
	oidcConfig := &oidc.Config{
		ClientID: clientId,
	   }
	verifier := provider.Verifier(oidcConfig)
	oauth2Config := oauth2.Config{
		ClientID:     clientId,
		// ClientSecret: clientSecret,
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
	// fmt.Println(r.URL.Query().Get("code"))
	// atoken :=oauth2Token [0]
	// fmt.Println(atoken)
	// fmt.Println(oauth2Token)
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
	// fmt.Println(rawIDToken)
	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	}{oauth2Token, new(json.RawMessage)}
     
	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// data, err := json.MarshalIndent(resp, "", "    ")
	// if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		// return
	// }
	// fmt.Print(data);
	// tokens :=strings.Split(data, "[")
	// fmt.Print(tokens[0])
	// req,_ :=http.NewRequest("GET","http://127.0.0.1:7090/events",nil)
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization",rawIDToken)
	// res,err :=http.DefaultClient.Do(req)
	// defer res.Body.Close()
	// fmt.Println(rawIDToken)
	// fmt.Println(r.Header)
    http.Redirect(w,r,"/events",http.StatusFound)
}

func Login(w http.ResponseWriter, r *http.Request){
	provider, err := oidc.NewProvider(r.Context(), configURL)
	if err != nil {
		panic(err)
	  }
	//   fmt.Print(provider.Endpoint())
	oauth2Config := oauth2.Config{
		ClientID:     clientId,
		// ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint: provider.Endpoint(),
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}

	rawAccessToken := r.Header.Get("Authorization")
	fmt.Print(rawAccessToken)
        if rawAccessToken == "" {
            http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
            return
        }
		w.Write([]byte("hello world"))
	// rq := &loginRequest{}
	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(rq); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// restyClient := client.RestyClient()
	// restyClient.SetDebug(true)
	// restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// jwt, err := client.Login(r.Context(),
	// 	clientId,
	// 	clientSecret,
	// 	realm,
	// 	rq.Username,
	// 	rq.Password)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusForbidden)
	// 	return
	// }

	// rs := &loginResponse{
	// 	AccessToken:  jwt.AccessToken,
	// 	RefreshToken: jwt.RefreshToken,
	// 	ExpiresIn:    jwt.ExpiresIn,
	// }

	// rsJs, _ := json.Marshal(rs)

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// _, _ = w.Write(rsJs)
}