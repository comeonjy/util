// @Description  redis
// @Author  	 jiangyang  
// @Created  	 2020/11/2 10:02 上午

// Example Config:
// redis:
//   addr: 127.0.0.1:6379
//   password:
//   db: 0
//   pool_size: 100

package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var (
	client *redis.Client
)

type Config struct {
	Addr     string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
	Db       int    `json:"db" yaml:"db"`
	PoolSize int    `json:"pool_size" yaml:"pool_size" mapstructure:"pool_size"`
}

func Init(cfg Config) {
	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.PoolSize,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("redis connect successfully")

}

func GetConn() *redis.Client {
	return client
}

func Close() {
	if client != nil {
		client.Close()
	}
	logrus.Info("redis connect closed")
}
