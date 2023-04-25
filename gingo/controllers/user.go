package controllers

import (
	"net/http"

	"github.com/MAHcodes/lets_go/gingo/initializers"
	"github.com/MAHcodes/lets_go/gingo/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	res := initializers.DB.First(&user, id)

	if res.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, user)

}

func CreateUser(c *gin.Context) {
	var body struct {
		Name, Bio string
		Age       int8
	}

	c.Bind(&body)

	user := models.User{
		Name: body.Name,
		Bio:  body.Bio,
		Age:  uint8(body.Age),
	}

	res := initializers.DB.Create(&user)

	if res.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	res := initializers.DB.Delete(&user, id)

	if res.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name, Bio string
		Age       int8
	}

	c.Bind(&body)

	var user models.User
	initializers.DB.First(&user, id)

	initializers.DB.Model(&user).Updates(models.User{
		Name: body.Name,
		Bio:  body.Bio,
		Age:  uint8(body.Age),
	})

	c.JSON(http.StatusOK, user)
}
