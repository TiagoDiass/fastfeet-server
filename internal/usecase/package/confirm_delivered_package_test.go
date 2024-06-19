package usecase

import (
	"testing"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/test"
	"github.com/stretchr/testify/require"
)

func makeConfirmDeliveredPackageSut() (*ConfirmDeliveredPackageUsecase, *test.InMemoryPackageRepository) {
	packageRepository := test.NewInMemoryPackageRepository()
	confirmDeliveredPackageUsecase := NewConfirmDeliveredPackageUsecase(packageRepository)

	pkg := entity.NewPackage(
		"recipient-id",
		"package from Amazon to John Doe",
	)
	pkg.ID = "fake-package-id"
	pkg.Withdraw("deliveryman-id")

	packageRepository.Create(pkg)

	return confirmDeliveredPackageUsecase, packageRepository
}

func TestConfirmDeliveredPackageSuccessCase(t *testing.T) {
	confirmDeliveredPackageUsecase, _ := makeConfirmDeliveredPackageSut()

	input := ConfirmDeliveredPackageInputDTO{
		PackageID:           "fake-package-id",
		DeliverymanID:       "deliveryman-id",
		DeliveredPictureURL: "https://example.com/picture.jpg",
	}

	output, err := confirmDeliveredPackageUsecase.Execute(input)

	require.Nil(t, err)
	require.NotNil(t, output)
	require.Equal(t, output.ID, "fake-package-id")
	require.Equal(t, output.Status, "DELIVERED")
	require.Equal(t, *output.DeliveredPicture, "https://example.com/picture.jpg")
	require.NotNil(t, output.DeliveredAt)
}

func TestConfirmDeliveredPackageWhenPackageDoesNotExist(t *testing.T) {
	confirmDeliveredPackageUsecase, _ := makeConfirmDeliveredPackageSut()

	input := ConfirmDeliveredPackageInputDTO{
		PackageID:           "non-existent-package-id",
		DeliverymanID:       "deliveryman-id",
		DeliveredPictureURL: "https://example.com/picture.jpg",
	}

	output, err := confirmDeliveredPackageUsecase.Execute(input)

	require.NotNil(t, err)
	require.Nil(t, output)
	require.Equal(t, err, ErrPackageNotExists)
}

func TestConfirmDeliveredPackageWhenPackageCannotBeDelivered(t *testing.T) {
	confirmDeliveredPackageUsecase, packageRepository := makeConfirmDeliveredPackageSut()

	pkg := entity.NewPackage(
		"recipient-id",
		"package from Amazon to John Doe",
	)
	pkg.ID = "package-that-is-waiting-withdraw"

	// Save package with WAITING_WITHDRAW status
	packageRepository.Create(pkg)

	input := ConfirmDeliveredPackageInputDTO{
		PackageID:           "package-that-is-waiting-withdraw",
		DeliverymanID:       "deliveryman-id",
		DeliveredPictureURL: "https://example.com/picture.jpg",
	}

	output, err := confirmDeliveredPackageUsecase.Execute(input)

	require.NotNil(t, err)
	require.Nil(t, output)
	require.Equal(t, ErrPackageCannotBeDelivered, err)
}

// func TestConfirmDeliveredPackageWhenPackageDeliverymanIsDifferent(t *testing.T) {
// 	usecase, repo := makeConfirmDeliveredPackageSut()

// 	// Add a package with a different deliveryman to the repository
// 	pkg := &entity.Package{
// 		ID:            "pkg1",
// 		RecipientId:   "rec1",
// 		DeliverymanId: "del2",
// 		Name:          "Package 1",
// 		Status:        "ON_GOING",
// 		PostedAt:      time.Now(),
// 	}
// 	repo.Create(pkg)

// 	input := ConfirmDeliveredPackageInputDTO{
// 		PackageID:           "pkg1",
// 		DeliverymanID:       "del1",
// 		DeliveredPictureURL: "http://example.com/picture.jpg",
// 	}

// 	output, err := usecase.Execute(input)

// 	require.NotNil(t, err)
// 	require.Nil(t, output)
// 	require.Equal(t, ErrDifferentDeliveryman, err)
// }
