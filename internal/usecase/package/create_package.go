package usecase

import (
	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/repository"
)

type CreatePackageInputDTO struct {
	UserID      string
	RecipientID string `json:"recipient_id"`
	Name        string `json:"name"`
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

func (u *CreatePackageUsecase) Execute(input CreatePackageInputDTO) (*entity.Package, error) {
	user, err := u.UserRepository.FindById(input.UserID)

	if err != nil || user.Role != "admin" {
		return nil, ErrUserIsNotAdmin
	}

	_, err = u.RecipientRepository.FindById(input.RecipientID)

	if err != nil {
		return nil, ErrRecipientNotExists
	}

	pkg := entity.NewPackage(
		input.RecipientID,
		input.Name,
	)

	err = u.PackageRepository.Create(pkg)

	if err != nil {
		return nil, err
	}

	return pkg, nil
}
