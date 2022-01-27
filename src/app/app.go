package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golanshy/load_tester/http"
	"github.com/golanshy/load_tester/repository/rest/leslie_api_repo"
	"github.com/golanshy/load_tester/services"
	"github.com/golanshy/plime_core-go/logger"
	"os"
	"strings"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApplication() {
	logger.Info("about to start the application")

	ccHandler := http.NewHandler(services.NewService(leslie_api_repo.NewRestLeslieApiRepository()))

	router.GET("/load_test/ping", ccHandler.Ping)
	router.POST("/load_test/test", ccHandler.LoadTest)

	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		port = "8080"
	}
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
		return
	}
}
