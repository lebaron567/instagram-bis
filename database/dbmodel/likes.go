package dbmodel

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	IDPost    int     `json:"id_post"`
	IDUser    int     `json:"id_user"`
	IDComment int     `json:"id_comment"` //  (peut être NULL)
	Post      Post    `gorm:"foreignKey:IDPost;references:ID;constraint:OnDelete:CASCADE;"`
	User      User    `gorm:"foreignKey:IDUser;references:ID;constraint:OnDelete:CASCADE;"`
	Comment   Comment `gorm:"foreignKey:IDComment;references:ID;constraint:OnDelete:CASCADE;"`
}

type LikeRepository interface {
	Create(like *Like) (*Like, error)
	Delete(postID, userID int) error
	FindByPostID(postID int) ([]*Like, error)
	FindByCommentID(commentID int) ([]*Like, error)
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

func (r *likeRepository) Delete(postID, userID int) error {
	return r.db.Where("post_id = ? AND user_id = ?", postID, userID).Delete(&Like{}).Error
}

func (r *likeRepository) FindByPostID(postID int) ([]*Like, error) {
	var likes []*Like
	if err := r.db.Where("post_id = ?", postID).Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}

func (r *likeRepository) FindByCommentID(commentID int) ([]*Like, error) {
	var likes []*Like
	if err := r.db.Where("comment_id = ?", commentID).Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}
