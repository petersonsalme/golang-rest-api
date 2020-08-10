package main

import (
	"log"
	"os"

	"github.com/petersonsalme/golang-rest-api/redis"
	"github.com/petersonsalme/golang-rest-api/router"

	"github.com/gin-gonic/gin"
)

var routerEngine *gin.Engine

func init() {
	os.Setenv("ACCESS_SECRET", "123456789ABCDEF")
	os.Setenv("REFRESH_SECRET", "987654321FEDCBA")

	redis.Connect()
	routerEngine = gin.Default()
}

func main() {
	routerEngine.POST("/login", router.Login)
	log.Fatal(routerEngine.Run(":8080"))
}
