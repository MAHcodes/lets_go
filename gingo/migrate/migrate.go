package main

import (
	"github.com/MAHcodes/lets_go/gingo/initializers"
	"github.com/MAHcodes/lets_go/gingo/models"
)

func init() {
  initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
