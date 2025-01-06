package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ID           uint       `json:"id"`
	IDUser       int        `json:"id_user"`
	IDDiscussion int        `json:"id_discussion"`
	Content      string     `json:"content_message"`
	User         User       `json:"user"`
	Discussion   Discussion `json:"discussion"`
}
