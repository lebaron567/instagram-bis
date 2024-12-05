package dbmodel

import (
	"errors"

	"gorm.io/gorm"
)

type Comment struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID  uint   `gorm:"not null" json:"id_user"`
	Content string `gorm:"not null" json:"content_comment"`
}

type CommentRepository interface {
	Create(comment *Comment) (*Comment, error)
	Delete(commentID uint) error
	FindByPostID(postID uint) ([]*Comment, error)
	FindAllByUserID(userID uint) ([]*Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment *Comment) (*Comment, error) {
	if err := r.db.Create(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *commentRepository) Delete(commentID uint) error {
	// Find the comment to ensure it exists
	var comment Comment
	if err := r.db.First(&comment, commentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("comment not found")
		}
		return err
	}

	// Delete the comment
	if err := r.db.Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}

func (r *commentRepository) FindByPostID(postID uint) ([]*Comment, error) {
	var comments []*Comment
	if err := r.db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentRepository) FindAllByUserID(userID uint) ([]*Comment, error) {
	var comments []*Comment
	if err := r.db.Where("user_id = ?", userID).Find(&comments).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no comments found for the user")
		}
		return nil, err
	}
	return comments, nil
}
