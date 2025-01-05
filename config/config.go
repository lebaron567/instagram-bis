package config

import (
	"log"

	"instagram-bis/database/dbmodel"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	DB                   *gorm.DB
	userRepository       dbmodel.UserRepository
	postRepository       dbmodel.PostRepository
	commentRepository    dbmodel.CommentRepository
	likeRepository       dbmodel.LikeRepository
	followerRepository   dbmodel.FollowerRepository
	messageRepository    dbmodel.MessageRepository
	discussionRepository dbmodel.DiscussionRepository
	memberRepository     dbmodel.MemberRepository
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
		userRepository:       userRepo,
		postRepository:       postRepo,
		commentRepository:    commentRepo,
		likeRepository:       likeRepo,
		followerRepository:   followRepo,
		messageRepository:    messageRepo,
		discussionRepository: discussionRepo,
		memberRepository:     memberRepo,
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
