package model

type Follower struct {
	ID         uint `json:"id"`
	IDUser     int  `json:"id_user"`
	IDFollower int  `json:"id_follower"`
	User       User `json:"user"`
	Follower   User `json:"follower"`
}
