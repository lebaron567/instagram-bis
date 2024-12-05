package dbmodel

import (
	"errors"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	IDUser  int    `json:"id_user"`
	Content string `json:"content_comment"`
	User    User   `gorm:"foreignKey:IDUser;references:ID;constraint:OnDelete:CASCADE;"`
	Post    Post   `gorm:"foreignKey:IDPost;references:ID;constraint:OnDelete:CASCADE;"`
}

type CommentRepository interface {
	Create(comment *Comment) (*Comment, error)
	Delete(commentID int) error
	FindByPostID(postID int) ([]*Comment, error)
	FindAllByUserID(userID int) ([]*Comment, error)
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

func (r *commentRepository) Delete(commentID int) error {
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

func (r *commentRepository) FindByPostID(postID int) ([]*Comment, error) {
	var comments []*Comment
	if err := r.db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentRepository) FindAllByUserID(userID int) ([]*Comment, error) {
	var comments []*Comment
	if err := r.db.Where("user_id = ?", userID).Find(&comments).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no comments found for the user")
		}
		return nil, err
	}
	return comments, nil
}
