package usecase

import (
	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/repository"
)

type ConfirmDeliveredPackageInputDTO struct {
	PackageID           string
	DeliverymanID       string
	DeliveredPictureURL string `json:"delivered_picture_url"`
}

type ConfirmDeliveredPackageUsecase struct {
	PackageRepository repository.PackageRepository
}

func NewConfirmDeliveredPackageUsecase(
	packageRepository repository.PackageRepository,
) *ConfirmDeliveredPackageUsecase {
	return &ConfirmDeliveredPackageUsecase{
		PackageRepository: packageRepository,
	}
}

func (u *ConfirmDeliveredPackageUsecase) Execute(input ConfirmDeliveredPackageInputDTO) (*entity.Package, error) {
	pkg, err := u.PackageRepository.FindById(input.PackageID)

	if err != nil {
		return nil, ErrPackageNotExists
	}

	if pkg.Status != "ON_GOING" {
		return nil, ErrPackageCannotBeDelivered
	}

	if *pkg.DeliverymanId != input.DeliverymanID {
		return nil, ErrDifferentDeliveryman
	}

	pkg = pkg.MarkAsDelivered(input.DeliveredPictureURL)

	err = u.PackageRepository.Update(pkg)

	if err != nil {
		return nil, err
	}

	return pkg, nil
}
