package messagerie

import (
	"net/http"
	"strconv"

	"instagram-bis/database/dbmodel"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
	repo dbmodel.MessageRepository
}

func NewMessageController(repo dbmodel.MessageRepository) *MessageController {
	return &MessageController{repo: repo}
}

// Créer un message
func (mc *MessageController) CreateMessage(c *gin.Context) {
	var message dbmodel.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMessage, err := mc.repo.Create(&message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create message"})
		return
	}

	c.JSON(http.StatusCreated, createdMessage)
}

// Récupérer les messages d'une discussion
func (mc *MessageController) GetMessagesByDiscussionID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("discussion_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid discussion ID"})
		return
	}

	messages, err := mc.repo.FindByDiscussionID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

// Supprimer un message
func (mc *MessageController) DeleteMessage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message ID"})
		return
	}

	if err := mc.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message deleted successfully"})
}

// Mettre à jour un message
func (mc *MessageController) UpdateMessage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message ID"})
		return
	}

	var updatedMessage dbmodel.Message
	if err := c.ShouldBindJSON(&updatedMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := mc.repo.Update(id, &updatedMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update message"})
		return
	}

	c.JSON(http.StatusOK, message)
}
