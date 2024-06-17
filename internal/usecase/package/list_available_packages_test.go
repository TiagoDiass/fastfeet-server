package usecase

import (
	"testing"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/test"
	"github.com/stretchr/testify/require"
)

func makeListAvailablePackagesSut() *ListAvailablePackagesUsecase {
	packageRepository := test.NewInMemoryPackageRepository()

	packageRepository.Create(entity.NewPackage(
		"fake-recipient-id",
		"Package 1",
	))

	deliveredPkg := entity.NewPackage(
		"fake-recipient-id",
		"Package 2",
	)
	deliveredPkg.Status = "DELIVERED"
	deliverymanId := "fake-deliveryman-id"
	deliveredPkg.DeliverymanId = &deliverymanId

	packageRepository.Create(deliveredPkg)
	packageRepository.Create(entity.NewPackage(
		"fake-recipient-id",
		"Package 3",
	))

	listAvailablePackages := NewListAvailablePackagesUsecase(packageRepository)

	return listAvailablePackages
}

func TestListAvailablePackages(t *testing.T) {
	listAvailablePackagesUsecase := makeListAvailablePackagesSut()

	output, err := listAvailablePackagesUsecase.Execute()

	require.Nil(t, err)
	require.NotNil(t, output)
	require.Len(t, output, 2)

	require.Equal(t, output[0].Name, "Package 1")
	require.Equal(t, output[0].Status, "WAITING_WITHDRAW")

	require.Equal(t, output[1].Name, "Package 3")
	require.Equal(t, output[1].Status, "WAITING_WITHDRAW")
}
