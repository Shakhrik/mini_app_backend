package handler

import (
	"fmt"
	"net/http"

	"github.com/Shakhrik/mini_app_backend/dto"
	"github.com/Shakhrik/mini_app_backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

func logAndReturnError(c *gin.Context, log logger.Logger, statusCode int, msg string, err error) {
	if statusCode == http.StatusNotFound {
		msg = err.Error()
	}
	fmt.Println()
	log.Error(msg, logger.Int("code", statusCode), logger.Error(err))
	c.JSON(statusCode, dto.ErrorModel{
		Message: msg,
	})
}
