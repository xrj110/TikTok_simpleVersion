package tools

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func RedisInit() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 如果没有密码则为空
		DB:       0,  // 默认 DB
	})
	ctx := context.Background()

	err := Client.Ping(ctx).Err()
	if err != nil {
		fmt.Println("redis con failed:", err)
	}
}
func DocRedisInit(name string, port string) {
	Client = redis.NewClient(&redis.Options{
		Addr:     name + ":" + port,
		Password: "", // 如果没有密码则为空
		DB:       0,  // 默认 DB
	})
	ctx := context.Background()

	err := Client.Ping(ctx).Err()
	if err != nil {
		fmt.Println("redis con failed:", err)
	}

}

func GetClient() *redis.Client {
	return Client
}
