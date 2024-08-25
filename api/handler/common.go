package handler

import (
	"github.com/Shakhrik/mini_app_backend/dto"
	"github.com/Shakhrik/mini_app_backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

func logAndReturnError(c *gin.Context, log logger.Logger, statusCode int, msg string, err error) {
	log.Error(msg, logger.Int("code", statusCode), logger.Error(err))
	c.JSON(statusCode, dto.ErrorModel{
		Message: msg,
	})
}
