package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "fmt"
)

func setupRouter() *gin.Engine {
    router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.Static("/js", "./templates/js")
	router.Static("/css", "./templates/css")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sign_up.html", gin.H{
			"title": "Sign Up",
		})
	})

	router.GET("/signin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sign_in.html", gin.H{
			"title": "Sign Ip",
		})
	})

	router.POST("/sign_up", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")
		address := c.PostForm("address")
		address2 := c.PostForm("address2")
		firstname := c.PostForm("firstname")
		lastname := c.PostForm("lastname")
		nickname := c.PostForm("nickname")
		city := c.PostForm("city")
		state := c.PostForm("state")
		zip := c.PostForm("zip")
        c.JSON(200, gin.H{
            "status": "posted",
            "email": email,
            "password": password,
			"address": address,
			"address2": address2,
			"city": city,
			"state": state,
			"zip": zip,
        })
	})

    return router
}



func main() {
    router := setupRouter()
    router.Run(":8080") // default localhost:8000
}