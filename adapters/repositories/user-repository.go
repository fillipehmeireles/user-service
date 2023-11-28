package repositories

import (
	"log"

	"github.com/fillipehmeireles/user-service/core/domain/user"
	"github.com/fillipehmeireles/user-service/core/domain/user/ports"
	"gorm.io/gorm"
)

type UserRepository struct {
	dbInstance *gorm.DB
}

// Create implements ports.UserRepository.
func (uRepo *UserRepository) Create(newUser user.User) error {
	if err := uRepo.dbInstance.Create(&newUser).Error; err != nil {
		log.Printf("[UserRepository:Create] Error on creating new user: %s", err)
		return err
	}

	return nil
}

// Delete implements ports.UserRepository.
func (uRepo *UserRepository) Delete(id int) error {
	if err := uRepo.dbInstance.Delete(&user.User{}, id).Error; err != nil {
		log.Printf("[UserRepository:Delete] Error on deleting user %d: %s", id, err)
		return err
	}

	return nil
}

// GetAll implements ports.UserRepository.
func (uRepo *UserRepository) GetAll() (user.Users, error) {
	var users user.Users

	if err := uRepo.dbInstance.Find(&users).Error; err != nil {
		log.Printf("[UserRepository:GetAll] Error on retrieving all users: %s", err)
		return user.Users{}, err
	}

	return users, nil
}

// GetOne implements ports.UserRepository.
func (uRepo *UserRepository) GetOne(id int) (user.User, error) {
	var usr user.User

	if err := uRepo.dbInstance.First(&usr, id).Error; err != nil {
		log.Printf("[UserRepository:GetOne] Error on retrieving one user: %s", err)
		return user.User{}, err
	}

	return usr, nil
}

// Update implements ports.UserRepository.
func (uRepo *UserRepository) Update(updatedUser user.User) error {
	if err := uRepo.dbInstance.Updates(updatedUser).Error; err != nil {
		log.Printf("[UserRepository:Update] Error on updating user %d: %s", updatedUser.ID, err)
		return err
	}

	return nil
}

func NewUserRepository(dbInstance *gorm.DB) ports.UserRepository {
	return &UserRepository{dbInstance}
}
