package handler

import (
	"fmt"

	"github.com/FUnigrad/funiverse-workspace-service/config"
	"github.com/FUnigrad/funiverse-workspace-service/goclient"
	"github.com/FUnigrad/funiverse-workspace-service/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	GoClient        goclient.GoClient
	Router          *gin.Engine
	UserService     service.IUserService
	WorkspaceSerive service.IWorkspaceService
}

func NewServer(goClient goclient.GoClient) *Server {

	userService := service.NewUserService()
	workService := service.NewWorkspaceService()

	server := &Server{
		GoClient:        goClient,
		UserService:     userService,
		WorkspaceSerive: workService,
	}
	router := gin.Default()
	//Add Route
	router.GET("/", HealthCheck())

	router.GET("/workspace", server.GetAllWorkspace)
	router.GET("/workspace/:id", server.GetWorkspaceById)
	router.POST("/workspace", server.CreateWorkspace)
	router.DELETE("/workspace", server.DeleteWorkspace)

	server.Router = router

	return server
}

func (server *Server) Start(config config.Config) error {

	address := fmt.Sprintf("0.0.0.0:%s", config.Port)
	return server.Router.Run(address)
}
