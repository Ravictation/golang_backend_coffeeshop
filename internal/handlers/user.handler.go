package handlers

import (
	"net/http"

	"github.com/Ravictation/golang_backend_coffeeshop/config"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/models"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/pkg"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/repositories"
	"github.com/asaskevich/govalidator"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	*repositories.RepoUser
}

func NewUser(r *repositories.RepoUser) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) PostData(ctx *gin.Context) {
	var ers error
	user := models.User{
		Role: "user",
	}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	_, ers = govalidator.ValidateStruct(&user)
	if ers != nil {
		pkg.NewRes(401, &config.Result{
			Data: ers.Error(),
		}).Send(ctx)
		return
	}

	user.Password, ers = pkg.HashPassword(user.Password)
	if ers != nil {
		pkg.NewRes(401, &config.Result{
			Data: ers.Error(),
		}).Send(ctx)
		return
	}

	response, err := h.CreateUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"data":        response,
	})
}

func (h *HandlerUser) UpdateData(ctx *gin.Context) {

	var user models.User
	user.Id_user = ctx.Param("id_user")

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.UpdateUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"data":        response,
	})
}

func (h *HandlerUser) GetDataUser(ctx *gin.Context) {
	var user models.User
	user.Id_user = ctx.Param("id_user")

	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.GetUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"data":        response,
	})
}

func (h *HandlerUser) GetAllData(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.GetAllUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"data":        response,
	})
}

func (h *HandlerUser) DeleteData(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.DeleteUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"data":        response,
	})
}
