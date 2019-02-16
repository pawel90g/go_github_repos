package routers

import (
	"github-user-statistics/controllers"

	"github.com/gin-gonic/gin"
)

func InitGithubRouter(router *gin.Engine) {

	router.GET("/callback", controllers.GithubCallback)
	router.GET("/me", controllers.GithubMyAccount)
	router.GET("/users", controllers.GithubUsers)
	router.GET("/users/:login", controllers.GithubUser)
	router.GET("/repo/:login/:repo", controllers.GithubRepo)
}
