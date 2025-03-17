package connect

import (
	"context"
	"redis_new/internal/config"
	"redis_new/internal/logger"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func NewClientRedis(ctx context.Context, cfg config.Config) (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "my-password",
		DB:           0,
		MaxRetries:   5,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})
	if err := db.Ping(ctx).Err(); err != nil {
		logger.GetLoggerFromCtx(ctx).Info(ctx, "error to connect redis", zap.Error(err))
		return nil, err
	}
	return db, nil
}
