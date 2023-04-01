package service

import (
	"errors"
	"log"

	"github.com/FUnigrad/funiverse-workspace-service/config"
	"github.com/FUnigrad/funiverse-workspace-service/goclient"
	httpclient "github.com/FUnigrad/funiverse-workspace-service/http-client"
	"github.com/FUnigrad/funiverse-workspace-service/model"
)

/*
Workspaces Service contain 4 function:
 1. Get All Workspace
 2. Get Workspace id
 3. Create Workspace
 4. Delete Workspace

Workspace service communicate with 2 client:
 1. GoClient -> interact to k8s
 2. HttpClient -> interact to authenservice
*/

type WorkspaceService struct {
	goClient   *goclient.GoClient
	httpClient *httpclient.HttpClient
}

func NewWorkspaceService(config config.Config) *WorkspaceService {
	goClient, err := goclient.NewClient(config)

	if err != nil {
		log.Fatal(err)
	}

	httpClient, err := httpclient.NewClient(config)

	if err != nil {
		log.Fatal(err)
	}

	return &WorkspaceService{
		goClient:   goClient,
		httpClient: httpClient,
	}
}
func (service *WorkspaceService) GetAllWorkspace(token string) []model.Workspace {

	workspaces := service.httpClient.GetAllWorkspace(token)

	return workspaces
}

func (service *WorkspaceService) GetWorkspaceById(id int, token string) *model.Workspace {

	workspace := service.httpClient.GetWorkspaceById(id, token)

	return workspace
}

func (service *WorkspaceService) DeleteWorkspace(workspace model.Workspace, token string) (err error) {

	err = service.goClient.DeleteWorkspace(workspace)

	if err != nil {
		return
	}

	is_deleted := service.httpClient.DeleteWorkspace(workspace.Id, token)
	if is_deleted {
		return nil
	} else {
		return errors.New("cannot delete ws resources")
	}

}

func (service *WorkspaceService) CreateWorkspace(workspace model.WorkspaceDTO, token string) (*model.Workspace, error) {

	if err := service.goClient.CreateWorkspace(workspace); err != nil {
		return nil, err
	}

	result, err := service.httpClient.CreateWorkspace(workspace, token)

	return result, err
}
