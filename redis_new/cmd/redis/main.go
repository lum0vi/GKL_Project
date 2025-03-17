package main

import (
	"context"
	"fmt"
	"redis_new/internal/config"
	"redis_new/internal/connect"
	"redis_new/internal/logger"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	ctx, err := logger.New(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	// cfg := config.Config{
	// 	Addr: "localhost:6379",
	// 	Password: ,
	// }
	db, err := connect.NewClientRedis(ctx, config.Config{})
	if err != nil {
		panic(err)
	}

	// Запись данных

	// db.Set(контекст, ключ, значение, время жизни в базе данных)
	if err := db.Set(ctx, "key", "test value", 0).Err(); err != nil {
		fmt.Printf("failed to set data, error: %s", err.Error())
	}

	if err := db.Set(ctx, "key2", 333, 30*time.Second).Err(); err != nil {
		fmt.Printf("failed to set data, error: %s", err.Error())
	}

	// Получение данных

	val, err := db.Get(ctx, "key").Result()
	if err == redis.Nil {
		fmt.Println("value not found")
	} else if err != nil {
		fmt.Printf("failed to get value, error: %v\n", err)
	}

	val2, err := db.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("value not found")
	} else if err != nil {
		fmt.Printf("failed to get value, error: %v\n", err)
	}

	fmt.Printf("value: %v\n", val)
	fmt.Printf("value: %v\n", val2)
}
