package main

import (
	"log"
	"net/http"

	containerpb "github.com/allnightmarel0Ng/godex/internal/app/container/proto"
	"github.com/allnightmarel0Ng/godex/internal/app/entrypoint/handler"
	"github.com/allnightmarel0Ng/godex/internal/app/entrypoint/usecase"
	parserpb "github.com/allnightmarel0Ng/godex/internal/app/parser/proto"
	"github.com/allnightmarel0Ng/godex/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("unable to load config")
	}

	parserConn, err := grpc.NewClient("parser:"+conf.ParserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("unable to create the grpc client: %s", err.Error())
	}
	defer parserConn.Close()
	parserClient := parserpb.NewParserClient(parserConn)

	containerConn, err := grpc.NewClient("container:"+conf.ContainerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("unable to create the grpc client: %s", err.Error())
	}
	defer containerConn.Close()
	containerClient := containerpb.NewContainerClient(containerConn)

	useCase := usecase.NewEntrypointUseCase(parserClient, containerClient)
	handle := handler.NewEntrypointHandler(useCase)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/store", handle.HandleStore)
	router.Get("/find", handle.HandleFind)

	err = http.ListenAndServe(":"+conf.EntrypointPort, router)
	if err != nil {
		log.Fatalf("unable to listen and serve: %s", err.Error())
	}
}
