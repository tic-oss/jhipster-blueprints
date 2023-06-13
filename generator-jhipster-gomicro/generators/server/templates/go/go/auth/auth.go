package auth

import (
	"context"
	"log"
	"fmt"
	"github.com/asim/go-micro/v3/server"
	"github.com/Nerzal/gocloak/v13"
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

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("[wrapper] server request: %v", req.Endpoint())
        headers := req.Header()
		tokenString, ok := headers["Authorization"]
		if !ok || len(tokenString) < 1 {
			log.Printf("Missing authorization token")
			return fmt.Errorf("Missing authorization token")
		}
        log.Printf("token:"+tokenString)
		rptResult, err := client.RetrospectToken(ctx, string(tokenString), clientid,clientsecret, realmname)
		if err != nil {
			log.Printf("Inspection failed: %s", err.Error())
			return fmt.Errorf("Inspection failed: %s", err.Error())
		}
        log.Printf("%v",rptResult)
		istokenvalid :=*rptResult.Active 

		if !istokenvalid {
			log.Printf("Token is not active")
			return fmt.Errorf("Token is not active")

		}
		err = fn(ctx, req, rsp)
		return err
	}
}