package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/makepeak/steve-operation/orders-service/handler"
	"github.com/makepeak/steve-operation/orders-service/subscriber"

	order "github.com/makepeak/steve-operation/orders-service/proto/order"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("steve.user.service.order"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	order.RegisterOrderHandler(service.Server(), new(handler.Order))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("steve.user.service.order", service.Server(), new(subscriber.Order))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
