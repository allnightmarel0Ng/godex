package main

import (
	"net/http"

	containerpb "github.com/allnightmarel0Ng/godex/internal/app/container/proto"
	"github.com/allnightmarel0Ng/godex/internal/app/entrypoint/handler"
	"github.com/allnightmarel0Ng/godex/internal/app/entrypoint/usecase"
	parserpb "github.com/allnightmarel0Ng/godex/internal/app/parser/proto"
	"github.com/allnightmarel0Ng/godex/internal/config"
	"github.com/allnightmarel0Ng/godex/internal/logger"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		logger.Error("unable to load config: %s", err.Error())
	}

	parserConn, err := grpc.NewClient("parser:"+conf.ParserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("unable to create the grpc client: %s", err.Error())
	}
	defer parserConn.Close()
	parserClient := parserpb.NewParserClient(parserConn)

	containerConn, err := grpc.NewClient("container:"+conf.ContainerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("unable to create the grpc client: %s", err.Error())
	}
	defer containerConn.Close()
	containerClient := containerpb.NewContainerClient(containerConn)

	useCase := usecase.NewEntrypointUseCase(parserClient, containerClient)
	handle := handler.NewEntrypointHandler(useCase)

	router := chi.NewRouter()

	router.Post("/store", handle.HandleStore)
	router.Get("/find", handle.HandleFind)

	err = http.ListenAndServe(":"+conf.EntrypointPort, router)
	if err != nil {
		logger.Error("unable to listen and serve: %s", err.Error())
	}
}
