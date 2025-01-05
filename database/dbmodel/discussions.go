package dbmodel

import (
	"errors"

	"gorm.io/gorm"
)

type Discussion struct {
	gorm.Model
	Name      string   `json:"name_discussion"`
	IDMembers int      `json:"id_members"`
	Members   []Member `gorm:"foreignKey:IDDiscussion;references:ID;constraint:OnDelete:CASCADE;"`
}

type DiscussionRepository interface {
	Create(discussion *Discussion) (*Discussion, error)
	FindByID(id int) (*Discussion, error)
	FindByUserID(userID int) ([]*Discussion, error)
}

type discussionRepository struct {
	db *gorm.DB
}

func NewDiscussionRepository(db *gorm.DB) DiscussionRepository {
	return &discussionRepository{db: db}
}

func (r *discussionRepository) Create(discussion *Discussion) (*Discussion, error) {
	if err := r.db.Create(discussion).Error; err != nil {
		return nil, err
	}
	return discussion, nil
}

func (r *discussionRepository) FindByUserID(userID int) ([]*Discussion, error) {
	var discussions []*Discussion
	if err := r.db.
		Joins("JOIN members ON members.discussion_id = discussions.id").
		Where("members.user_id = ?", userID).
		Find(&discussions).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no discussions found for the user")
		}
		return nil, err
	}
	return discussions, nil
}

func (r *discussionRepository) FindByID(id int) (*Discussion, error) {
	var discussion Discussion
	if err := r.db.First(&discussion, id).Error; err != nil {
		return nil, err
	}
	return &discussion, nil
}
