package model

type Discussion struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name_discussion"`
	IDMembers int      `json:"id_members"`
	Members   []Member `json:"members"`
}
