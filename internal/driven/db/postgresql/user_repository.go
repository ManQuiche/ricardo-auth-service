package postgresql

import (
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	"gitlab.com/ricardo134/auth-service/internal/core/ports/user"
	"gorm.io/gorm"
)

type userRepository struct {
	client *gorm.DB
}

func NewUserRepository(client *gorm.DB) user.Repository {
	return userRepository{
		client: client,
	}
}

func (u userRepository) Get(userID uint) (*entities.User, error) {
	var gUser *entities.User
	err := u.client.First(&gUser, userID).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return gUser, nil
}

func (u userRepository) Update(user entities.User) (*entities.User, error) {
	err := u.client.Save(&user).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return &user, nil
}

func (u userRepository) Delete(userID uint) (*entities.User, error) {
	var delUser *entities.User
	err := u.client.Delete(&delUser, userID).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return delUser, nil
}
