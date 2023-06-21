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
	clientId string
	clientSecret string
	realmName string
	keycloakUrl string
)

func SetClient() {
	clientId =os.Getenv("GO_MICRO_CLIENT_ID")
	clientSecret = os.Getenv("GO_MICRO_CLIENT_SECRET")
	realmName =os.Getenv("GO_MICRO_REALM_NAME")
	keycloakUrl=os.Getenv("GO_MICRO_KEYCLOAK_URL")
	client = gocloak.NewClient(keycloakUrl)
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
		rptResult, err := client.RetrospectToken(context.TODO(),string(accessToken), clientId,clientSecret, realmName)
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
