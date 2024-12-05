package dbmodel

import (
	"gorm.io/gorm"
)

type Message struct {
	ID           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       uint   `gorm:"not null" json:"id_user"`
	DiscussionID uint   `gorm:"not null" json:"id_discussion"`
	Content      string `gorm:"type:text;not null" json:"content_message"`
}

type MessageRepository interface {
	Create(message *Message) (*Message, error)
	FindByDiscussionID(discussionID uint) ([]*Message, error)
	Delete(id uint) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) Create(message *Message) (*Message, error) {
	if err := r.db.Create(message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (r *messageRepository) FindByDiscussionID(discussionID uint) ([]*Message, error) {
	var messages []*Message
	if err := r.db.Where("discussion_id = ?", discussionID).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *messageRepository) Delete(id uint) error {
	return r.db.Delete(&Message{}, id).Error
}
