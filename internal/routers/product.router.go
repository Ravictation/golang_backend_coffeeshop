package routers

import (
	"github.com/Ravictation/golang_backend_coffeeshop/internal/handlers"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/middleware"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	// dependcy injection
	repo := repositories.NewProduct(d)
	handler := handlers.NewProduct(repo)

	route.POST("/", middleware.Authjwt("admin"), middleware.UploadFile("product_image"), handler.PostProduct)
	route.PATCH("/:id_product", middleware.Authjwt("admin"), middleware.UploadFile("product_image"), handler.UpdateData)
	route.GET("/:id_product", handler.GetDataProduct)
	route.GET("/", handler.GetAllDataProduct)
	route.DELETE("/:id_product", middleware.Authjwt("admin"), handler.DeleteData)

}
