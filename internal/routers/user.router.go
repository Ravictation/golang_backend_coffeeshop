package routers

import (
	"github.com/Ravictation/golang_backend_coffeeshop/internal/handlers"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// ! /movie
func movie(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	// dependcy injection
	repo := repositories.NewUser(d)
	handler := handlers.NewUser(repo)

	route.POST("/", handler.PostData)
	route.PATCH("/", handler.UpdateData)
	route.GET("/", handler.GetDataUser)
	route.DELETE("/", handler.DeleteData)

}
