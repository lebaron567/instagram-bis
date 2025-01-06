package model

type Message struct {
	ID           uint       `json:"id"`
	IDUser       int        `json:"id_user"`
	IDDiscussion int        `json:"id_discussion"`
	Content      string     `json:"content_message"`
	User         User       `json:"user"`
	Discussion   Discussion `json:"discussion"`
}
