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
		"Package from Amazon to John Doe",
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

func TestWithdrawPackageWhenPackageNotExists(t *testing.T) {
	sut := makeWithdrawPackageSut()

	input := WithdrawPackageInputDTO{
		PackageID:     "non-existent-package-id",
		DeliverymanID: "deliveryman-id",
	}

	pkg, err := sut.withdrawPackageUsecase.Execute(input)

	require.Nil(t, pkg)
	require.NotNil(t, err)
	require.ErrorIs(t, err, ErrPackageNotExists)
}

func TestWithdrawPackageWhenPackageWasAlreadyWithdrew(t *testing.T) {
	sut := makeWithdrawPackageSut()

	pkg := entity.NewPackage(
		"fake-recipient-id",
		"Package from Amazon to John Doe",
	)
	fakeDeliverymanId := "fake-deliveryman-id"
	pkg.ID = "already-withdrew-package-id"
	pkg.Status = "ON_GOING"
	pkg.DeliverymanId = &fakeDeliverymanId

	err := sut.packageRepository.Create(pkg)
	require.Nil(t, err)

	input := WithdrawPackageInputDTO{
		PackageID:     "already-withdrew-package-id",
		DeliverymanID: "deliveryman-id",
	}

	pkg, err = sut.withdrawPackageUsecase.Execute(input)

	require.Nil(t, pkg)
	require.NotNil(t, err)
	require.ErrorIs(t, err, ErrPackageWasAlreadyWithdrew)
}

func TestWithdrawPackageWhenDeliverymanNotExists(t *testing.T) {
	sut := makeWithdrawPackageSut()

	input := WithdrawPackageInputDTO{
		PackageID:     "package-id",
		DeliverymanID: "non-existent-deliveryman-id",
	}

	pkg, err := sut.withdrawPackageUsecase.Execute(input)

	require.Nil(t, pkg)
	require.NotNil(t, err)
	require.ErrorIs(t, err, ErrDeliverymanNotExists)
}
