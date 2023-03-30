package model

type User struct {
	Id           int    `json:"id"`
	PersonalMail string `json:"personalMail"`
	EduMail      string `json:"eduMail"`
}
