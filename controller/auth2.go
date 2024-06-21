package controller

import (
	"encoding/json"
	"fmt"
	"io"

	// "io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"github.com/macabrabits/go_template/configs"

	"golang.org/x/oauth2"
)

var (
	keycloak         configs.KeycloakConfig = configs.GetConfig().Keycloak
	oauth2Config     oauth2.Config          = configs.GetConfig().Oauth2Config
	keycloakClientID                        = oauth2Config.ClientID
)

type Auth2Controller struct {
	provider *oidc.Provider
}

var ()

func NewAuth2Controller(
	provider *oidc.Provider,

) Auth2Controller {
	return Auth2Controller{provider}
}

func (s *Auth2Controller) IndexHandler(c *gin.Context) {
	spanName := c.Request.Method + " - " + c.Request.URL.Path
	_, span := tracer.Start(c, spanName)
	defer span.End()
	c.String(http.StatusOK, "Welcome to the Home Page!")
}

// @Summary		Initiates login process
// @Description	Redirects to Keycloak login page
// @Tags			OAuth2
// @Produce		html
// @Success		302	{string}	string	"Redirects to Keycloak login page"
// @Router			/auth2/login [get]
func (s *Auth2Controller) LoginHandler() gin.HandlerFunc {

	// // Redirect to Keycloak login page
	// url := oauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline)
	// fmt.Println(url)
	// url = strings.Replace(url, "keycloak:8080", "localhost:8082", -1)
	// // url = strings.Replace(url, "localhost%3A8080", "localhost%3A8082", -1)
	// c.Redirect(http.StatusFound, url)
	return func(c *gin.Context) {
		// Redirect to Keycloak login page
		url := oauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline)
		c.Redirect(http.StatusFound, url)
	}

}

// @Summary		Handles authentication callback
// @Description	Handles the OAuth2 callback from Keycloak
// @Tags			OAuth2
// @Produce		plain
// @Success		200	{string}	string	"Authentication successful!"
// @Router			/auth2/callback [get]
func (s *Auth2Controller) CallbackHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")

		// Exchange the authorization code for an access token
		token, err := oauth2Config.Exchange(c, code)
		if err != nil {
			log.Printf("Failed to exchange token: %v", err)
			c.String(http.StatusInternalServerError, "Failed to exchange token")
			return
		}

		// Verify ID Token
		rawIDToken, ok := token.Extra("id_token").(string)
		if !ok {
			log.Println("No id_token field in oauth2 token")
			c.String(http.StatusInternalServerError, "No id_token field in oauth2 token")
			return
		}

		verifier := s.provider.Verifier(&oidc.Config{ClientID: oauth2Config.ClientID, SkipIssuerCheck: true})
		idToken, err := verifier.Verify(c, rawIDToken)
		if err != nil {
			log.Printf("Failed to verify ID token: %v", err)
			c.String(http.StatusInternalServerError, "Failed to verify ID token")
			return
		}

		// At this point, user is authenticated successfully
		// You can get user info from ID token if needed

		// Extract claims
		var claims map[string]interface{}
		if err := idToken.Claims(&claims); err != nil {
			log.Printf("Failed to extract claims: %v", err)
			c.String(http.StatusInternalServerError, "Failed to extract claims")
			return
		}

		c.Redirect(http.StatusFound, "http://localhost:8082/swagger/index.html")

		// c.JSON(http.StatusOK, gin.H{
		// 	"OAuth2Token": token,
		// 	"rawIDToken":  rawIDToken,
		// 	"claims":      claims,
		// 	"idToken":     idToken,
		// 	"code":        code,
		// })
	}
}

// @Summary		Logs out the user
// @Description	Logs out the user from Keycloak
// @Tags			OAuth2
// @Produce		html
// @Success		302	{string}	string	"Redirects to Keycloak logout"
// @Router			/auth2/logout [get]
func (s *Auth2Controller) LogoutHandler(c *gin.Context) {
	// Redirect to Keycloak logout endpoint
	oidcLogoutPath := "/protocol/openid-connect/logout"
	logoutURL := keycloak.PublicUrl + oidcLogoutPath
	c.Redirect(http.StatusTemporaryRedirect, logoutURL)
}
func (s *Auth2Controller) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tctx, initialSpan := tracer.Start(c, "AuthMiddleware")
		defer initialSpan.End()
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		// if token == authHeader {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
		// 	return
		// }
		verifier := s.provider.Verifier(&oidc.Config{ClientID: keycloakClientID, SkipIssuerCheck: true, SkipClientIDCheck: true})
		_, span := tracer.Start(tctx, "verify")
		idToken, err := verifier.Verify(c, token)
		span.End()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			return
		}

		fmt.Println("I'm heere!")

		var claims map[string]interface{}
		if err := idToken.Claims(&claims); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to extract claims"})
			return
		}

		// Attach the claims to the context
		c.Set("claims", claims)
		c.Next()

	}
}

// type TokenIntrospectionResponse struct {
// 	active bool   `json:"active"`
// 	sid    string `json:"sid"`
// 	// Add other fields as needed
// }

type TokenIntrospectionResponse struct {
	Active bool   `json:"active"`
	Sid    string `json:"sid"`
}

func (s *Auth2Controller) introspectToken(token string) (*TokenIntrospectionResponse, error) {
	introspectionEndpoint := "http://mykeycloak:8080/realms/app/protocol/openid-connect/token/introspect"
	resp, err := http.PostForm(introspectionEndpoint, url.Values{
		"client_id":     {oauth2Config.ClientID},
		"client_secret": {oauth2Config.ClientSecret},
		"token":         {token},
	})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	res := TokenIntrospectionResponse{}
	json.Unmarshal(body, &res)
	return &res, nil
}

func (s *Auth2Controller) RefreshTokenIfNeeded() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		refreshToken := c.GetHeader("refresh_token")
		if accessToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Access token not found or expired"})
			return
		}

		if refreshToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
			return
		}

		introspectionResponse, err := s.introspectToken(accessToken)
		if err != nil {
			log.Printf("Failed to introspect token: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to introspect token"})
			return
		}

		if !introspectionResponse.Active {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User is logged out"})
			return
		}

		token, err := oauth2Config.TokenSource(c, &oauth2.Token{RefreshToken: refreshToken}).Token()
		if err != nil {
			log.Printf("Failed to refresh token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to refresh token"})
			return
		}

		accessToken = token.AccessToken
		refreshToken = token.RefreshToken

		c.Next()
	}
}
