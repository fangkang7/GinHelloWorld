package tool

import "github.com/go-redis/redis"

type RedisStore struct {
	*redis.Client
}

// 初始化redis
func InitRedisStore() *RedisStore {
	config := GetConfig().RedisConfig

	Redis := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.Db,
	})

	redisStore := &RedisStore{Redis}

	return redisStore
}
