package controllers

import (
	"manga-sage/initializers"
	"manga-sage/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserCreate(c *gin.Context) {

	var body struct {
		Username string
		Email    string
		Password string
	}

	c.Bind(&body)

	hash, err := HashPassword(body.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user := models.User{Username: body.Username, Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"user": user,
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func UserLogin(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	c.Bind(&body)

	var user models.User
	result := initializers.DB.Where(&models.User{Email: body.Email}).First(&user)

	if result.Error != nil || !CheckPasswordHash(body.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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
