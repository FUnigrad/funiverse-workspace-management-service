package model

type Workspace struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Admin  User   `json:"admin"`
	Domain string `json:"domain"`
	Active bool   `json:"active"`
}

type WorkspaceDTO struct {
	Name         string `json:"name"`
	Code         string `json:"code"`
	Domain       string `json:"domain"`
	PersonalMail string `json:"personalMail"`
	EduMail      string `json:"eduMail"`
}
