package handler

import (
	"net/http"

	"github.com/FUnigrad/funiverse-workspace-service/model"
	"github.com/gin-gonic/gin"
)

func (server *Server) GetAllWorkspace(ctx *gin.Context) {

	ctx.AbortWithStatus(http.StatusNotImplemented)

}

func (server *Server) GetWorkspaceById(ctx *gin.Context) {

	ctx.AbortWithStatus(http.StatusNotImplemented)
}

func (server *Server) CreateWorkspace(ctx *gin.Context) {

	var workspace model.Workspace
	if err := ctx.ShouldBindJSON(&workspace); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := server.GoClient.CreateWorkspace(workspace)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "create OK",
	})
}

func (server *Server) DeleteWorkspace(ctx *gin.Context) {
	var workspace model.Workspace
	if err := ctx.ShouldBindJSON(&workspace); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := server.GoClient.DeleteWorkspace(workspace)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete OK",
	})
}
