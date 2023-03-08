package main

import (
	"github.com/FUnigrad/funiverse-workspace-service/internal/goclient"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/pods", func(ctx *gin.Context) {
		data, err := goclient.GetPods()
		if err != nil {
			ctx.JSON(500, err.Error())
		}
		ctx.JSON(200, data)
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}
