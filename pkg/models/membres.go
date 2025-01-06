package model

type Member struct {
	ID           uint       `json:"id"`
	IDUser       int        `json:"id_user"`
	IDDiscussion int        `json:"id_discussion"`
	User         User       `json:"user"`
	Discussion   Discussion `json:"discussion"`
}
