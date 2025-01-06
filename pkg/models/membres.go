package model

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	ID           uint       `json:"id"`
	IDUser       int        `json:"id_user"`
	IDDiscussion int        `json:"id_discussion"`
	User         User       `json:"user"`
	Discussion   Discussion `json:"discussion"`
}
