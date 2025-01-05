package config

import (
	"log"

	"instagram-bis/database/dbmodel"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	DB                   *gorm.DB
	UserRepository       dbmodel.UserRepository
	PostRepository       dbmodel.PostRepository
	CommentRepository    dbmodel.CommentRepository
	LikeRepository       dbmodel.LikeRepository
	FollowerRepository   dbmodel.FollowerRepository
	MessageRepository    dbmodel.MessageRepository
	DiscussionRepository dbmodel.DiscussionRepository
	MemberRepository     dbmodel.MemberRepository
}

func New() (*Config, error) {
	db, err := gorm.Open(sqlite.Open("intagrambis.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Impossible de se connecter à la base de données : %v", err)
		return nil, err
	}
	if err := db.AutoMigrate(&dbmodel.User{}, &dbmodel.Post{}, &dbmodel.Comment{}, &dbmodel.Like{}, &dbmodel.Follower{}, &dbmodel.Message{}, &dbmodel.Discussion{}, &dbmodel.Member{}); err != nil {
		log.Fatalf("Erreur lors de la migration : %v", err)
		return nil, err
	}

	userRepo := dbmodel.NewUserRepository(db)
	postRepo := dbmodel.NewPostRepository(db)
	commentRepo := dbmodel.NewCommentRepository(db)
	likeRepo := dbmodel.NewLikeRepository(db)
	followRepo := dbmodel.NewFollowerRepository(db)
	messageRepo := dbmodel.NewMessageRepository(db)
	discussionRepo := dbmodel.NewDiscussionRepository(db)
	memberRepo := dbmodel.NewMemberRepository(db)

	config := Config{
		DB:                   db,
		UserRepository:       userRepo,
		PostRepository:       postRepo,
		CommentRepository:    commentRepo,
		LikeRepository:       likeRepo,
		FollowerRepository:   followRepo,
		MessageRepository:    messageRepo,
		DiscussionRepository: discussionRepo,
		MemberRepository:     memberRepo,
	}

	return &config, nil
}

func InitDB() {
	var err error
	db, err := gorm.Open(sqlite.Open("intagrambis.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Impossible de se connecter à la base de données : %v", err)
	}
	if err := db.AutoMigrate(&dbmodel.User{}, &dbmodel.Post{}, &dbmodel.Comment{}, &dbmodel.Like{}, &dbmodel.Follower{}, &dbmodel.Message{}, &dbmodel.Discussion{}, &dbmodel.Member{}); err != nil {
		log.Fatalf("Erreur lors de la migration : %v", err)
	}
}
