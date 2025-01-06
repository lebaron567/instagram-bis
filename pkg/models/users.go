package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID             uint   `json:"id"`
	LastName       string `json:"lastename_user"`
	FirstName      string `json:"firstname_user"`
	Email          string `json:"email_user"`
	Pseudo         string `json:"pseudo_user"`
	Birthdate      string `json:"birthdate"`
	Password       string `json:"password_user"`
	IsPrivate      bool   `json:"isprivate_user"`
	ProfilePicture string `json:"profilpicture_user"`
	WantsNotify    bool   `json:"wantsnotify_user"`
}
