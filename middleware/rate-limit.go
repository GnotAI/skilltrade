package middleware

import (
	"net/http"
	"time"

	redis "github.com/redis/go-redis/v9" // ✅ redis/v9, not redis/v8
	"github.com/ulule/limiter/v3"
	limiterRedis "github.com/ulule/limiter/v3/drivers/store/redis"
	stdlibmiddleware "github.com/ulule/limiter/v3/drivers/middleware/stdlib"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func RateLimiterMiddleware() func(http.Handler) http.Handler {
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  100,
	}

	// ✅ Use the compatible client with limiter
	store, err := limiterRedis.NewStoreWithOptions(redisClient, limiter.StoreOptions{
		Prefix:   "rate_limit",
		MaxRetry: 3,
	})
	if err != nil {
		panic("Failed to create limiter store: " + err.Error())
	}

	limiterInstance := limiter.New(store, rate)
	return stdlibmiddleware.NewMiddleware(limiterInstance).Handler
}
