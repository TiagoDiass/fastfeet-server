package usecase

import (
	"testing"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/repository"
	"github.com/TiagoDiass/fastfeet-server/internal/test"
	"github.com/stretchr/testify/require"
)

type MakeWithdrawPackageSutOutput struct {
	withdrawPackageUsecase *WithdrawPackageUsecase
	packageRepository      repository.PackageRepository
}

func makeWithdrawPackageSut() *MakeWithdrawPackageSutOutput {
	deliveryman, _ := entity.NewUser(
		"fake-document",
		"fake-password",
		"fake-name",
		"email@example.com",
		"fake-phone",
		"deliveryman",
	)
	deliveryman.ID = "deliveryman-id"

	pkg := entity.NewPackage(
		"fake-recipient-id",
		"fake-deliveryman-id",
		"Package from Amazon to John Doe",
		"WAITING_WITHDRAW",
	)
	pkg.ID = "package-id"

	packageRepository := test.NewInMemoryPackageRepository()
	userRepository := test.NewInMemoryUserRepository()

	userRepository.Create(deliveryman)
	packageRepository.Create(pkg)

	withdrawPackageUsecase := NewWithdrawPackageUsecase(
		packageRepository,
		userRepository,
	)

	return &MakeWithdrawPackageSutOutput{
		withdrawPackageUsecase: withdrawPackageUsecase,
		packageRepository:      packageRepository,
	}
}

func TestWithdrawPackageSuccessCase(t *testing.T) {
	sut := makeWithdrawPackageSut()

	input := WithdrawPackageInputDTO{
		PackageID:     "package-id",
		DeliverymanID: "deliveryman-id",
	}

	pkg, err := sut.withdrawPackageUsecase.Execute(input)

	require.Nil(t, err)
	require.NotNil(t, pkg)
	require.Equal(t, pkg.Status, "ON_GOING")
	require.NotEmpty(t, pkg.WithdrewAt)

	pkgFromDb, _ := sut.packageRepository.FindById("package-id")

	require.Equal(t, pkgFromDb.Status, "ON_GOING")
	require.NotEmpty(t, pkgFromDb.WithdrewAt)
}