package handlers

// import (
// 	"github.com/Ravictation/golang_backend_coffeeshop/internal/models"
// 	"github.com/Ravictation/golang_backend_coffeeshop/internal/repositories"

// 	// "fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type HandlerCustOrder struct {
// 	*repositories.RepoCustOrder
// }

// func NewCustOrder(r *repositories.RepoCustOrder) *HandlerCustOrder {
// 	return &HandlerCustOrder{r}
// }

// func (h *HandlerCustOrder) PostData(ctx *gin.Context) {
// 	var CustOrder models.CustOrder

// 	if err := ctx.ShouldBind(&CustOrder); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	response, err := h.CreateOrder(&CustOrder)
// 	if err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}
// 	ctx.JSON(200, response)
// }

// // func (h *HandlerUser) GetDataUser(ctx *gin.Context) {

// // 	var user models.User
// // 	user.Id_user = ctx.Param("id_user")

// // 	if err := ctx.ShouldBindUri(&user); err != nil {
// // 		ctx.AbortWithError(http.StatusBadRequest, err)
// // 		return
// // 	}

// // 	respone, err := h.GetUser(&user)
// // 	if err != nil {
// // 		ctx.AbortWithError(http.StatusBadRequest, err)
// // 		return
// // 	}

// // 	ctx.JSON(200, respone)
// // }

// func (h *HandlerCustOrder) GetAllDataorder(ctx *gin.Context) {

// 	var CustOrder models.CustOrder

// 	if err := ctx.ShouldBind(&CustOrder); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	response, err := h.GetAllData(&CustOrder)
// 	if err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	ctx.JSON(200, response)
// }

// func (h *HandlerCustOrder) UpdateData(ctx *gin.Context) {

// 	var CustOrder models.CustOrder
// 	CustOrder.Id_order = ctx.Param("id_order")

// 	if err := ctx.ShouldBind(&CustOrder); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	response, err := h.UpdateOrder(&CustOrder)
// 	if err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	ctx.JSON(200, response)
// }

// func (h *HandlerCustOrder) DeleteData(ctx *gin.Context) {

// 	var CustOrder models.CustOrder

// 	if err := ctx.ShouldBindUri(&CustOrder); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	response, err := h.DeleteOrder(&CustOrder)
// 	if err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	ctx.JSON(200, response)
// }
