package db

import (
	"fmt"
	"os"

	"github.com/gin-contrib/sessions/redis"
)

var SessionsStore *redis.Store

func CreateSessionStorage() (redis.Store, error) {
	store, err := redis.NewStore(10, "tcp", fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), os.Getenv("REDIS_PASS"), []byte("id"))
	if err != nil {
		return nil, err
	}

	SessionsStore = &store
	return store, nil
}
