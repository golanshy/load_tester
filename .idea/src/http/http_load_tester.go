package http

import (
	"github.com/gin-gonic/gin"
	data_models "github.com/golanshy/load_tester/model"
	"github.com/golanshy/load_tester/services"
	"github.com/golanshy/plime_core-go/logger"
	"github.com/golanshy/plime_core-go/utils/rest_errors"
	"net/http"
)

type LoadTestersHandler interface {
	Ping(*gin.Context)
	LoadTest(*gin.Context)
}

type loadTestersHandler struct {
	service services.Service
}

func NewHandler(service services.Service) LoadTestersHandler {
	return &loadTestersHandler{
		service: service,
	}
}

func (handler *loadTestersHandler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (handler *loadTestersHandler) LoadTest(c *gin.Context) {

	var request data_models.LoadTestRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("error when trying to create load test request invalid json body", err)
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	err := handler.service.LoadTest(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, request)
}
