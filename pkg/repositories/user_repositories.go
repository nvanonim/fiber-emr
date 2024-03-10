package repositories

import (
	"github.com/nvanonim/fiber-emr/pkg/configs"
	"github.com/nvanonim/fiber-emr/pkg/models"
)

func NewUserRepository() UserRepository {
	return UserRepository{}
}

// UserRepository is the repository for the user model
type UserRepository struct{}

var db = configs.GetDB()

// FindByUsername finds the user by username and password
func (UserRepository) FindByUsername(username string) (models.User, error) {
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Create creates a new user
func (UserRepository) Create(user *models.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
