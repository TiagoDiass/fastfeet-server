package usecase

import (
	"testing"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/test"
	"github.com/stretchr/testify/require"
)

func makeCreatePackageSut() *CreatePackageUsecase {
	userRepository := test.NewInMemoryUserRepository()
	packageRepository := test.NewInMemoryPackageRepository()
	recipientRepository := test.NewInMemoryRecipientRepository()

	admin, _ := entity.NewUser(
		"42163301001",
		"beautiful-password",
		"Admin",
		"admin@example.com",
		"fake-phone",
		"admin",
	)
	admin.ID = "admin-id"
	userRepository.Create(admin)

	address := entity.NewAddress("Main St", "123", "Downtown", "CA", "12345")
	recipient := entity.NewRecipient(
		"Jane Doe",
		"jane@example.com",
		address,
	)
	recipient.ID = "recipient-id"
	recipientRepository.Create(recipient)

	createPackageUsecase := NewCreatePackageUsecase(
		packageRepository,
		userRepository,
		recipientRepository,
	)

	return createPackageUsecase
}

func TestCreatePackageSuccessCase(t *testing.T) {
	createPackageUsecase := makeCreatePackageSut()

	input := CreatePackageInputDTO{
		UserID:      "admin-id",
		RecipientID: "recipient-id",
		Name:        "Sample Package",
	}

	output, err := createPackageUsecase.Execute(input)

	require.Nil(t, err)
	require.NotNil(t, output)
	require.Equal(t, output.Name, input.Name)
	require.Equal(t, output.RecipientId, input.RecipientID)
	require.Equal(t, output.Status, "WAITING_WITHDRAW")
	require.Nil(t, output.DeliverymanId)
}

func TestCreatePackageUnauthorizedCase(t *testing.T) {
	createPackageUsecase := makeCreatePackageSut()

	input := CreatePackageInputDTO{
		UserID:      "non-existent-admin",
		RecipientID: "recipient-id",
		Name:        "Sample Package",
	}

	output, err := createPackageUsecase.Execute(input)

	require.Nil(t, output)
	require.NotNil(t, err)
	require.ErrorIs(t, err, ErrUserIsNotAdmin)
}

func TestCreatePackageRecipientNotFound(t *testing.T) {
	createPackageUsecase := makeCreatePackageSut()

	input := CreatePackageInputDTO{
		UserID:      "admin-id",
		RecipientID: "non-existent-recipient-id",
		Name:        "Sample Package",
	}

	output, err := createPackageUsecase.Execute(input)

	require.Nil(t, output)
	require.NotNil(t, err)
	require.ErrorIs(t, err, ErrRecipientNotExists)
}
