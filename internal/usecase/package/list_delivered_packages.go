package usecase

import (
	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/repository"
)

type ListDeliveredPackagesUsecase struct {
	PackageRepository repository.PackageRepository
}

func NewListDeliveredPackagesUsecase(packageRepository repository.PackageRepository) *ListDeliveredPackagesUsecase {
	return &ListDeliveredPackagesUsecase{
		PackageRepository: packageRepository,
	}
}

func (u *ListDeliveredPackagesUsecase) Execute() ([]*entity.Package, error) {
	packages, err := u.PackageRepository.FindAllByStatus("DELIVERED")

	if err != nil {
		return nil, err
	}

	return packages, nil
}
