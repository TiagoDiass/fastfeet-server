package usecase

import (
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

func (u *ListDeliveredPackagesUsecase) Execute() ([]ListPackagesOutputDTO, error) {
	packages, err := u.PackageRepository.FindAllByStatus("DELIVERED")

	if err != nil {
		return nil, err
	}

	output := []ListPackagesOutputDTO{}

	for _, pkg := range packages {
		output = append(output, ListPackagesOutputDTO{
			ID:               pkg.ID,
			RecipientId:      pkg.RecipientId,
			DeliverymanId:    pkg.DeliverymanId,
			Name:             pkg.Name,
			Status:           pkg.Status,
			PostedAt:         pkg.PostedAt,
			DeliveredPicture: pkg.DeliveredPicture,
			WithdrewAt:       pkg.WithdrewAt,
			DeliveredAt:      pkg.DeliveredAt,
		})
	}

	return output, nil
}
