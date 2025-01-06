package dbmodel

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	IDUser       int        `json:"id_user"`
	IDDiscussion int        `json:"id_discussion"`
	Content      string     `json:"content_message"`
	User         User       `gorm:"foreignKey:IDUser;references:ID;constraint:OnDelete:CASCADE;"`
	Discussion   Discussion `gorm:"foreignKey:IDDiscussion;references:ID;constraint:OnDelete:CASCADE;"`
}

type MessageRepository interface {
	Create(message *Message) (*Message, error)
	FindByDiscussionID(discussionID int) ([]*Message, error)
	Delete(id int) error
	Update(id int, updatedMessage *Message) (*Message, error) // Nouvelle méthode pour la mise à jour
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

func (r *messageRepository) FindByDiscussionID(discussionID int) ([]*Message, error) {
	var messages []*Message
	if err := r.db.Where("discussion_id = ?", discussionID).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *messageRepository) Update(id int, updatedMessage *Message) (*Message, error) {
	// Rechercher le message existant
	var existingMessage Message
	if err := r.db.First(&existingMessage, id).Error; err != nil {
		return nil, err // Retourne une erreur si le message n'existe pas
	}

	// Mettre à jour les champs du message
	existingMessage.Content = updatedMessage.Content

	// Sauvegarder les modifications
	if err := r.db.Save(&existingMessage).Error; err != nil {
		return nil, err
	}

	return &existingMessage, nil
}

func (r *messageRepository) Delete(id int) error {
	return r.db.Delete(&Message{}, id).Error
}
