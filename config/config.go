package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	DB                       *gorm.DB
	authenticationRepository dbmodel.authenticationRepository
	userRepository           dbmodel.userRepository
	postRepository           dbmodel.postRepository
	commentRepository        dbmodel.commentRepository
	likeRepository           dbmodel.likeRepository
	followRepository         dbmodel.followRepository
	notificationRepository   dbmodel.notificationRepository
	messageRepository        dbmodel.messageRepository
	discussionRepository     dbmodel.discussionRepository
	membersRepository        dbmodel.membersRepository
}

func New() (*Config, error) {
	db, err := gorm.Open(sqlite.Open("intagrambis.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Impossible de se connecter à la base de données : %v", err)
		return nil, err
	}
	if err := db.AutoMigrate(&dbmodel.User{}, &dbmodel.Post{}, &dbmodel.Comment{}, &dbmodel.Like{}, &dbmodel.Follow{}, &dbmodel.Notification{}, &dbmodel.Message{}, &dbmodel.Discussion{}, &dbmodel.Members{}); err != nil {
		log.Fatal("Erreur lors de la migration : %v", err)
		return nil, err
	}

	authenticationRepo := dbmodel.NewAuthenticationRepository(db)
	userRepo := dbmodel.NewUserRepository(db)
	postRepo := dbmodel.NewPostRepository(db)
	commentRepo := dbmodel.NewCommentRepository(db)
	likeRepo := dbmodel.NewLikeRepository(db)
	followRepo := dbmodel.NewFollowRepository(db)
	notificationRepo := dbmodel.NewNotificationRepository(db)
	messageRepo := dbmodel.NewMessageRepository(db)
	discussionRepo := dbmodel.NewDiscussionRepository(db)
	membersRepo := dbmodel.NewMembersRepository(db)

	config := Config{
		DB:                       db,
		authenticationRepository: authenticationRepo,
		userRepository:           userRepo,
		postRepository:           postRepo,
		commentRepository:        commentRepo,
		likeRepository:           likeRepo,
		followRepository:         followRepo,
		notificationRepository:   notificationRepo,
		messageRepository:        messageRepo,
		discussionRepository:     discussionRepo,
		membersRepository:        membersRepo,
	}

	return &config, nil
}
