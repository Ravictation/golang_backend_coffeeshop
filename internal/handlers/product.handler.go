package handlers

import (
	"net/http"
	"strconv"

	"github.com/Ravictation/golang_backend_coffeeshop/config"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/models"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/pkg"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/repositories"
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

	product.Product_image = ctx.MustGet("image").(string)
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.CreateProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	pkg.NewRes(200, &config.Result{Data: response}).Send(ctx)
}

func (h *HandlerProduct) GetAllDataProduct(ctx *gin.Context) {

	var product models.Product
	search := ctx.Query("search")
	categories := ctx.Query("categories")
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, pgnt, err := h.GetAllProduct(search, page, limit, categories)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"data":        response,
		"meta":        pgnt,
	})
}

func (h *HandlerProduct) GetDataProduct(ctx *gin.Context) {

	var product models.Product
	product.Id_product = ctx.Param("id_product")

	if err := ctx.ShouldBindUri(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.GetProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pkg.NewRes(200, &config.Result{Data: response}).Send(ctx)
}

func (h *HandlerProduct) UpdateData(ctx *gin.Context) {

	var product models.Product
	product.Id_product = ctx.Param("id_product")
	product.Product_image = ctx.MustGet("image").(string)

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.UpdateProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pkg.NewRes(200, &config.Result{Data: response}).Send(ctx)
}

func (h *HandlerProduct) DeleteData(ctx *gin.Context) {

	var product models.Product
	product.Id_product = ctx.Param("id_product")

	if err := ctx.ShouldBindUri(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.DeleteProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pkg.NewRes(200, &config.Result{Data: response}).Send(ctx)
}
