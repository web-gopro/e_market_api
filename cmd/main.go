package main

import (
	"context"
	"log"

	"github.com/saidamir98/udevs_pkg/logger"
	"github.com/web-gopro/e_market_api/api"
	"github.com/web-gopro/e_market_api/config"
	"github.com/web-gopro/e_market_api/pkg/db"
	"github.com/web-gopro/e_market_api/redis"
	"github.com/web-gopro/e_market_api/services"
)

var ctx = context.Background()

func main() {

	logr := logger.NewLogger("", logger.LevelDebug)
	cfg := config.Load()

	service := services.Service()

	log.Println(service)

	rediCli, err := db.ConnRedis(logr, context.Background(), cfg.RedisConfig)

	if err != nil {

		logr.Debug(err.Error())
		return

	}

	cache := redis.NewRedisRepo(rediCli, logr)

	engine := api.Api(api.Options{Services: service, Redis: cache, Log: logr})

	engine.Run(":8080")

}



//chmod +x ./scripts/gen_proto.sh

// ./scripts/gen_proto.sh
