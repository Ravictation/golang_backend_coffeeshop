package handlers

import (
	"github.com/Ravictation/golang_backend_coffeeshop/internal/models"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/repositories"

	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerProduct struct {
	*repositories.RepoProduct
}

func NewProduct(r *repositories.RepoProduct) *HandlerProduct {
	return &HandlerProduct{r}
}

func (h *HandlerProduct) PostProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.CreateProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, response)
}

func (h *HandlerProduct) GetAllDataProduct(ctx *gin.Context) {

	var product models.Product
	product.Id_product = ctx.Param("id_product")

	if err := ctx.ShouldBindUri(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := h.GetAllProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, respone)
}

func (h *HandlerProduct) GetDataProduct(ctx *gin.Context) {

	var product models.Product
	product.Id_product = ctx.Param("id_product")

	if err := ctx.ShouldBindUri(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := h.GetProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, respone)
}

func (h *HandlerProduct) UpdateData(ctx *gin.Context) {

	var product models.Product
	product.Id_product = ctx.Param("id_product")

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := h.UpdateProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, respone)
}

func (h *HandlerProduct) DeleteData(ctx *gin.Context) {

	var product models.Product

	if err := ctx.ShouldBindUri(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := h.DeleteProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, respone)
}
