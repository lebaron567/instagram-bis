package dbmodel

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	LastName       string     `json:"lastename_user"`
	FirstName      string     `json:"firstname_user"`
	Email          string     `json:"email_user"`
	Pseudo         string     `json:"pseudo_user"`
	Birthdate      string     `json:"birthdate"`
	Password       string     `json:"password_user"`
	IsPrivate      bool       `json:"isprivate_user"`
	ProfilePicture string     `json:"profilpicture_user"`
	WantsNotify    bool       `json:"wantsnotify_user"`
	Followers      []Follower `gorm:"foreignKey:IDUser;references:ID"`
	Posts          []Post     `gorm:"foreignKey:IDUser;references:ID"`
	Likes          []Like     `gorm:"foreignKey:IDUser;references:ID"`
	Comments       []Comment  `gorm:"foreignKey:IDUser;references:ID"`
	Messages       []Message  `gorm:"foreignKey:IDUser;references:ID"`
}

type UserRepository interface {
	Create(user *User) (*User, error)
	FindAll() ([]*User, error)
	FindByID(id int) (*User, error)
	UpdateUser(id int, updatedUser *User) (*User, error)
	Delete(userID int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *User) (*User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindAll() ([]*User, error) {
	var users []*User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) UpdateUser(id int, updatedUser *User) (*User, error) {
	// Fetch the existing user
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Update user fields
	user.LastName = updatedUser.LastName
	user.LastName = updatedUser.LastName
	user.Email = updatedUser.Email
	user.Pseudo = updatedUser.Pseudo
	user.Birthdate = updatedUser.Birthdate
	user.IsPrivate = updatedUser.IsPrivate
	user.ProfilePicture = updatedUser.ProfilePicture
	user.WantsNotify = updatedUser.WantsNotify

	// Only update the password if provided
	if updatedUser.Password != "" {
		user.Password = updatedUser.Password
	}

	// Save changes
	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByID(id int) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Delete(userID int) error {
	var user User
	// Trouver l'utilisateur par ID
	if err := r.db.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	// Supprimer l'utilisateur (cela supprimera automatiquement les relations si les clés étrangères sont en cascade)
	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}