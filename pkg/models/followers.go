package model

import "gorm.io/gorm"

// Follower represents a follower relationship between users
// @Description Follower represents a follower relationship between users
type Follower struct {
	gorm.Model
	ID         uint `json:"id"`
	IDUser     int  `json:"id_user"`
	IDFollower int  `json:"id_follower"`
	User       User `json:"user"`
	Follower   User `json:"follower"`
}
