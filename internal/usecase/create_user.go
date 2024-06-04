package usecase

import (
	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/repository"
)

type CreateUserInputDTO struct {
	Document string `json:"document"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"` // "admin" or "deliveryman"
}

type CreateUserOutputDTO struct {
	ID       string `json:"id"`
	Document string `json:"document"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type CreateUserUsecase struct {
	userRepository repository.UserRepository
}

func NewCreateUserUsecase(userRepository repository.UserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{
		userRepository: userRepository,
	}
}

func (u *CreateUserUsecase) Execute(input CreateUserInputDTO) (*CreateUserOutputDTO, error) {
	user, err := entity.NewUser(
		input.Document,
		input.Password,
		input.Name,
		input.Email,
		input.Phone,
		input.Role,
	)

	if err != nil {
		return nil, err
	}

	err = u.userRepository.Create(user)

	if err != nil {
		return nil, err
	}

	return &CreateUserOutputDTO{
		ID:       user.ID,
		Document: user.Document,
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Role:     user.Role,
	}, nil
}
