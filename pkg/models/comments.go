package model

type Comment struct {
	ID      uint   `json:"id"`
	IDUser  int    `json:"id_user"`
	IDPost  uint   `json:"id_post"`
	Content string `json:"content_comment"`
	User    User   `json:"user"`
	Post    Post   `json:"post"`
}
