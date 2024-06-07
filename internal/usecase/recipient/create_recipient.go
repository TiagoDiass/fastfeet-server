package usecase

import (
	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/repository"
)

type CreateRecipientInputDTO struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	District    string `json:"district"`
	State       string `json:"state"`
	Zipcode     string `json:"zipcode"`
}

type CreateRecipientOutputDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	District    string `json:"district"`
	State       string `json:"state"`
	Zipcode     string `json:"zipcode"`
}

type CreateRecipientUsecase struct {
	RecipientRepository repository.RecipientRepository
}

func NewCreateRecipientUsecase(recipientRepository repository.RecipientRepository) *CreateRecipientUsecase {
	return &CreateRecipientUsecase{
		RecipientRepository: recipientRepository,
	}
}

func (u *CreateRecipientUsecase) Execute(input CreateRecipientInputDTO) (*CreateRecipientOutputDTO, error) {
	address := entity.NewAddress(
		input.Street,
		input.HouseNumber,
		input.District,
		input.State,
		input.Zipcode,
	)

	recipient := entity.NewRecipient(input.Name, input.Email, address)

	err := u.RecipientRepository.Create(recipient)

	if err != nil {
		return nil, err
	}

	output := &CreateRecipientOutputDTO{
		ID:          recipient.ID,
		Name:        recipient.Name,
		Email:       recipient.Email,
		Street:      recipient.Street,
		HouseNumber: recipient.HouseNumber,
		District:    recipient.District,
		State:       recipient.State,
		Zipcode:     recipient.Zipcode,
	}

	return output, nil
}
