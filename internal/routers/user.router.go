package routers

import (
	"github.com/Ravictation/golang_backend_coffeeshop/internal/handlers"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// ! /movie
func user(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	// dependcy injection
	repo := repositories.NewUser(d)
	handler := handlers.NewUser(repo)

	route.POST("/", handler.PostData)
	route.PATCH("/:id_user", handler.UpdateData)
	route.GET("/:id_user", handler.GetDataUser)
	route.DELETE("/:id_user", handler.DeleteData)

}
