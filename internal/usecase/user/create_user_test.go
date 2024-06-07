package usecase

import (
	"testing"

	"github.com/TiagoDiass/fastfeet-server/internal/test"
	"github.com/stretchr/testify/require"
)

func TestCreateUserUsecase(t *testing.T) {
	userRepository := test.NewInMemoryUserRepository()
	createUserUsecase := NewCreateUserUsecase(userRepository)

	input := CreateUserInputDTO{
		Document: "87847048027",
		Password: "beautiful-password",
		Name:     "John Doe",
		Email:    "johndoe@gmail.com",
		Phone:    "19912341234",
		Role:     "admin",
	}

	output, err := createUserUsecase.Execute(input)

	require.Nil(t, err)
	require.NotEmpty(t, output.ID)
	require.Equal(t, output.Document, input.Document)
	require.Equal(t, output.Name, input.Name)
	require.Equal(t, output.Email, input.Email)
}
