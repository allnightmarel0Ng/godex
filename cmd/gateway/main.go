package main

import (
	"net/http"
	"net/url"

	"github.com/allnightmarel0Ng/godex/internal/app/gateway/handler"
	"github.com/allnightmarel0Ng/godex/internal/app/gateway/usecase"
	"github.com/allnightmarel0Ng/godex/internal/config"
	"github.com/allnightmarel0Ng/godex/internal/logger"
	"github.com/go-chi/chi/v5"
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

	router := chi.NewRouter()

	router.Post("/store", handle.HandleStore)
	router.Get("/find", handle.HandleFind)

	err = http.ListenAndServe(":"+conf.GatewayPort, router)
	if err != nil {
		logger.Error("unable to listen and serve: %s", err.Error())
	}
}
