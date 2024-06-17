package usecase

import (
	"testing"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/test"
	"github.com/stretchr/testify/require"
)

func makeListDeliveredPackagesSut() *ListDeliveredPackagesUsecase {
	packageRepository := test.NewInMemoryPackageRepository()
	deliverymanId := "fake-deliveryman-id"

	deliveredPkg1 := entity.NewPackage(
		"fake-recipient-id",
		"Package 1",
	)
	deliveredPkg1.Status = "DELIVERED"
	deliveredPkg1.DeliverymanId = &deliverymanId
	packageRepository.Create(deliveredPkg1)

	deliveredPkg2 := entity.NewPackage(
		"fake-recipient-id",
		"Package 2",
	)
	deliveredPkg2.Status = "DELIVERED"
	deliveredPkg2.DeliverymanId = &deliverymanId
	packageRepository.Create(deliveredPkg2)

	packageRepository.Create(entity.NewPackage(
		"fake-recipient-id",
		"Package 3",
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
