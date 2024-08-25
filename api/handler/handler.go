package handler

import "github.com/Shakhrik/mini_app_backend/service"

type Handler struct {
	productHandler
	telegramBotHandler
}

func NewHandler(service *service.Service) Handler {
	return Handler{
		productHandler:     newProductHandler(service.Product),
		telegramBotHandler: newTelegramBotHandler(service.TelegramBot),
	}
}
