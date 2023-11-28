package ports

import "github.com/fillipehmeireles/user-service/core/domain/user"

type UserRepository interface {
	Create(newUser user.User) error
	GetAll() (user.Users, error)
	GetOne(id int) (user.User, error)
	Update(updatedUser user.User) error
	Delete(id int) error
}
