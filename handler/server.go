package handler

import (
	"fmt"

	"github.com/FUnigrad/funiverse-workspace-service/config"
	"github.com/FUnigrad/funiverse-workspace-service/service"
	"github.com/gin-gonic/gin"
)

/*
Server contain service
*/
type Server struct {
	config          config.Config
	Router          *gin.Engine
	WorkspaceSerive *service.WorkspaceService
}

func NewServer(config config.Config) *Server {
	//Create Service
	workService := service.NewWorkspaceService(config)
	server := &Server{
		config:          config,
		WorkspaceSerive: workService,
	}

	router := gin.Default()
	//Add Route
	router.GET("/", server.HealthCheck)

	router.GET("/workspace", server.GetAllWorkspace)
	router.GET("/workspace/:id", server.GetWorkspaceById)
	router.POST("/workspace", server.CreateWorkspace)
	router.DELETE("/workspace", server.DeleteWorkspace)

	server.Router = router

	return server
}

func (server *Server) Start() error {

	config := server.config

	address := fmt.Sprintf("0.0.0.0:%s", config.Port)
	return server.Router.Run(address)
}
