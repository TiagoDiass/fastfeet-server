package usecase

import (
	"errors"
	"time"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/repository"
)

type CreatePackageInputDTO struct {
	UserID        string
	RecipientID   string `json:"recipient_id"`
	DeliverymanID string `json:"deliveryman_id"`
	Name          string `json:"name"`
}

type CreatePackageOutputDTO struct {
	ID               string     `json:"id"`
	RecipientId      string     `json:"recipient_id"`
	DeliverymanId    string     `json:"deliveryman_id"`
	Name             string     `json:"name"`
	Status           string     `json:"status"` // WAITING_WITHDRAW
	PostedAt         time.Time  `json:"posted_at"`
	DeliveredPicture *string    `json:"delivered_picture"`
	WithdrewAt       *time.Time `json:"withdrew_at"`
	DeliveredAt      *time.Time `json:"delivered_at"`
}

type CreatePackageUsecase struct {
	PackageRepository   repository.PackageRepository
	UserRepository      repository.UserRepository
	RecipientRepository repository.RecipientRepository
}

func NewCreatePackageUsecase(
	packageRepository repository.PackageRepository,
	userRepository repository.UserRepository,
	recipientRepository repository.RecipientRepository,
) *CreatePackageUsecase {
	return &CreatePackageUsecase{
		PackageRepository:   packageRepository,
		UserRepository:      userRepository,
		RecipientRepository: recipientRepository,
	}
}

var (
	ErrUserIsNotAdmin       = errors.New("unauthorized: only admins can create packages")
	ErrRecipientNotExists   = errors.New("recipient does not exist")
	ErrDeliverymanNotExists = errors.New("deliveryman does not exist")
)

func (u *CreatePackageUsecase) Execute(input CreatePackageInputDTO) (*CreatePackageOutputDTO, error) {
	user, err := u.UserRepository.FindById(input.UserID)

	if err != nil || user.Role != "admin" {
		return nil, ErrUserIsNotAdmin
	}

	_, err = u.RecipientRepository.FindById(input.RecipientID)

	if err != nil {
		return nil, ErrRecipientNotExists
	}

	deliveryman, err := u.UserRepository.FindById(input.DeliverymanID)

	if err != nil || deliveryman.Role != "deliveryman" {
		return nil, ErrDeliverymanNotExists
	}

	pkg := entity.NewPackage(
		input.RecipientID,
		input.DeliverymanID,
		input.Name,
		"WAITING_WITHDRAW",
	)

	err = u.PackageRepository.Create(pkg)

	if err != nil {
		return nil, err
	}

	output := &CreatePackageOutputDTO{
		ID:               pkg.ID,
		RecipientId:      pkg.RecipientId,
		DeliverymanId:    pkg.DeliverymanId,
		Name:             pkg.Name,
		Status:           pkg.Status,
		PostedAt:         pkg.PostedAt,
		DeliveredPicture: pkg.DeliveredPicture,
		WithdrewAt:       pkg.WithdrewAt,
		DeliveredAt:      pkg.DeliveredAt,
	}

	return output, nil
}
