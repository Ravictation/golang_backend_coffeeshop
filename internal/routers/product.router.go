package routers

import (
	"github.com/Ravictation/golang_backend_coffeeshop/internal/handlers"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// ! /movie
func movie(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	// dependcy injection
	repo := repositories.NewProduct(d)
	handler := handlers.NewProduct(repo)

	route.POST("/", handler.PostProduct)
	route.PATCH("/", handler.UpdateData)
	route.GET("/", handler.GetDataProduct)
	route.DELETE("/", handler.DeleteData)

}
