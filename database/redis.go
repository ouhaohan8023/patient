package database

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func ConnectRedis() *redis.Client {
	environment := "dev"
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString(environment + `.redis.host`),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}