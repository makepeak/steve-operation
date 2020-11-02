package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/makepeak/steve-operation/inventory-service/handler"
	"github.com/makepeak/steve-operation/inventory-service/subscriber"

	inventory "github.com/makepeak/steve-operation/inventory-service/proto/inventory"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("steve.user.service.inventory"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	inventory.RegisterInventoryHandler(service.Server(), new(handler.Inventory))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("steve.user.service.inventory", service.Server(), new(subscriber.Inventory))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
