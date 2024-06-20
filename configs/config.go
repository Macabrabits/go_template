package configs

import (
	"os"

	"github.com/coreos/go-oidc"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/oauth2"
)

type KeycloakConfig struct {
	ClientId    string
	Secret      string
	InternalUrl string
	PublicUrl   string
	Realm       string
}

type Config struct {
	Port         string
	MysqlCFG     mysql.Config
	Oauth2Config oauth2.Config
	Keycloak     KeycloakConfig
}

func config(defaultValue string, optionalValue string) string {
	if optionalValue == "" {
		return defaultValue
	}
	return optionalValue
}

func GetConfig() Config {
	keyCloakConfig := KeycloakConfig{
		ClientId:    config("appclient", os.Getenv("KEYCLOAK_CLIENT_ID")),
		Secret:      config("bJg5JMPTukmsF9F51TTLvll4u7dnX4vv", os.Getenv("KEYCLOAK_CLIENT_SECRET")),
		InternalUrl: config("http://keycloak:8080/realms/app", os.Getenv("KEYCLOAK_ISSUER_URL")),
		PublicUrl:   config("http://localhost:8082/realms/app", os.Getenv("KEYCLOAK_PUBLIC_URL")),
		Realm:       config("app", os.Getenv("KEYCLOAK_REALM")),
	}

	cfg := Config{
		Port: config("8080", os.Getenv("PORT")),
		MysqlCFG: mysql.Config{
			User:      config("root", os.Getenv("DBUSER")),
			Passwd:    config("root", os.Getenv("DBPASS")),
			Net:       "tcp",
			Addr:      config("localhost", os.Getenv("DBHOST")) + ":3306",
			DBName:    "app",
			ParseTime: true,
		},
		Keycloak: keyCloakConfig,
		Oauth2Config: oauth2.Config{
			ClientID:     config("appclient", os.Getenv("KEYCLOAK_CLIENT_ID")),
			ClientSecret: config("", os.Getenv("KEYCLOAK_CLIENT_SECRET")),
			RedirectURL:  config("http://localhost:8080/api/v1/auth2/callback", os.Getenv("KEYCLOAK_REDIRECT_URL")),
			Endpoint: oauth2.Endpoint{
				AuthURL:  config("http://localhost:8082/realms/app", os.Getenv("KEYCLOAK_PUBLIC_URL")) + "/protocol/openid-connect/auth",
				TokenURL: config("http://keycloak:8080/realms/app", os.Getenv("KEYCLOAK_ISSUER_URL")) + "/protocol/openid-connect/token",
			},
			Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
		},
	}
	return cfg
}
