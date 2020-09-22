package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sesi12/controller"
)


func main() {
//Untuk Alamat Urlnya
	router := gin.Default()
	router.LoadHTMLGlob("html/*")
	router.Static("images", "images")
	v1 := router.Group("/contact/service")
	{
		v1.POST("", controller.CretedFeedback)
	}
	router.GET("/", func(c *gin.Context) {
		c.HTML(
		  http.StatusOK,
		  "index.html",
		  gin.H{
			"title": "Home Page",
		  },
		)
	
	  })
	  router.GET("/view", func(c *gin.Context) {
		c.HTML(
		  http.StatusOK,
		  "view.html",
		  gin.H{
			"title": "Contact Us",
		  },
		)
	
	  })

	  router.GET("/login", func(c *gin.Context) {
		c.HTML(
		  http.StatusOK,
		  "login.html",
		  gin.H{
			"title": "Login",
		  },
		)
	
	  })

	router.Run(":9000")
}
