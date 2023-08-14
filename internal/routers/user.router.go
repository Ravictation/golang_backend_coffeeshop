package routers

import (
	"github.com/Ravictation/golang_backend_coffeeshop/internal/handlers"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/middleware"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repositories.NewUser(d)
	handler := handlers.NewUser(repo)

	route.POST("/", middleware.UploadFile("image_user"), handler.PostData)
	route.PATCH("/:username", middleware.UploadFile("image_user"), handler.UpdateData)
	route.GET("/", handler.GetAllData)
	route.GET("/:username", middleware.Authjwt("admin"), handler.GetDataUser)
	route.DELETE("/:username", handler.DeleteData)

}
