package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/heroku/github-user-statistics/routers"
)

func main() {
	port := "8000" //os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	routers.InitGithubRouter(router)

	router.Run(":" + port)
}
