package services

import (
	"encoding/json"
	"github-user-statistics/models"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	clientID     = "08a723e17a18bbf7226f"
	clientSecret = "a97026b2da5a2dcfa9db1d5b317bb84f03349a88"
)

func Auth(code string) models.AuthResponse {

	response, _ := http.Post("https://github.com/login/oauth/access_token?code="+code+"&client_id="+clientID+"&client_secret="+clientSecret, "application/json", nil)

	var accessToken models.AuthResponse

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

	return accessToken
}

func GetAuthUser(accessToken models.AuthResponse) models.User {
	userReq, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	userReq.Header.Add("Authorization", accessToken.TokenType+" "+accessToken.AccessToken)

	client := http.Client{}
	userResp, _ := client.Do(userReq)
	userBytes, _ := ioutil.ReadAll(userResp.Body)

	var user *models.User

	json.Unmarshal(userBytes, &user)

	return *user
}

func GetUsers() []models.User {
	req, _ := http.NewRequest("GET", "https://api.github.com/users", nil)

	client := http.Client{}
	resp, _ := client.Do(req)
	usersBytes, _ := ioutil.ReadAll(resp.Body)

	var users *[]models.User

	json.Unmarshal(usersBytes, &users)

	return *users
}

func GetUser(login string) models.User {
	req, _ := http.NewRequest("GET", "https://api.github.com/users/"+login, nil)

	client := http.Client{}
	resp, _ := client.Do(req)
	usersBytes, _ := ioutil.ReadAll(resp.Body)

	var user *models.User

	json.Unmarshal(usersBytes, &user)

	return *user
}

func GetUserEmail(accessToken models.AuthResponse) models.Email {
	req, _ := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	req.Header.Add("Authorization", accessToken.TokenType+" "+accessToken.AccessToken)
	client := http.Client{}
	resp, _ := client.Do(req)
	respBytes, _ := ioutil.ReadAll(resp.Body)

	var emails *[]models.Email

	json.Unmarshal(respBytes, &emails)

	for _, email := range *emails {
		if email.Primary {
			return email
		}
	}
	return (*emails)[0]
}

func GetAuthRepos(accessToken models.AuthResponse) []models.Repo {
	req, _ := http.NewRequest("GET", "https://api.github.com/user/repos", nil)
	req.Header.Add("Authorization", accessToken.TokenType+" "+accessToken.AccessToken)

	client := http.Client{}
	resp, _ := client.Do(req)
	reposBytes, _ := ioutil.ReadAll(resp.Body)

	var repos *[]models.Repo

	json.Unmarshal(reposBytes, &repos)

	return *repos
}

func GetRepo(login string, repo string) models.Repo {

	req, _ := http.NewRequest("GET", "https://api.github.com/repos/"+login+"/"+repo, nil)
	client := http.Client{}
	resp, _ := client.Do(req)
	repoBytes, _ := ioutil.ReadAll(resp.Body)

	var repoObj *models.Repo

	json.Unmarshal(repoBytes, &repoObj)

	return *repoObj
}

func GetRepos(login string) []models.Repo {
	req, _ := http.NewRequest("GET", "https://api.github.com/users/"+login+"/repos", nil)

	client := http.Client{}
	resp, _ := client.Do(req)
	reposBytes, _ := ioutil.ReadAll(resp.Body)

	var repos *[]models.Repo

	json.Unmarshal(reposBytes, &repos)

	return *repos
}

func GetRepoLanguages(user string, repo string) map[string]int32 {
	req, _ := http.NewRequest("GET", "https://api.github.com/repos/"+user+"/"+repo+"/languages", nil)

	client := http.Client{}
	resp, _ := client.Do(req)
	langsBytes, _ := ioutil.ReadAll(resp.Body)

	var repoLangs *map[string]int32

	json.Unmarshal(langsBytes, &repoLangs)

	return *repoLangs
}

func GetClientID() string {
	return clientID
}
