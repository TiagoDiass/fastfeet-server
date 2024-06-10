package usecase

import (
	"testing"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/test"
	"github.com/stretchr/testify/require"
)

func makeListDeliveredPackagesSut() *ListDeliveredPackagesUsecase {
	packageRepository := test.NewInMemoryPackageRepository()

	packageRepository.Create(entity.NewPackage(
		"fake-recipient-id",
		"fake-deliveryman-id",
		"Package 1",
		"DELIVERED",
	))
	packageRepository.Create(entity.NewPackage(
		"fake-recipient-id",
		"fake-deliveryman-id",
		"Package 2",
		"DELIVERED",
	))
	packageRepository.Create(entity.NewPackage(
		"fake-recipient-id",
		"fake-deliveryman-id",
		"Package 3",
		"WAITING_WITHDRAW",
	))

	listDeliveredPackages := NewListDeliveredPackagesUsecase(packageRepository)

	return listDeliveredPackages
}

func TestListDeliveredPackages(t *testing.T) {
	listDeliveredPackagesUsecase := makeListDeliveredPackagesSut()

	output, err := listDeliveredPackagesUsecase.Execute()

	require.Nil(t, err)
	require.NotNil(t, output)
	require.Len(t, output, 2)

	require.Equal(t, output[0].Name, "Package 1")
	require.Equal(t, output[0].Status, "DELIVERED")

	require.Equal(t, output[1].Name, "Package 2")
	require.Equal(t, output[1].Status, "DELIVERED")
}
