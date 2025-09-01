package main

import (
	redis_utils "Go-Backend/redis/utils"
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	ctx := context.Background()

	rp1D := &redis_utils.RedisPayload{Id: "id-1", Description: "Welcome"}
	rp1B, _ := json.Marshal(rp1D)

	err := client.Set(ctx, "redis/foo", rp1B, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "redis/foo").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("redis/foo", val)
}
