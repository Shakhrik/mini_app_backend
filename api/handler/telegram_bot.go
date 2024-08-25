package handler

import (
	"net/http"

	model "github.com/Shakhrik/mini_app_backend/api/swagger"
	srdto "github.com/Shakhrik/mini_app_backend/dto/service"
	"github.com/Shakhrik/mini_app_backend/pkg/logger"
	"github.com/Shakhrik/mini_app_backend/service"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/swag"
)

type telegramBotHandler struct {
	telegramBotService service.TelegramBotI
	logger             logger.Logger
}

func newTelegramBotHandler(tgBotService service.TelegramBotI) telegramBotHandler {
	return telegramBotHandler{
		telegramBotService: tgBotService,
	}
}

func (t telegramBotHandler) CreateTelegramBot(c *gin.Context) {
	req := model.CreateTelegramBotJSONRequestBody{}
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	params := srdto.TelegramBotCreateParams{
		Name:        req.Name,
		Token:       req.Token,
		ImageUrl:    swag.StringValue(req.ImageUrl),
		Description: swag.StringValue(req.Description),
	}
	id, err := t.telegramBotService.Create(c, &params)
	if err != nil {
		logAndReturnError(c, t.logger, http.StatusInternalServerError, "failed to create telegram bot", err)
		return
	}

	c.JSON(http.StatusCreated, model.IDResponse{Id: swag.String(id)})
}
