package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateProduct(ctx *gin.Context) {
	// createdProduct, err := logic.CreateProduct(req)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	ctx.JSON(http.StatusCreated, "something")
}
