package repositories

import (
	"github.com/nvanonim/fiber-emr/app/models"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

// UserRepository is the repository for the user model
type UserRepository struct {
	db *gorm.DB
}

// FindByUsername finds the user by username and password
func (ur UserRepository) FindByUsername(username string) (models.User, error) {
	var user models.User
	if err := ur.db.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Create creates a new user
func (ur UserRepository) Create(user *models.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
