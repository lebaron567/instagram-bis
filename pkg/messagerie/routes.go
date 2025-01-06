package messagerie

import (
	"instagram-bis/database/dbmodel"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterMessageRoutes(router *gin.Engine, db *gorm.DB) {
	repo := dbmodel.NewMessageRepository(db)
	controller := NewMessageController(repo)

	messageRoutes := router.Group("/messages")
	{
		messageRoutes.POST("/", controller.CreateMessage)                                     // Créer un message
		messageRoutes.GET("/discussion/:discussion_id", controller.GetMessagesByDiscussionID) // Récupérer les messages d'une discussion
		messageRoutes.DELETE("/:id", controller.DeleteMessage)                                // Supprimer un message
		messageRoutes.PUT("/:id", controller.UpdateMessage)                                   // Mettre à jour un message
	}
}
