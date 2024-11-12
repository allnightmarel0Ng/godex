package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/allnightmarel0Ng/godex/internal/app/gateway/usecase"
	"github.com/allnightmarel0Ng/godex/internal/domain/model"
	"github.com/allnightmarel0Ng/godex/internal/logger"
	"github.com/gin-gonic/gin"
)

type GatewayHandler struct {
	useCase usecase.GatewayUseCase
}

func NewGatewayHandler(useCase usecase.GatewayUseCase) GatewayHandler {
	return GatewayHandler{
		useCase: useCase,
	}
}

func (e *GatewayHandler) sendError(c *gin.Context, statusCode int, message string) {
    c.JSON(statusCode, model.Response{
        Code:    statusCode,
        Message: message,
    })
}

func (e *GatewayHandler) HandleStore(c *gin.Context) {
    logger.Debug("HandleStore: start")
    defer logger.Debug("HandleStore: end")

    if c.GetHeader("Content-Type") != "application/json" {
        e.sendError(c, http.StatusBadRequest, "wrong content type: should be application/json")
        logger.Warning("wrong content type")
        return
    }

    body, err := io.ReadAll(c.Request.Body)
    if err != nil {
        e.sendError(c, http.StatusInternalServerError, "error reading request body")
        logger.Warning("error reading request body")
        return
    }
    defer c.Request.Body.Close()

    payload, err := e.useCase.Store(body)
    if err != nil {
        e.sendError(c, http.StatusInternalServerError, "unexpected error")
        logger.Warning("unexpected http error: %s", err.Error())
        return
    }

    var response model.Response
    if json.Unmarshal(payload, &response) != nil {
        e.sendError(c, http.StatusInternalServerError, "unexpected error")
        logger.Warning("json unmarshalling error")
        return
    }

    c.JSON(response.Code, response)
}

func (e *GatewayHandler) HandleFind(c *gin.Context) {
    logger.Debug("HandleFind: start")
    defer logger.Debug("HandleFind: end")

	
    if c.GetHeader("Content-Type") != "application/json" {
		e.sendError(c, http.StatusBadRequest, "wrong content type: should be application/json")
        logger.Warning("wrong content type")
        return
    }

    body, err := io.ReadAll(c.Request.Body)
    if err != nil {
        e.sendError(c, http.StatusInternalServerError, "error reading request body")
        logger.Warning("error reading request body")
        return
    }
    defer c.Request.Body.Close()

    var signature model.Signature
    err = json.Unmarshal(body, &signature)
    if err != nil {
        e.sendError(c, http.StatusBadRequest, "error parsing JSON")
        logger.Warning("error parsing JSON")
        return
    }

    payload, err := e.useCase.Find(signature.Signature)
    if err != nil || payload == nil {
        e.sendError(c, http.StatusNotFound, "error finding signature")
        logger.Warning("error finding signature")
        return
    }

    c.JSON(http.StatusOK, json.RawMessage(payload))
}
