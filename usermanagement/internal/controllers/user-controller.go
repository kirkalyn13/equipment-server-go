package controllers

import (
	"kirky/equipment-server/usermanagement/internal/config"
	"kirky/equipment-server/usermanagement/internal/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	secret   = os.Getenv("SECRET")
	userBody struct {
		Username string
		Password string
		Role     string
	}
	credentials struct {
		Username string
		Password string
	}
)

// Signup
func Signup(c *gin.Context) {
	if c.Bind(&userBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read userBody.",
		})

		return
	}

	// Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(userBody.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password.",
		})

		return
	}

	user := models.User{Username: userBody.Username, Password: string(hash), Role: userBody.Role}
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user.",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

// Login
func Login(c *gin.Context) {
	if c.Bind(&credentials) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read user.",
		})

		return
	}

	var user models.User
	config.DB.First(&user, "username = ?", credentials.Username)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid credentials.",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid credentials.",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token.",
		})

		return
	}

	//Set cookies
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*2, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})

}

// Edit User
func EditUser(c *gin.Context) {
	id := c.Param("id")

	c.Bind(&userBody)
	var toEdit models.User
	config.DB.First(&toEdit, id)

	// Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(userBody.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password.",
		})

		return
	}

	config.DB.Model(&toEdit).Updates(models.User{
		Username: userBody.Username,
		Password: string(hash),
		Role:     userBody.Role,
	})

	c.JSON(http.StatusOK, gin.H{})
}

// Validate User Login
func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
