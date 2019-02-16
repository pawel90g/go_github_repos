package controllers

import (
	"github-user-statistics/models"
	"github-user-statistics/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var authResponse models.AuthResponse

func GetClientID() string {
	return services.GetClientID()
}

func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{"clientId": services.GetClientID(), "host": "http://localhost:8000"})
}

func GithubCallback(c *gin.Context) {
	code := c.Query("code")

	authResponse = services.Auth(code)
	me := services.GetAuthUser(authResponse)
	email := services.GetUserEmail(authResponse)

	me.Email = email.Email

	c.HTML(http.StatusOK, "user.tmpl.html", gin.H{"user": me})
}

func GithubMyAccount(c *gin.Context) {
	email := services.GetUserEmail(authResponse)
	me := services.GetAuthUser(authResponse)

	me.Email = email.Email

	c.HTML(http.StatusOK, "user.tmpl.html", gin.H{"user": me})
}

func GithubUsers(c *gin.Context) {
	users := services.GetUsers()

	c.HTML(http.StatusOK, "users.tmpl.html", gin.H{"users": users})
}

func GithubUser(c *gin.Context) {
	login := c.Param("login")

	user := services.GetUser(login)
	repos := services.GetRepos(login)

	c.HTML(http.StatusOK, "user.tmpl.html", gin.H{"user": user, "repos": repos})
}

func GithubRepo(c *gin.Context) {
	login := c.Param("login")
	repo := c.Param("repo")

	repoObj := services.GetRepo(login, repo)
	langStats := services.GetRepoLanguages(login, repo)

	c.HTML(http.StatusOK, "repo.tmpl.html", gin.H{"userLogin": login, "repo": repoObj, "langStats": langStats})
}
