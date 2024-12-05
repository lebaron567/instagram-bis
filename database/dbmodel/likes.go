package dbmodel

import (
	"gorm.io/gorm"
)

type Like struct {
	PostID    uint  `gorm:"not null" json:"id_post"`
	UserID    uint  `gorm:"not null" json:"id_user"`
	CommentID *uint `json:"id_comment"` // Nullable foreign key
}

type LikeRepository interface {
	Create(like *Like) (*Like, error)
	Delete(postID, userID uint) error
	FindByPostID(postID uint) ([]*Like, error)
	FindByCommentID(commentID uint) ([]*Like, error)
}

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) LikeRepository {
	return &likeRepository{db: db}
}

func (r *likeRepository) Create(like *Like) (*Like, error) {
	if err := r.db.Create(like).Error; err != nil {
		return nil, err
	}
	return like, nil
}

func (r *likeRepository) Delete(postID, userID uint) error {
	return r.db.Where("post_id = ? AND user_id = ?", postID, userID).Delete(&Like{}).Error
}

func (r *likeRepository) FindByPostID(postID uint) ([]*Like, error) {
	var likes []*Like
	if err := r.db.Where("post_id = ?", postID).Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}

func (r *likeRepository) FindByCommentID(commentID uint) ([]*Like, error) {
	var likes []*Like
	if err := r.db.Where("comment_id = ?", commentID).Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}
