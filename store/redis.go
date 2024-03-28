package store

import (
	"fmt"
	"market_back/conf"
	"time"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func GetRDB() *redis.Client {
	return rdb
}

// InitRedis 初始化redis
func InitRedis(c *conf.RedisConf) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", c.RedisHost, c.RedisPort),
		Password:     c.RedisPassword,
		DB:           c.RedisDB,
		PoolSize:     1000,
		MinIdleConns: 100,
	})
}

func RedisClose() {
	rdb.Close()
}

// ExistsKey 判断key是否存在
func ExistsKey(key string) bool {
	return rdb.Exists(key).Val() > 0
}

// GetKey 获取key对应的value值
func GetKey(key string) (string, error) {
	return rdb.Get(key).Result()
}

// SetKey 设置key对应的value值
func SetKey(key string, value string) error {
	return rdb.Set(key, value, 0).Err()
}

// 设置key对应的value值，并设置过期时间
func SetKeyWithExpire(key string, value string, expire time.Duration) error {
	return rdb.Set(key, value, expire).Err()
}
