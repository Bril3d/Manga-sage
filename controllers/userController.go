package controllers

import (
	"manga-sage/initializers"
	"manga-sage/models"

	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {

	var body struct {
		Username string
		Email    string
		Password string
	}

	c.Bind(&body)
	user := models.User{Username: body.Username, Email: body.Email, Password: body.Password}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"user": user,
	})
}

func FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := initializers.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func Update() error {
	var u models.User
	result := initializers.DB.Save(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Delete() error {
	var u models.User
	result := initializers.DB.Delete(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
