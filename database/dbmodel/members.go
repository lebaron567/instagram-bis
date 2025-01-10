package dbmodel

import (
	model "instagram-bis/pkg/models"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	IDUser       int        `json:"id_user"`
	IDDiscussion int        `json:"id_discussion"`
	User         User       `gorm:"foreignKey:IDUser;references:ID;constraint:OnDelete:CASCADE;"`
	Discussion   Discussion `gorm:"foreignKey:IDDiscussion;references:ID;constraint:OnDelete:CASCADE;"`
}

type MemberRepository interface {
	Add(member *Member) (*Member, error)
	Remove(userID, discussionID int) error
	FindByDiscussionID(discussionID int) ([]*Member, error)
	FindByUserID(userID int) ([]*Member, error)
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

func (r *memberRepository) Remove(userID, discussionID int) error {
	return r.db.Where("user_id = ? AND discussion_id = ?", userID, discussionID).Delete(&Member{}).Error
}

func (r *memberRepository) FindByDiscussionID(discussionID int) ([]*Member, error) {
	var members []*Member
	if err := r.db.Where("discussion_id = ?", discussionID).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *memberRepository) FindByUserID(userID int) ([]*Member, error) {
	var members []*Member
	if err := r.db.Where("user_id = ?", userID).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (Member *Member) ToModel() model.Member {
	return model.Member{
		ID:           Member.ID,
		IDUser:       Member.IDUser,
		IDDiscussion: Member.IDDiscussion,
	}
}
