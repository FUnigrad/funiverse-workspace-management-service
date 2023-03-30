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
	Name         string `json:"name" binding:"required"`
	Code         string `json:"code" binding:"required"`
	Domain       string `json:"domain" binding:"required"`
	PersonalMail string `json:"personalMail" binding:"required,email"`
	EduMail      string `json:"eduMail" binding:"required,email"`
}
