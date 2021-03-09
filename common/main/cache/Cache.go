package cache

import (
	"time"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"

	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
	
)

// RDB for redis client
var rdb *redis.Client
// CTX for context
var ctx = context.Background()

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	log.Error(err)
	if err != nil {
		panic("failed to connect redis")
	}
}

// StoreBy put value to cache
func StoreBy(namepace string, key string, value interface {}) error {
	serialized, err := json.Marshal(value)
	log.Info(serialized)
	if err != nil {
		log.Error(err)
		return err
	}
	_, err = rdb.Set(ctx, namepace + ":" + key, serialized, 0).Result()
	return err
}

// GetBy get value from cache
func GetBy(namepace string, key string, v interface {}) error {
	serialized, err := rdb.Get(ctx, namepace + ":" + key).Result()
	if err != nil {
		log.Error(err)
		return err
	}
	err = json.Unmarshal([]byte(serialized), &v)
	return err
} 

// DeleteBy value form cache
func DeleteBy(namepace string, key string) error {
	_, err := rdb.Del(ctx, namepace + ":" + key).Result()
	return err
}