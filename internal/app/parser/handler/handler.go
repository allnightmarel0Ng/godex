package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/allnightmarel0Ng/godex/internal/app/parser/usecase"
	"github.com/allnightmarel0Ng/godex/internal/domain/model"
	"github.com/allnightmarel0Ng/godex/internal/logger"
	"github.com/gin-gonic/gin"
)

type ParserHandler struct {
	useCase usecase.ParserUseCase
}

func NewParserHandler(useCase usecase.ParserUseCase) ParserHandler {
	return ParserHandler{
		useCase: useCase,
	}
}

func send(c *gin.Context, code int, message string) {
	c.JSON(code, model.Response{
		Code:    code,
		Message: message,
	})
}

func (p *ParserHandler) HandleLink(c *gin.Context) {
	logger.Debug("HandleLink: start")
	defer logger.Debug("HandleLink: end")

	if c.GetHeader("Content-Type") != "application/json" {
		send(c, http.StatusBadRequest, "wrong content type: should be application/json")
		logger.Warning("wrong content type")
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		send(c, http.StatusInternalServerError, "error reading request body")
		logger.Warning("error reading request body")
		return
	}
	defer c.Request.Body.Close()

	var link model.Link
	err = json.Unmarshal(body, &link)
	if err != nil {
		send(c, http.StatusBadRequest, "error parsing JSON")
		logger.Warning("error parsing JSON")
		return
	}

	fileName, packageName, url, err := p.useCase.ParseUrl(link.Link)
	if err != nil {
		message := fmt.Sprintf("invalid link: %s", err.Error())
		logger.Warning(message)
		send(c, http.StatusBadRequest, message)
		return
	}

	bytes, err := p.useCase.FetchFile(url)
	if err != nil {
		message := fmt.Sprintf("unable to fetch the data from link: %s", err.Error())
		logger.Warning(message)
		send(c, http.StatusNotFound, message)
		return
	}

	functions, err := p.useCase.ExtractFunctions(bytes, fileName, packageName, url)
	if err != nil {
		message := fmt.Sprintf("unable to get functions from file: %s", err.Error())
		logger.Warning(message)
		send(c, http.StatusInternalServerError, message)
		return
	}

	for _, function := range functions {
		err = p.useCase.ProduceMessage(function)
		if err != nil {
			logger.Warning("producer error: %s", err.Error())
		}
	}

	send(c, http.StatusOK, "success")
}
