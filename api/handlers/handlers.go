package handlers

import (
	"github.com/saidamir98/udevs_pkg/logger"
	"github.com/web-gopro/e_market_api/redis"
	"github.com/web-gopro/e_market_api/services"
)

type Handlers struct {
	log      logger.LoggerI
	services services.ServiceManagerI
	redis    redis.RedisRepoI
}

func NewHandler(log logger.LoggerI, services services.ServiceManagerI, redis redis.RedisRepoI) Handlers {

	return Handlers{log: log, services: services, redis: redis}
}
