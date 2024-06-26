package initialize

import (
	"context"
	"github.com/littlestar2023/common_pkg/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitialRedis() *redis.Client {

	redisCfg := global.CMP_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.CMP_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.CMP_LOG.Info("redis connect ping response:", zap.String("pong", pong))
	}

	return client
}
