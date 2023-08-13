package handlers

import (
	"github.com/Ravictation/golang_backend_coffeeshop/config"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/pkg"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/repositories"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `db:"username" json:"username" form:"username"`
	Password string `db:"password" json:"password,omitempty"`
}

type HandlerAuth struct {
	*repositories.RepoUser
}

func NewAuth(r *repositories.RepoUser) *HandlerAuth {
	return &HandlerAuth{r}
}

func (h *HandlerAuth) Login(ctx *gin.Context) {
	var data User
	if ers := ctx.ShouldBind(&data); ers != nil {
		pkg.NewRes(500, &config.Result{
			Data: ers.Error(),
		}).Send(ctx)
		return
	}

	users, err := h.GetAuthData(data.Username)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	if err := pkg.VerifyPassword(users.Password, data.Password); err != nil {
		pkg.NewRes(401, &config.Result{
			Data: "wrong password",
		}).Send(ctx)
		return
	}

	jwtt := pkg.NewToken(users.Id_user, users.Role)
	token, err := jwtt.Generate()
	if err != nil {
		pkg.NewRes(500, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, &config.Result{Data: token}).Send(ctx)

}
