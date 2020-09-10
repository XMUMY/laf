package main

import (
	"github.com/XMUMY/api/core/auth"
	"github.com/XMUMY/lib/micro/wrapper"
	"github.com/XMUMY/lost_found/handler"
	"github.com/XMUMY/lost_found/proto/lost_found"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("xmux.lost_found.v4"),
		micro.Version("latest"),
		micro.Address(":9000"),
		micro.WrapHandler(wrapper.ErrorLoggerHandler(lostfound.SvcID)),
	)

	// Remove default auth handler.
	wrapper.RemoveAuthHandler(service)

	// Initialise service
	service.Init()
	auth.InitAuthService(service.Client())

	// Register Handler
	_ = lostfound.RegisterLostAndFoundHandler(service.Server(), handler.New())

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
