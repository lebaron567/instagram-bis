package dbmodel

import (
	"gorm.io/gorm"
)

type Post struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint   `gorm:"not null" json:"id_user"`
	Title       string `gorm:"size:50;not null" json:"title_post"`
	Description string `gorm:"type:text;not null" json:"description_post"`
	IsStory     bool   `gorm:"not null" json:"isstory_post"`
}

type PostRepository interface {
	Create(post *Post) (*Post, error)
	FindAll() ([]*Post, error)
	FindByUserID(userID uint) ([]*Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *Post) (*Post, error) {
	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postRepository) FindAll() ([]*Post, error) {
	var posts []*Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) FindByUserID(userID uint) ([]*Post, error) {
	var posts []*Post
	if err := r.db.Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
