package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"hello/handler"

	hello "hello/proto/hello"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.hello"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	_ = hello.RegisterHelloHandler(service.Server(), new(handler.Hello))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
