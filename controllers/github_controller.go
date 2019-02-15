package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthResponse struct {
	AccessToken string
	Scope       string
	TokenType   string
}

var clientID = "08a723e17a18bbf7226f"
var clientSecret = "a97026b2da5a2dcfa9db1d5b317bb84f03349a88"

func GetClientID() string {
	return clientID
}

func GithubCallback(c *gin.Context) {
	code := c.Query("code")

	response, _ := http.Post("https://github.com/login/oauth/access_token?code="+code+"&client_id="+clientID+"&client_secret="+clientSecret, "application/json", nil)

	var accessToken AuthResponse

	body, _ := ioutil.ReadAll(response.Body)
	bodyStr := string(body)

	accessTokenParts := strings.Split(bodyStr, "&")

	for _, part := range accessTokenParts {
		if strings.HasPrefix(part, "access_token") {
			accessToken.AccessToken = strings.Split(part, "=")[1]
		} else if strings.HasPrefix(part, "scope") {
			accessToken.Scope = strings.Split(part, "=")[1]
		} else if strings.HasPrefix(part, "token_type") {
			accessToken.TokenType = strings.Split(part, "=")[1]
		}
	}

	log.Println(accessToken)
}
