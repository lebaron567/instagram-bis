package dbmodel

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Lastname       string `gorm:"size:30;not null" json:"lastename_user"`
	Firstname      string `gorm:"size:30;not null" json:"firstname_user"`
	Email          string `gorm:"size:70;not null" json:"email_user"`
	Pseudo         string `gorm:"size:30;not null" json:"pseudo_user"`
	Birthdate      string `gorm:"type:date;not null" json:"birthdate"`
	Password       string `gorm:"size:100;not null" json:"password_user"`
	IsPrivate      bool   `gorm:"not null" json:"isprivate_user"`
	ProfilePicture string `gorm:"type:text;not null" json:"profilpicture_user"`
	WantsNotify    bool   `gorm:"not null" json:"wantsnotify_user"`
}

type UserRepository interface {
	Create(user *User) (*User, error)
	FindAll() ([]*User, error)
	FindByID(id uint) (*User, error)
	UpdateUser(id uint, updatedUser *User) (*User, error)
	Delete(userID uint) error
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

func (r *userRepository) UpdateUser(id uint, updatedUser *User) (*User, error) {
	// Fetch the existing user
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Update user fields
	user.Lastname = updatedUser.Lastname
	user.Firstname = updatedUser.Firstname
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

func (r *userRepository) FindByID(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Delete(userID uint) error {
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
