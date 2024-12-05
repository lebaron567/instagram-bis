package dbmodel

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	IDUser   int       `json:"id_user"`
	Title    string    `json:"title_post"`
	Content  string    `json:"description_post"`
	IsStory  bool      `json:"isstory_post"`
	User     User      `gorm:"foreignKey:IDUser;references:ID;constraint:OnDelete:CASCADE;"`
	Likes    []Like    `gorm:"foreignKey:IDPost;references:ID;constraint:OnDelete:CASCADE;"`
	Comments []Comment `gorm:"foreignKey:IDPost;references:ID;constraint:OnDelete:CASCADE;"`
}

type PostRepository interface {
	Create(post *Post) (*Post, error)
	FindAll() ([]*Post, error)
	FindByUserID(userID int) ([]*Post, error)
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

func (r *postRepository) FindByUserID(userID int) ([]*Post, error) {
	var posts []*Post
	if err := r.db.Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
