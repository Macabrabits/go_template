package services

import (
	"context"
	"strings"

	"encoding/json"
	"fmt"
	"io"

	// "strings"
	"time"

	// "fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	// "github.com/hashicorp/vault/sdk/helper/password"
	// "golang.org/x/oauth2"
	// "go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/metric"
	// oidc "github.com/coreos/go-oidc"
)

type AuthService struct {
}

type AuthParams struct {
	Username string `json:"username" example:"appuser" validate:"required"`
	Password string `json:"password" example:"test"    validate:"required"`
}
type AuthResponse struct {
	Access_token       string `json:"access_token"`
	Expires_in         int    `json:"expires_in" example:"300"`
	Refresh_expires_in int    `json:"refresh_expires_in" example:"1800"`
	Refresh_token      string `json:"refresh_token"`
}

var (
	// tracer = otel.Tracer(name)
	// meter = otel.Meter(name)
	// logger =
	clientID     = "appclient"
	clientSecret = "bJg5JMPTukmsF9F51TTLvll4u7dnX4vv"
	rsaPublicKey = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAut7lCwO95fB/C6Jj59cl
gZUe5ZYgQVQ30XGdij1nTmnGZIYqHaJNKOJvRD2yoM1X7DUPSlhlOcdefTtv9r5Z
8zuBiZWTkE6s4pbJroeY53HnSk1oXnK7dto4MXgWbm1+G8VvANBAIew6FeRFDp/B
22ZdyefQxBFXqGPO5hXf7rGrYZEXSAx6AY1qnBwBAwiCVwC7DmsTmMk3ZPE642+Z
tWQUTfp8A27BF8UlmNMVTB65fB7Ze7FS5sktLAhNunqv9mlMS7pOlvsgdOHvnBPQ
ngz6hbXa2xDw5E3aoK5gRbfIOXsIDWPtEq3xSc3GSBK3/xLbQ9JVED0hYwwNzFTI
FwIDAQAB
-----END PUBLIC KEY-----
`
)

func NewAuthsService(
// repository *repository.AuthRepository,
) AuthService {
	return AuthService{
		// repository,
	}
}

func (svc *AuthService) GetToken(ctx context.Context, authParams AuthParams) (any, error) {
	// provider, err := oidc.NewProvider(ctx, "http://keycloak:8080/realms/app")
	// if err != nil {
	// 	panic(err)
	// }
	// config := oauth2.Config{
	// 	ClientID:     clientID,
	// 	ClientSecret: clientSecret,
	// 	Endpoint:     provider.Endpoint(),
	// 	RedirectURL:  "http://app:8080/api/v1/auth/callback",
	// 	Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	// }
	// state := "maginca"
	// fmt.Println(config.AuthCodeURL(state))
	fmt.Println(authParams)
	response, err := http.PostForm("http://keycloak:8080/realms/app/protocol/openid-connect/token", url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"username":      {authParams.Username},
		"password":      {authParams.Password},
		"grant_type":    {"password"},
	})
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error getting de token: %w", string(body))
	}
	res := AuthResponse{}
	json.Unmarshal(body, &res)
	return res, nil
}

func (svc *AuthService) AuthCallback(ctx context.Context) (gin.H, error) {

	return gin.H{
		"message": "success",
		"data":    "data",
	}, nil
}

func (svc *AuthService) Auth(ctx context.Context, jwtToken string) (any, error) {
	jwtToken = strings.TrimPrefix(jwtToken, "Bearer ")

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(rsaPublicKey))
	if err != nil {
		return nil, fmt.Errorf("error parsing public key: %v", err)
	}

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	fmt.Println(token.Claims.GetExpirationTime())
	fmt.Println(time.Now())

	return gin.H{
		"message": "success",
		"data":    "token validated",
	}, nil

}
