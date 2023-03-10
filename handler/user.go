package handler

import (
	"net/http"

	"github.com/FUnigrad/funiverse-workspace-service/model"
	"github.com/gin-gonic/gin"
)

func (server *Server) GetUser(ctx *gin.Context) {
	user := server.UserService.Get()
	ctx.JSON(http.StatusOK, user)
}

func (server *Server) AddUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	server.UserService.Add(user)

	ctx.JSON(http.StatusAccepted, user)

}
