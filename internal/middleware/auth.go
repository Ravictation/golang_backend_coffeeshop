package middleware

import (
	"strings"

	"github.com/Ravictation/golang_backend_coffeeshop/config"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/pkg"
	"github.com/gin-gonic/gin"
)

func Authjwt(role ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var valid bool
		var header string

		if header = ctx.GetHeader("Authorization"); header == "" {
			pkg.NewRes(401, &config.Result{
				Message: "need to login first",
			}).Send(ctx)
			return
		}

		if !strings.Contains(header, "Bearer") {
			pkg.NewRes(401, &config.Result{
				Message: "Invalid Header",
			}).Send(ctx)
			return
		}

		tokens := strings.Replace(header, "Bearer ", "", -1)
		check, err := pkg.VerifyToken(tokens)
		if err != nil {
			pkg.NewRes(401, &config.Result{
				Message: err.Error(),
			}).Send(ctx)
			return
		}
		for _, r := range role {
			if r == check.Role {
				valid = true
			}
		}
		if !valid {
			pkg.NewRes(401, &config.Result{
				Data: "No Permission To Access",
			}).Send(ctx)
			return
		}
		ctx.Set("userId", check.Id)
		ctx.Next()
	}

}
