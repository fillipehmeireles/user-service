package usecases_test

import (
	"testing"

	"github.com/fillipehmeireles/user-service/core/domain/user"
	"github.com/fillipehmeireles/user-service/core/domain/user/ports"
	"github.com/fillipehmeireles/user-service/core/domain/user/ports/mocks"
	"github.com/fillipehmeireles/user-service/core/usecases"
	"github.com/fillipehmeireles/user-service/pkg/handlers/user/dto"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	userRepoMock             *mocks.UserRepository
	userService              ports.UserUseCase
	mockUserCreateRequestDto dto.CreateUserRequestDto
	mockUserUpdateRequestDto dto.UpdateUserRequestDto
	mockUserDomain           user.User
	mockUsersDomain          user.Users
	userID                   int
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.userRepoMock = &mocks.UserRepository{}
	newDto := dto.CreateUserRequestDto{}
	newDto.Name = "Fillipe"
	newDto.Email = "fillipe.dev@gmail.com"
	newDto.PhoneNumber = "00000000000"
	suite.mockUserCreateRequestDto = newDto

	newUpdateDto := dto.UpdateUserRequestDto{}
	newUpdateDto.Name = "Heung Min Son"
	newUpdateDto.Email = "fillipe.dev@gmail.com"
	newUpdateDto.PhoneNumber = "00000000000"
	suite.mockUserUpdateRequestDto = newUpdateDto
	suite.userID = 1
	suite.mockUserDomain = user.User{
		Name:        "Fillipe",
		Email:       "fillipe.dev@gmail.com",
		PhoneNumber: "00000000000",
	}

	suite.mockUsersDomain = user.Users{
		{
			Name:        "Henrique",
			Email:       "henrique.dev@gmail.com",
			PhoneNumber: "11111111111",
		},
		{
			Name:        "Leal",
			Email:       "leal.dev@gmail.com",
			PhoneNumber: "22222222222",
		},
		{
			Name:        "Meireles",
			Email:       "meireles.dev@gmail.com",
			PhoneNumber: "33333333333",
		},
	}

	suite.userService = usecases.NewUserUseCase(suite.userRepoMock)
}

func (suite *UserUseCaseTestSuite) TestCreate_ShouldThrowAnErrorWhenEntityIsNotCreated() {
	suite.userRepoMock.On("Create", suite.mockUserDomain).Return(nil).Once()

	err := suite.userService.Create(suite.mockUserCreateRequestDto)
	require.Nil(suite.T(), err)

}

func (suite *UserUseCaseTestSuite) TestDelete_ShouldThrowAnErrorWhenEntityIsNotDeleted() {
	suite.userRepoMock.On("Delete", suite.userID).Return(nil).Once()

	err := suite.userService.Delete(suite.userID)
	require.Nil(suite.T(), err)

}

func (suite *UserUseCaseTestSuite) TestGetAll_ShouldThrowAnErrorWhenRepoCannotFetchUsers() {
	suite.userRepoMock.On("GetAll").Return(suite.mockUsersDomain, nil).Once()

	users, err := suite.userService.GetAll()

	require.Nil(suite.T(), err)
	require.Equal(suite.T(), len(users.Users), len(suite.mockUsersDomain))
}

func (suite *UserUseCaseTestSuite) TestGetOne_ShouldThrowAnErrorWhenRepoCannotFetchOneUserByID() {
	suite.userRepoMock.On("GetOne", suite.userID).Return(suite.mockUserDomain, nil).Once()

	user, err := suite.userService.GetOne(suite.userID)

	require.Nil(suite.T(), err)
	user_test := user
	user.FromDomain(suite.mockUserDomain)
	require.Equal(suite.T(), user_test, user)

}

func (suite *UserUseCaseTestSuite) TestUpdate_ShouldThrowAnErrorWhenEntityIsNotUpdated() {
	suite.userRepoMock.On("GetOne", suite.userID).Return(suite.mockUserDomain, nil).Once()
	userToUpdate := suite.mockUserUpdateRequestDto
	userToUpdate.Name = "Heung Min Son"

	var userToUpdateDomain user.User
	userToUpdateDomain.ID = suite.userID
	userToUpdate.ToDomain(&userToUpdateDomain)
	suite.userRepoMock.On("Update", userToUpdateDomain).Return(nil).Once()

	err := suite.userService.Update(suite.userID, suite.mockUserUpdateRequestDto)
	require.Nil(suite.T(), err)
}
func TestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}
