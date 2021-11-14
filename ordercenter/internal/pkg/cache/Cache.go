package cache

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	once   sync.Once
	rdb *redis.Client  // RDB for redis client
	ctx = context.Background()  // CTX for context
)

func Init() {
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     viper.GetString("redis.addr"),
			Password: "",  // no password set
			DB:       0,   // use default DB
			PoolSize: 100, // 连接池大小
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err := rdb.Ping(ctx).Result()
		logrus.Error(err)
		if err != nil {
			panic("failed to connect redis")
		}
	})
}

// StoreBy put value to cache
func StoreBy(namespace string, key string, value interface {}) error {
	serialized, err := json.Marshal(value)
	logrus.Info(serialized)
	if err != nil {
		logrus.Error(err)
		return err
	}
	_, err = rdb.Set(ctx, namespace+ ":" + key, serialized, 0).Result()
	return err
}

// GetBy get value from cache
func GetBy(namespace string, key string, v interface {}) error {
	serialized, err := rdb.Get(ctx, namespace+ ":" + key).Result()
	if err != nil {
		logrus.Error(err)
		return err
	}
	err = json.Unmarshal([]byte(serialized), &v)
	return err
} 

// DeleteBy value form cache
func DeleteBy(namespace string, key string) error {
	_, err := rdb.Del(ctx, namespace+ ":" + key).Result()
	return err
}