package main

import (
	"log"
	"os"

	"github.com/Komilov31/distributed-cut/pkg/cut"
	"github.com/Komilov31/distributed-cut/pkg/flags"
	"github.com/Komilov31/distributed-cut/pkg/handler"
	"github.com/Komilov31/distributed-cut/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/wb-go/wbf/zlog"
)

func main() {
	address := os.Getenv("SERVER_ADDRESS")

	flags := flags.Parse()
	cut := cut.New(flags)
	service := service.New(cut)
	handler := handler.New(service)

	engine := gin.New()
	engine.POST("/process", handler.HandleInput)

	zlog.Logger.Info().Msg("starting server on " + address)
	if err := engine.Run(address); err != nil {
		log.Fatal("could not start serever")
	}
}
