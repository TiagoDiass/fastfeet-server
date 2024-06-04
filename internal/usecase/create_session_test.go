package usecase

import (
	"testing"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/test"
	"github.com/go-chi/jwtauth"
	"github.com/stretchr/testify/require"
)

func makeSut() *CreateSessionUsecase {
	userRepository := test.NewInMemoryUserRepository()
	jwt := jwtauth.New("HS256", []byte("secret"), nil)
	expiresIn := 600

	user, _ := entity.NewUser(
		"87847048027",
		"beautiful-password",
		"John Doe",
		"john@example.com",
		"fake-phone",
		"admin",
	)
	userRepository.Create(user)

	createSessionUsecase := NewCreateSessionUsecase(userRepository, jwt, expiresIn)

	return createSessionUsecase
}

func TestCreateSessionSuccessCase(t *testing.T) {
	createSessionUsecase := makeSut()

	input := CreateSessionInputDTO{
		Document: "87847048027",
		Password: "beautiful-password",
	}

	output, err := createSessionUsecase.Execute(input)

	require.Nil(t, err)
	require.NotNil(t, output)
	require.NotEmpty(t, output.AccessToken)
}

func TestCreateSessionWhenUserNotExists(t *testing.T) {
	createSessionUsecase := makeSut()

	input := CreateSessionInputDTO{
		Document: "wrong-document",
		Password: "any-password",
	}

	output, err := createSessionUsecase.Execute(input)

	require.Nil(t, output)
	require.NotNil(t, err)
	require.ErrorContains(t, err, "user does not exist")
}

func TestCreateSessionWithWrongPassword(t *testing.T) {
	createSessionUsecase := makeSut()

	input := CreateSessionInputDTO{
		Document: "87847048027",
		Password: "any-password",
	}

	output, err := createSessionUsecase.Execute(input)

	require.Nil(t, output)
	require.NotNil(t, err)
	require.ErrorContains(t, err, "unauthorized")
}
