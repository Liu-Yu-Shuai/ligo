package redis

import (
	"github.com/go-redis/redis"
	"github.com/lfuture/easygin/pkg/app"
)

var (
	RedisClient *redis.Client
)

func Bootstrap()  {

	redisConfig, err := app.GetConfig().GetSection("redis")

	if err != nil {
		panic("get redis config fail")
	}

	db,err := redisConfig.Key("db").Int()

	maxRetry,err1 := redisConfig.Key("maxRetry").Int()
	poolSize,err2 := redisConfig.Key("poolSize").Int()

	if err != nil || err1 != nil || err2 != nil {
		panic(err)
		panic(err1)
		panic(err2)
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr:    redisConfig.Key("host").String() + ":" + redisConfig.Key("port").String(),
		Password: redisConfig.Key("password").String(),
		DB:       db,
		PoolSize: poolSize,
		MaxRetries: maxRetry,
	})
}