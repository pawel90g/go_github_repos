package main

import (
	"log"
	"net/http"

	"github-user-statistics/controllers"
	"github-user-statistics/routers"

	"github.com/gin-gonic/gin"
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

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{"clientId": controllers.GetClientID(), "host": "http://localhost:8000"})
	})

	routers.InitGithubRouter(router)

	router.Run(":" + port)
}
