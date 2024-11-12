package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/allnightmarel0Ng/godex/internal/app/gateway/handler"
	"github.com/allnightmarel0Ng/godex/internal/app/gateway/repository"
	"github.com/allnightmarel0Ng/godex/internal/app/gateway/usecase"
	"github.com/allnightmarel0Ng/godex/internal/config"
	domainRepository "github.com/allnightmarel0Ng/godex/internal/domain/repository"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/postgres"
	"github.com/allnightmarel0Ng/godex/internal/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		logger.Error("unable to load config: %s", err.Error())
	}

	db, err := postgres.NewDatabase(context.Background(), fmt.Sprintf("postgresql://%s:%s@postgres:%s/%s?sslmode=disable", conf.PostgresUser, conf.PostgresPassword, conf.PostgresPort, conf.PostgresDB))
	if err != nil {
		logger.Error("unable to connect to the database: %s", err.Error())
	}
	defer db.Close()

	repository := repository.NewGatewayRepositiry(domainRepository.NewFunctionRepository(db))
	useCase := usecase.NewGatewayUseCase(repository, conf.ParserPort)
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
