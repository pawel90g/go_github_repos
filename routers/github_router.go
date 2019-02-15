package routers

import (
	"net/http"

	"github.com/heroku/github-user-statistics/controllers"

	"github.com/gin-gonic/gin"
)

func InitGithubRouter(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{"clientId": controllers.GetClientID(), "host": "http://localhost:8000"})
	})

	router.GET("/callback", controllers.GithubCallback)
}
