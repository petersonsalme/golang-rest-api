package redis

import (
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/petersonsalme/golang-rest-api/model"
)

var client *redis.Client

// Connect should connect to Redis
func Connect() *redis.Client {
	dsn := os.Getenv("REDIS_DSN")

	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}

	client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}

	return client
}

// CreateAuth createAuth
func CreateAuth(userid uint64, token *model.Token) error {
	at := time.Unix(token.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(token.RtExpires, 0)
	now := time.Now()

	errAccess := client.Set(token.AccessUUID, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := client.Set(token.RefreshUUID, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}
