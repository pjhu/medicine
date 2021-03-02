package cache

import (
	"time"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"

	log "github.com/sirupsen/logrus"
)

// RDB for redis client
var rdb *redis.Client
// CTX for context
var ctx = context.Background()

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:56379",
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

// Set put value to cache
func Set (namepace string, key string, value interface {}) {
	serialized, err := json.Marshal(value)
	log.Info(serialized)
	if err != nil {
		log.Error(err)
		return
	}
	v := rdb.Set(ctx, namepace + ":" + key, serialized, 0)
	log.Info("set:  " + v.FullName())
}

// Get get value from cache
func Get(namepace string, key string, v interface {}) error {
	serialized, err := rdb.Get(ctx, namepace + ":" + key).Result()
	if err != nil {
		log.Error(err)
		return err
	}
	err = json.Unmarshal([]byte(serialized), &v)
	return err
} 

// Delete value form cache
func Delete(namepace string, key string) {
	_, err := rdb.Del(ctx, namepace + ":" + key).Result()
	if err != nil {
		log.Error(err)
		return
	}
}