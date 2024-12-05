package dbmodel

import (
	"gorm.io/gorm"
)

type Follower struct {
	UserID     uint `gorm:"not null" json:"id_user"`
	FollowerID uint `gorm:"not null" json:"id_follower"`
}

type FollowerRepository interface {
	Follow(follower *Follower) (*Follower, error)
	Unfollow(userID, followerID uint) error
	FindFollowersByUserID(userID uint) ([]*Follower, error)
	FindFollowingByUserID(userID uint) ([]*Follower, error)
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

func (r *followerRepository) Unfollow(userID, followerID uint) error {
	return r.db.Where("user_id = ? AND follower_id = ?", userID, followerID).Delete(&Follower{}).Error
}

func (r *followerRepository) FindFollowersByUserID(userID uint) ([]*Follower, error) {
	var followers []*Follower
	if err := r.db.Where("id_follower = ?", userID).Find(&followers).Error; err != nil {
		return nil, err
	}
	return followers, nil
}

func (r *followerRepository) FindFollowingByUserID(userID uint) ([]*Follower, error) {
	var following []*Follower
	if err := r.db.Where("id_user = ?", userID).Find(&following).Error; err != nil {
		return nil, err
	}
	return following, nil
}
