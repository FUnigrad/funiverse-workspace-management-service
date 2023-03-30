package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetAllWorkspace(ctx *gin.Context) {

	token := ctx.Request.Header["Authorization"][0]
	workspaces := server.WorkspaceSerive.GetAllWorkspace(token)

	ctx.JSON(http.StatusOK, workspaces)

}

func (server *Server) GetWorkspaceById(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	token := ctx.Request.Header["Authorization"][0]

	workspace := server.WorkspaceSerive.GetWorkspaceById(id, token)

	if workspace == nil {
		ctx.JSON(http.StatusNotFound, workspace)
	} else {
		ctx.JSON(http.StatusOK, workspace)
	}

}

func (server *Server) CreateWorkspace(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{
		"func": "CreateWorkspace",
	})
}

func (server *Server) DeleteWorkspace(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{
		"func": "DeleteWorkspace",
	})
}

// func (server *Server) CreateWorkspace(ctx *gin.Context) {

// 	var workspace model.Workspace
// 	if err := ctx.ShouldBindJSON(&workspace); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	err := server.GoClient.CreateWorkspace(workspace)

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": "create OK",
// 	})
// }

// func (server *Server) DeleteWorkspace(ctx *gin.Context) {
// 	var workspace model.Workspace
// 	if err := ctx.ShouldBindJSON(&workspace); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	err := server.GoClient.DeleteWorkspace(workspace)

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": "delete OK",
// 	})
// }
