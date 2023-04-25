package main

import (
	"github.com/MAHcodes/lets_go/gingo/controllers"
	"github.com/MAHcodes/lets_go/gingo/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.POST("/user", controllers.CreateUser)
	r.GET("/user/:id", controllers.GetUser)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	r.Run()
}
