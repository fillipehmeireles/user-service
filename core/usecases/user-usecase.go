package usecases

import (
	"errors"

	"github.com/fillipehmeireles/user-service/core/domain/user"
	"github.com/fillipehmeireles/user-service/core/domain/user/ports"
	"github.com/fillipehmeireles/user-service/pkg/handlers/user/dto"
	"gorm.io/gorm"
)

type UserUseCase struct {
	userRepo ports.UserRepository
}

// Create implements ports.UserUseCase.
func (uUC *UserUseCase) Create(newUser dto.CreateUserRequestDto) error {

	var user user.User

	newUser.ToDomain(&user)

	if err := uUC.userRepo.Create(user); err != nil {
		return err
	}

	return nil
}

// Delete implements ports.UserUseCase.
func (uUC *UserUseCase) Delete(id int) error {
	if err := uUC.userRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

// GetAll implements ports.UserUseCase.
func (uUC *UserUseCase) GetAll() (dto.GetAllUsersResponseDto, error) {
	users, err := uUC.userRepo.GetAll()
	if err != nil {
		return dto.GetAllUsersResponseDto{}, err
	}

	var usrs dto.GetAllUsersResponseDto

	usrs.FromDomain(users)
	return usrs, nil
}

// GetOne implements ports.UserUseCase.
func (uUC *UserUseCase) GetOne(id int) (dto.GetOneResponseDto, error) {
	u, err := uUC.userRepo.GetOne(id)
	if err != nil {
		return dto.GetOneResponseDto{}, err
	}

	var usr dto.GetOneResponseDto

	usr.FromDomain(u)
	return usr, nil
}

// Update implements ports.UserUseCase.
func (uUC *UserUseCase) Update(id int, updatedUser dto.UpdateUserRequestDto) error {

	_, err := uUC.userRepo.GetOne(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("cannot update a inexistent user")
		}

		return err
	}

	var user user.User

	updatedUser.ToDomain(&user)
	user.ID = id
	if err := uUC.userRepo.Update(user); err != nil {
		return err
	}

	return nil
}

func NewUserUseCase(userRepo ports.UserRepository) ports.UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}
