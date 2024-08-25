package handler

import (
	"net/http"

	model "github.com/Shakhrik/mini_app_backend/api/swagger"
	srdto "github.com/Shakhrik/mini_app_backend/dto/service"
	"github.com/Shakhrik/mini_app_backend/pkg/errors"
	"github.com/Shakhrik/mini_app_backend/pkg/logger"
	"github.com/Shakhrik/mini_app_backend/service"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/swag"
)

type telegramBotHandler struct {
	telegramBotService service.TelegramBotI
	logger             logger.Logger
}

func newTelegramBotHandler(tgBotService service.TelegramBotI, l logger.Logger) telegramBotHandler {
	return telegramBotHandler{
		telegramBotService: tgBotService,
		logger:             l,
	}
}

func (t telegramBotHandler) CreateTelegramBot(c *gin.Context) {
	req := model.CreateTelegramBotJSONRequestBody{}
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		logAndReturnError(c, t.logger, http.StatusBadRequest, "invalid request payload: "+err.Error(), err)
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
		logAndReturnError(c, t.logger, errors.HttpErrCode(err), "failed to create telegram bot", err)
		return
	}

	c.JSON(http.StatusCreated, model.IDResponse{Id: swag.String(id)})
}

func (t telegramBotHandler) DeleteTelegramBot(c *gin.Context, params model.DeleteTelegramBotParams) {
	if err := t.telegramBotService.Delete(c, params.Id); err != nil {
		logAndReturnError(c, t.logger, errors.HttpErrCode(err), "failed to delete telegram bot", err)
		return
	}

	c.JSON(http.StatusNoContent, model.DeleteTelegramBot204Response{})
}

func (t telegramBotHandler) GetTelegramBot(c *gin.Context, id string) {
	res, err := t.telegramBotService.Get(c, id)
	if err != nil {
		logAndReturnError(c, t.logger, errors.HttpErrCode(err), "failed to get telegram bot", err)
		return
	}

	c.JSON(http.StatusOK, model.GetTelegramBot200JSONResponse{
		Id:          &res.ID,
		Name:        &res.Name,
		Token:       &res.Token,
		ImageUrl:    &res.ImageUrl,
		Description: &res.Description,
	})
}
