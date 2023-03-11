package model

type Workspace struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Admin  User   `json:"admin"`
	Domain string `json:"domain"`
}
