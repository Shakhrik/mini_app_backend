package api

import (
	"net/http"
	"time"

	"github.com/Shakhrik/mini_app_backend/api/handler"
	"github.com/Shakhrik/mini_app_backend/api/swagger"
	"github.com/Shakhrik/mini_app_backend/service"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer(service *service.Service) *http.Server {
	r := gin.Default()

	sw, err := swagger.GetSwagger()
	if err != nil {
		return nil
	}

	r.GET("/swagger.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, sw)
	})
	// Route to redirect to Swagger UI
	url := ginSwagger.URL("/swagger.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Use(middleware.OapiRequestValidator(sw))
	h := handler.NewHandler(service)
	swagger.RegisterHandlersWithOptions(r, h, swagger.GinServerOptions{
		BaseURL: "api/v1",
	})

	return &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
		// go-swagger uses 10 seconds which seems a bit short. http client default is 90 seconds
		IdleTimeout: 90 * time.Second,
		// protect against Slowloris: https://en.wikipedia.org/wiki/Slowloris_(computer_security)
		// go-swagger servers have a ReadTimeout (which includes headers) of 30 second
		ReadHeaderTimeout: 30 * time.Second,
		// we should have some timelimit on requests and we have th write timeout set to 10 mins in simstore
		ReadTimeout:  15 * time.Minute,
		WriteTimeout: 15 * time.Minute,
	}
}
