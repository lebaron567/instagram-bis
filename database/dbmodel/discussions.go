package dbmodel

import (
	"errors"

	"gorm.io/gorm"
)

type Discussion struct {
	ID       uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string    `gorm:"size:255;not null" json:"name_discussion"`
	Messages []Message `gorm:"foreignKey:DiscussionID" json:"messages"`
}

type DiscussionRepository interface {
	Create(discussion *Discussion) (*Discussion, error)
	FindByID(id uint) (*Discussion, error)
	FindByUserID(userID uint) ([]*Discussion, error)
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

func (r *discussionRepository) FindByUserID(userID uint) ([]*Discussion, error) {
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

func (r *discussionRepository) FindByID(id uint) (*Discussion, error) {
	var discussion Discussion
	if err := r.db.First(&discussion, id).Error; err != nil {
		return nil, err
	}
	return &discussion, nil
}
