package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	ID        uint    `json:"id"`
	IDPost    int     `json:"id_post"`
	IDUser    int     `json:"id_user"`
	IDComment int     `json:"id_comment"` // (peut être NULL)
	Post      Post    `json:"post"`
	User      User    `json:"user"`
	Comment   Comment `json:"comment"`
}
