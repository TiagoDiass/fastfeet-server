package usecase

import (
	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/repository"
)

type ListAvailablePackagesUsecase struct {
	PackageRepository repository.PackageRepository
}

func NewListAvailablePackagesUsecase(packageRepository repository.PackageRepository) *ListAvailablePackagesUsecase {
	return &ListAvailablePackagesUsecase{
		PackageRepository: packageRepository,
	}
}

func (u *ListAvailablePackagesUsecase) Execute() ([]*entity.Package, error) {
	packages, err := u.PackageRepository.FindAllByStatus("WAITING_WITHDRAW")

	if err != nil {
		return nil, err
	}

	return packages, nil
}
