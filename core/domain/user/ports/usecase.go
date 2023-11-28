package ports

import "github.com/fillipehmeireles/user-service/pkg/handlers/user/dto"

type UserUseCase interface {
	Create(newUser dto.CreateUserRequestDto) error
	GetAll() (dto.GetAllUsersResponseDto, error)
	GetOne(id int) (dto.GetOneResponseDto, error)
	Update(id int, updatedUser dto.UpdateUserRequestDto) error
	Delete(id int) error
}
