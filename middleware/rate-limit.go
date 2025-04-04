package middleware

import (
	"net/http"
	"os"
	"time"

	redis "github.com/redis/go-redis/v9"
	"github.com/ulule/limiter/v3"
	stdlibmiddleware "github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	limiterRedis "github.com/ulule/limiter/v3/drivers/store/redis"
)

func RateLimiterMiddleware() func(http.Handler) http.Handler {
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  100,
	}

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		panic("REDIS_URL environment variable not set")
	}

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic("Failed to parse REDIS_URL: " + err.Error())
	}

	redisClient := redis.NewClient(opt)

	// Create limiter store
	store, err := limiterRedis.NewStoreWithOptions(redisClient, limiter.StoreOptions{
		Prefix:   "rate_limit",
		MaxRetry: 3,
	})
	if err != nil {
		panic("Failed to create limiter store: "+ err.Error())
	}

	limiterInstance := limiter.New(store, rate)
	return stdlibmiddleware.NewMiddleware(limiterInstance).Handler
}
