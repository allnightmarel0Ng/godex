package main

import (
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"

	containerpb "github.com/allnightmarel0Ng/godex/internal/app/container/proto"
	"github.com/allnightmarel0Ng/godex/internal/app/entrypoint/handler"
	"github.com/allnightmarel0Ng/godex/internal/app/entrypoint/usecase"
	parserpb "github.com/allnightmarel0Ng/godex/internal/app/parser/proto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"
)

func main() {
	parserConn, err := grpc.NewClient("parser:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("unable to create the grpc client: %s", err.Error())
	}
	defer parserConn.Close()
	parserClient := parserpb.NewParserClient(parserConn)

	containerConn, err := grpc.NewClient("container:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("unable to listen and serve: %s", err.Error())
	}
}
