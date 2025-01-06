package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID      uint   `json:"id"`
	IDUser  int    `json:"id_user"`
	IDPost  uint   `json:"id_post"`
	Content string `json:"content_comment"`
	User    User   `json:"user"`
	Post    Post   `json:"post"`
}
