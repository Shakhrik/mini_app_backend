package handler

import (
	"net/http"

	model "github.com/Shakhrik/mini_app_backend/api/swagger"
	"github.com/Shakhrik/mini_app_backend/dto"
	"github.com/Shakhrik/mini_app_backend/service"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService service.ProductI
}

func newProductHandler(productService service.ProductI) *productHandler {
	return &productHandler{
		productService: productService,
	}
}

func (h productHandler) CreateProduct(c *gin.Context) {
	// createdProduct, err := logic.CreateProduct(req)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	req := model.ProductRequest{}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	prodDto := dto.Product{
		Id:    1,
		Name:  "shakh",
		Price: 21,
	}
	p, err := h.productService.CreateProduct(prodDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
	}
	c.JSON(http.StatusCreated, model.CreateProduct201JSONResponse{
		Id: &p.Id,
	})
}
