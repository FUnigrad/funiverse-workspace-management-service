package handler

import (
	"fmt"
	"net/http"

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

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Request.Header["Authorization"]) == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}
}

func NewServer(config config.Config) *Server {
	//Create Service
	workService := service.NewWorkspaceService(config)
	server := &Server{
		config:          config,
		WorkspaceSerive: workService,
	}

	router := gin.New()

	// Global middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//Healcheck
	router.GET("/", server.HealthCheck)

	authorized := router.Group("/")

	authorized.Use(AuthRequired())
	{
		authorized.GET("/workspace", server.GetAllWorkspace)
		authorized.GET("/workspace/:id", server.GetWorkspaceById)
		authorized.POST("/workspace", server.CreateWorkspace)
		authorized.DELETE("/workspace/:id", server.DeleteWorkspace)
	}

	server.Router = router

	return server
}

func (server *Server) Start() error {

	config := server.config

	if config.Enviroment == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	address := fmt.Sprintf("0.0.0.0:%s", config.Port)

	err := server.Router.Run(address)
	return err
}
