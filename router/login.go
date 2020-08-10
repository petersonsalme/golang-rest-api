package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petersonsalme/golang-rest-api/middleware"
	"github.com/petersonsalme/golang-rest-api/model"
	"github.com/petersonsalme/golang-rest-api/redis"
)

var user = model.User{
	ID:       1,
	Username: "username",
	Password: "password",
}

// Login should verify user credentials
func Login(c *gin.Context) {
	var u model.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	token, err := middleware.CreateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := redis.CreateAuth(user.ID, token)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}

	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}

	c.JSON(http.StatusOK, tokens)
}
