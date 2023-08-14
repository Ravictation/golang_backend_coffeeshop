package middleware

import (
	"fmt"
	"net/http"

	"github.com/Ravictation/golang_backend_coffeeshop/internal/pkg"
	"github.com/gin-gonic/gin"
)

func UploadFile(fieldName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, err := ctx.FormFile(fieldName)
		if err != nil {
			if err.Error() == "http: no such file" {
				ctx.Set("image", "")
				ctx.Next()
				return
			}

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing file"})
			return
		}

		// Open the file
		src, err := file.Open()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
			return
		}
		defer src.Close()

		result, err := pkg.CloudInary(src)
		fmt.Print(result)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
			return
		}

		ctx.Set("image", result)
		ctx.Next()
	}
}
