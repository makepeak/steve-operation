package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/makepeak/steve-operation/payment-service/handler"
	"github.com/makepeak/steve-operation/payment-service/subscriber"

	payment "github.com/makepeak/steve-operation/payment-service/proto/payment"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("steve.user.service.payment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	payment.RegisterPaymentHandler(service.Server(), new(handler.Payment))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("steve.user.service.payment", service.Server(), new(subscriber.Payment))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
