package model

import "gorm.io/gorm"

type Discussion struct {
	gorm.Model
	ID        uint     `json:"id"`
	Name      string   `json:"name_discussion"`
	IDMembers int      `json:"id_members"`
	Members   []Member `json:"members"`
}
