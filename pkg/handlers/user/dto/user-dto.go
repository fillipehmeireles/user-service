package dto

import "github.com/fillipehmeireles/user-service/core/domain/user"

type userResponseDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type userRequestDto struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
type requestDto struct {
	userRequestDto
}

func (rDto *requestDto) ToDomain(userDomain *user.User) {
	userDomain.Name = rDto.Name
	userDomain.Email = rDto.Email
	userDomain.PhoneNumber = rDto.PhoneNumber
}

type (
	CreateUserRequestDto struct {
		requestDto
	}
	UpdateUserRequestDto struct {
		requestDto
	}
	GetAllUsersResponseDto struct {
		Users []userResponseDto `json:"users"`
	}
	GetOneResponseDto struct {
		userResponseDto
	}
)

func (g *GetAllUsersResponseDto) FromDomain(usersDomain user.Users) {
	for _, u := range usersDomain {
		g.Users = append(g.Users, userResponseDto{
			Id:          u.ID,
			Name:        u.Name,
			Email:       u.Email,
			PhoneNumber: u.PhoneNumber,
		})
	}
}

func (g *GetOneResponseDto) FromDomain(userDomain user.User) {
	g.Id = userDomain.ID
	g.Name = userDomain.Name
	g.Email = userDomain.Email
	g.PhoneNumber = userDomain.PhoneNumber
}
