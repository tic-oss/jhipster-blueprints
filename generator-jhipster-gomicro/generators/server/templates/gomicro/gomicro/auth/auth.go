package auth

import (
	"context"
	"github.com/micro/micro/v3/service/logger"
	"net/http"
	"github.com/Nerzal/gocloak/v13"
	"encoding/json"
	"os"
)

var (
	client *gocloak.GoCloak
	clientid string
	clientsecret string
	realmname string
	keycloakurl string
)

func SetClient() {
	clientid =os.Getenv("CLIENT_ID")
	clientsecret = os.Getenv("CLIENT_SECRET")
	realmname =os.Getenv("REALM_NAME")
	keycloakurl=os.Getenv("KEYCLOAK_URL")
	client = gocloak.NewClient(keycloakurl)
}


func Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) < 1 {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode("Unauthorized")
			return
		}
		accessToken := authHeader
		rptResult, err := client.RetrospectToken(context.TODO(),string(accessToken), clientid,clientsecret, realmname)
		if err != nil {
			logger.Errorf("Inspection failed: %s", err.Error())
			return 
		}
		istokenvalid :=*rptResult.Active 
		if !istokenvalid {
			logger.Errorf("Token is not active")
			return 
		}
		next.ServeHTTP(w,r)
	})
}
