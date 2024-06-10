package usecase

import (
	"time"

	"github.com/TiagoDiass/fastfeet-server/internal/repository"
)

type ListPackagesOutputDTO struct {
	ID               string     `json:"id"`
	RecipientId      string     `json:"recipient_id"`
	DeliverymanId    string     `json:"deliveryman_id"`
	Name             string     `json:"name"`
	Status           string     `json:"status"`
	PostedAt         time.Time  `json:"posted_at"`
	DeliveredPicture *string    `json:"delivered_picture"`
	WithdrewAt       *time.Time `json:"withdrew_at"`
	DeliveredAt      *time.Time `json:"delivered_at"`
}

type ListAvailablePackagesUsecase struct {
	PackageRepository repository.PackageRepository
}

func NewListAvailablePackagesUsecase(packageRepository repository.PackageRepository) *ListAvailablePackagesUsecase {
	return &ListAvailablePackagesUsecase{
		PackageRepository: packageRepository,
	}
}

func (u *ListAvailablePackagesUsecase) Execute() ([]ListPackagesOutputDTO, error) {
	packages, err := u.PackageRepository.FindAllByStatus("WAITING_WITHDRAW")

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
