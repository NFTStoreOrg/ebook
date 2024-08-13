package models

import "github.com/redis/go-redis/v9"

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"127.0.0.1:26379", "127.0.0.1:26380"},
	})
}
