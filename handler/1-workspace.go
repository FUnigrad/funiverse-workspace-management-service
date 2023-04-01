package handler

import (
	"net/http"
	"strconv"

	"github.com/FUnigrad/funiverse-workspace-service/model"
	"github.com/gin-gonic/gin"
)

func GetTokenFromRequest(ctx *gin.Context) string {
	token := ctx.Request.Header["Authorization"][0]
	return token
}

/*
Get a list of workspace
*/
func (server *Server) GetAllWorkspace(ctx *gin.Context) {

	if len(ctx.Request.Header["Authorization"]) == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := GetTokenFromRequest(ctx)

	workspaces := server.WorkspaceSerive.GetAllWorkspace(token)

	ctx.JSON(http.StatusOK, workspaces)

}

/*
Get workspace by its ID -> return a workspace
*/
func (server *Server) GetWorkspaceById(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	token := GetTokenFromRequest(ctx)

	workspace := server.WorkspaceSerive.GetWorkspaceById(id, token)

	if workspace == nil {
		ctx.JSON(http.StatusNotFound, workspace)
	} else {
		ctx.JSON(http.StatusOK, workspace)
	}

}

func (server *Server) CreateWorkspace(ctx *gin.Context) {
	var workspace model.WorkspaceDTO

	if err := ctx.ShouldBindJSON(&workspace); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token := GetTokenFromRequest(ctx)
	created_workspace, err := server.WorkspaceSerive.CreateWorkspace(workspace, token)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusCreated, created_workspace)
	}

}

func (server *Server) DeleteWorkspace(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	token := GetTokenFromRequest(ctx)

	workspace := server.WorkspaceSerive.GetWorkspaceById(id, token)

	err = server.WorkspaceSerive.DeleteWorkspace(*workspace, token)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Delete ok",
		})
	}

}
