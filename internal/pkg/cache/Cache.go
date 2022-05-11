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

// CTX for context
var repo *RedisRepository
var once sync.Once
var ctx context.Context

type ICacheRepository interface {
	StoreBy(namespace string, key string, value interface{}) error
	GetBy(namespace string, key string, v interface{}) error
	DeleteBy(namespace string, key string) error
}

type RedisRepository struct {
	client *redis.Client
}

func Repository() *RedisRepository {
	return repo
}

func Builder() {
	once.Do(func() {
		initRedis()
	})
}

// initRedis for redis connection
func initRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	logrus.Error(err)
	if err != nil {
		panic("failed to connect redis")
	}
	repo = &RedisRepository{client: rdb}
}

// StoreBy put value to cache
func (repo RedisRepository) StoreBy(namespace string, key string, value interface{}) error {
	serialized, err := json.Marshal(value)
	logrus.Info(serialized)
	if err != nil {
		logrus.Error(err)
		return err
	}
	_, err = repo.client.Set(ctx, namespace+":"+key, serialized, 0).Result()
	return err
}

// GetBy get value from cache
func (repo RedisRepository) GetBy(namepace string, key string, v interface{}) error {
	serialized, err := repo.client.Get(ctx, namepace+":"+key).Result()
	if err != nil {
		logrus.Error(err)
		return err
	}
	err = json.Unmarshal([]byte(serialized), &v)
	return err
}

// DeleteBy value form cache
func (repo RedisRepository) DeleteBy(namepace string, key string) error {
	_, err := repo.client.Del(ctx, namepace+":"+key).Result()
	return err
}
