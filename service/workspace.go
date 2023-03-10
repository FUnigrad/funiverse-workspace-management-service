package service

import (
	"github.com/FUnigrad/funiverse-workspace-service/model"
)

type IWorkspaceService interface {
	GetAllWorkspace() []model.Workspace
	GetWorkspaceById(int64) model.Workspace
	CreateWorkspace(model.Workspace) model.Workspace
}

type WorkspaceService struct {
	worspaces []model.Workspace
}

func NewWorkspaceService() IWorkspaceService {
	return &WorkspaceService{}
}
func (service *WorkspaceService) GetAllWorkspace() []model.Workspace {
	return service.worspaces
}

func (service *WorkspaceService) GetWorkspaceById(id int64) model.Workspace {
	return service.worspaces[id]
}
func (service *WorkspaceService) CreateWorkspace(workspace model.Workspace) model.Workspace {
	service.worspaces = append(service.worspaces, workspace)
	return workspace
}
