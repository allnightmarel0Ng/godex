package main

import (
	"net/http"
	"net/url"

	"github.com/allnightmarel0Ng/godex/internal/app/gateway/handler"
	"github.com/allnightmarel0Ng/godex/internal/app/gateway/usecase"
	"github.com/allnightmarel0Ng/godex/internal/config"
	"github.com/allnightmarel0Ng/godex/internal/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		logger.Error("unable to load config: %s", err.Error())
	}

	useCase := usecase.NewGatewayUseCase(
		url.URL{Scheme: "ws", Host: "parser:"+conf.ParserPort, Path: "/"}, 
		url.URL{Scheme: "ws", Host: "container:"+conf.ContainerPort, Path: "/"})
	handle := handler.NewGatewayHandler(useCase)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://localhost"},
        AllowMethods: []string{"POST"},
        AllowHeaders: []string{"Content-Type"},
    }))

	router.POST("/store", handle.HandleStore)
	router.POST("/find", handle.HandleFind)

	err = http.ListenAndServe(":"+conf.GatewayPort, router)
	if err != nil {
		logger.Error("unable to listen and serve: %s", err.Error())
	}
}
