package main

import (
	"search-service/internal"
	"search-service/internal/util"

	"github.com/hadanhtuan/go-sdk"
	"github.com/hadanhtuan/go-sdk/amqp"
	"github.com/hadanhtuan/go-sdk/config"
	"github.com/hadanhtuan/go-sdk/db/elasticsearch"
	cache "github.com/hadanhtuan/go-sdk/db/redis"
)

func main() {
	config, _ := config.InitConfig(".")

	amqp.ConnectRabbit(util.SEARCH_EXCHANGE, util.SEARCH_QUEUE, amqp.ExchangeType.Topic)
	cache.ConnectRedis()
	es.ConnectElasticSearch()

	app := sdk.App{
		Config: config,
	}

	internal.InitGRPCServer(&app)
}
