package usecase

import (
	"testing"

	"github.com/TiagoDiass/fastfeet-server/internal/test"
	"github.com/stretchr/testify/require"
)

func makeRecipientSut() *CreateRecipientUsecase {
	recipientRepository := test.NewInMemoryRecipientRepository()
	createRecipientUsecase := NewCreateRecipientUsecase(recipientRepository)
	return createRecipientUsecase
}

func TestCreateRecipientSuccessCase(t *testing.T) {
	createRecipientUsecase := makeRecipientSut()

	input := CreateRecipientInputDTO{
		Name:        "John Doe",
		Email:       "john@example.com",
		Street:      "Main St",
		HouseNumber: "123",
		District:    "Downtown",
		State:       "CA",
		Zipcode:     "12345",
	}

	output, err := createRecipientUsecase.Execute(input)

	require.Nil(t, err)
	require.NotNil(t, output)
	require.Equal(t, output.Name, input.Name)
	require.Equal(t, output.Email, input.Email)
	require.Equal(t, output.Street, input.Street)
	require.Equal(t, output.HouseNumber, input.HouseNumber)
	require.Equal(t, output.District, input.District)
	require.Equal(t, output.State, input.State)
	require.Equal(t, output.Zipcode, input.Zipcode)
}

func TestCreateRecipientWhenRepositoryReturnsAnError(t *testing.T) {
	createRecipientUsecase := makeRecipientSut()

	input := CreateRecipientInputDTO{
		Name:        "John Doe",
		Email:       test.EmailThatReturnsErrorOnCreateRecipient,
		Street:      "Main St",
		HouseNumber: "123",
		District:    "Downtown",
		State:       "CA",
		Zipcode:     "12345",
	}

	output, err := createRecipientUsecase.Execute(input)

	require.Nil(t, output)
	require.NotNil(t, err)
	require.ErrorIs(t, err, test.MockErrorOnCreateRecipient)
}
