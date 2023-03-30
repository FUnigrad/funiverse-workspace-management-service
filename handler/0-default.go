package handler

import (
	"net/http"

	"github.com/FUnigrad/funiverse-workspace-service/goclient"
	"github.com/gin-gonic/gin"
)

func (server *Server) HealthCheck(ctx *gin.Context) {

	_, err := goclient.NewClient(server.config)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Ok",
			"env":     server.config.Enviroment,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
	}

}
