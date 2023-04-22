package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s", name)
	})

	r.GET("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action"
		c.String(http.StatusOK, "%t", b)
	})

	r.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "Yoo Groups!")
	})

	r.GET("/hello", func(c *gin.Context) {
		fn := c.DefaultQuery("firstname", "World")
		ln := c.Query("lastname")
		c.String(http.StatusOK, "Hello, %s %s", fn, ln)
	})

	r.POST("/form_post", func(c *gin.Context) {
		fn := c.DefaultQuery("firstname", "World")
		ln := c.Query("lastname")
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"first":   fn,
			"last":    ln,
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	r.POST("/map", func(c *gin.Context) {
		ids := c.QueryMap("id")
		names := c.PostFormMap("names")
		c.String(http.StatusOK, "IDs: %v\nNames: %v", ids, names)
	})

	r.Run()
}
