package usecase

import (
	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/repository"
)

type ConfirmDeliveredPackageInputDTO struct {
	PackageID     string `json:"package_id"`
	DeliverymanID string `json:"deliveryman_id"`
}

type ConfirmDeliveredPackageUsecase struct {
	PackageRepository repository.PackageRepository
	UserRepository    repository.UserRepository
}

func NewConfirmDeliveredPackageUsecase(
	packageRepository repository.PackageRepository,
	userRepository repository.UserRepository,
) *ConfirmDeliveredPackageUsecase {
	return &ConfirmDeliveredPackageUsecase{
		PackageRepository: packageRepository,
		UserRepository:    userRepository,
	}
}

func (u *ConfirmDeliveredPackageUsecase) Execute(input ConfirmDeliveredPackageInputDTO) (*entity.Package, error) {
	pkg, err := u.PackageRepository.FindById(input.PackageID)

	if err != nil {
		return nil, ErrPackageNotExists
	}

	if pkg.Status != "WAITING_WITHDRAW" {
		return nil, ErrPackageWasAlreadyWithdrew
	}

	_, err = u.UserRepository.FindDeliverymanById(input.DeliverymanID)

	if err != nil {
		return nil, ErrDeliverymanNotExists
	}

	// pkg = pkg.Withdraw()

	err = u.PackageRepository.Update(pkg)

	if err != nil {
		return nil, err
	}

	return pkg, nil
}
