package dbmodel

import (
	"gorm.io/gorm"
)

type Member struct {
	ID           uint `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       uint `gorm:"not null" json:"id_user"`
	DiscussionID uint `gorm:"not null" json:"id_discussion"`
}

type MemberRepository interface {
	Add(member *Member) (*Member, error)
	Remove(userID, discussionID uint) error
	FindByDiscussionID(discussionID uint) ([]*Member, error)
	FindByUserID(userID uint) ([]*Member, error)
}

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &memberRepository{db: db}
}

func (r *memberRepository) Add(member *Member) (*Member, error) {
	if err := r.db.Create(member).Error; err != nil {
		return nil, err
	}
	return member, nil
}

func (r *memberRepository) Remove(userID, discussionID uint) error {
	return r.db.Where("user_id = ? AND discussion_id = ?", userID, discussionID).Delete(&Member{}).Error
}

func (r *memberRepository) FindByDiscussionID(discussionID uint) ([]*Member, error) {
	var members []*Member
	if err := r.db.Where("discussion_id = ?", discussionID).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *memberRepository) FindByUserID(userID uint) ([]*Member, error) {
	var members []*Member
	if err := r.db.Where("user_id = ?", userID).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}
