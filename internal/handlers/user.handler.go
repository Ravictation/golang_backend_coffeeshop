package handlers

import (
	"fmt"
	"net/http"

	"github.com/Ravictation/golang_backend_coffeeshop/internal/models"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/repositories"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	*repositories.RepoUser
}

func NewUser(r *repositories.RepoUser) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) PostData(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := h.CreateUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, respone)
}

func (h *HandlerUser) UpdateData(ctx *gin.Context) {

	var user models.User
	user.Id_user = ctx.Param("id_user")

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Print(user)

	respone, err := h.UpdateUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, respone)
}

func (h *HandlerUser) GetDataUser(ctx *gin.Context) {

	var user models.User
	user.Id_user = ctx.Param("id_user")

	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := h.GetUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, respone)
}

func (h *HandlerUser) GetAllData(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := h.GetAllUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, respone)
}

func (h *HandlerUser) DeleteData(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := h.DeleteUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, respone)
}
