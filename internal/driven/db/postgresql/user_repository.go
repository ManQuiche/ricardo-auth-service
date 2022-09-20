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
	err := u.client.First(&gUser, "user_id", userID).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return gUser, nil
}

func (u userRepository) Update(user entities.User) (*entities.User, error) {
	var updUser *entities.User
	*updUser = user
	err := u.client.Save(&updUser).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return updUser, nil
}

func (u userRepository) Delete(userID uint) (*entities.User, error) {
	var delUser *entities.User
	err := u.client.Delete(&delUser, "user_id", userID).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return delUser, nil
}
