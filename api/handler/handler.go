package handler

import (
	"github.com/Shakhrik/mini_app_backend/pkg/logger"
	"github.com/Shakhrik/mini_app_backend/service"
)

type Handler struct {
	productHandler
	telegramBotHandler
}

func NewHandler(service *service.Service, logger logger.Logger) Handler {
	return Handler{
		productHandler:     newProductHandler(service.Product),
		telegramBotHandler: newTelegramBotHandler(service.TelegramBot, logger),
	}
}
