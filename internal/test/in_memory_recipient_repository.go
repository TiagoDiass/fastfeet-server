package test

import (
	"errors"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
)

var (
	ErrRecipientDoesNotExist = errors.New("recipient does not exist")
)

type InMemoryRecipientRepository struct {
	recipients map[string]*entity.Recipient
}

func NewInMemoryRecipientRepository() *InMemoryRecipientRepository {
	return &InMemoryRecipientRepository{
		recipients: make(map[string]*entity.Recipient),
	}
}

func (r *InMemoryRecipientRepository) Create(recipientId *entity.Recipient) error {
	r.recipients[recipientId.ID] = recipientId

	return nil
}

func (r *InMemoryRecipientRepository) FindById(recipientId string) (*entity.Recipient, error) {
	pkg, exists := r.recipients[recipientId]

	if !exists {
		return nil, ErrRecipientDoesNotExist
	}

	return pkg, nil
}

func (r *InMemoryRecipientRepository) FindAll() ([]*entity.Recipient, error) {
	var recipients []*entity.Recipient

	for _, product := range r.recipients {
		recipients = append(recipients, product)
	}

	return recipients, nil
}
