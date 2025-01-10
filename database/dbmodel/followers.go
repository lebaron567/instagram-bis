package dbmodel

import (
	model "instagram-bis/pkg/models"

	"gorm.io/gorm"
)

type Follower struct {
	gorm.Model
	IDUser     int  `json:"id_user"`
	IDFollower int  `json:"id_folower"`
	User       User `gorm:"foreignKey:IDUser;references:ID;constraint:OnDelete:CASCADE;"`
	Follower   User `gorm:"foreignKey:IDFollower;references:ID;constraint:OnDelete:CASCADE;"`
}
type FollowerRepository interface {
	Follow(follower *Follower) (*Follower, error)
	Unfollow(userID, followerID int) error
	FindFollowersByUserID(userID int) ([]*Follower, error)
	FindFollowingByUserID(userID int) ([]*Follower, error)
}

type followerRepository struct {
	db *gorm.DB
}

func NewFollowerRepository(db *gorm.DB) FollowerRepository {
	return &followerRepository{db: db}
}

func (r *followerRepository) Follow(follower *Follower) (*Follower, error) {
	if err := r.db.Create(follower).Error; err != nil {
		return nil, err
	}
	return follower, nil
}

func (r *followerRepository) Unfollow(userID, followerID int) error {
	return r.db.Where("user_id = ? AND follower_id = ?", userID, followerID).Delete(&Follower{}).Error
}

func (r *followerRepository) FindFollowersByUserID(userID int) ([]*Follower, error) {
	var followers []*Follower
	if err := r.db.Where("id_follower = ?", userID).Find(&followers).Error; err != nil {
		return nil, err
	}
	return followers, nil
}

func (r *followerRepository) FindFollowingByUserID(userID int) ([]*Follower, error) {
	var following []*Follower
	if err := r.db.Where("id_user = ?", userID).Find(&following).Error; err != nil {
		return nil, err
	}
	return following, nil
}

func (Follower *Follower) ToModel() model.Follower {
	return model.Follower{
		ID:         Follower.ID,
		IDUser:     Follower.IDUser,
		IDFollower: Follower.IDFollower,
	}
}
