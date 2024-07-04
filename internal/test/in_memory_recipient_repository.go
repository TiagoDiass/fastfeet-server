package test

import (
	"errors"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
)

var (
	ErrRecipientNotExists                  = errors.New("recipient does not exist")
	ErrOnCreateRecipient                   = errors.New("mocked error while creating recipient")
	EmailThatReturnsErrorOnCreateRecipient = "error@example.com"
)

type InMemoryRecipientRepository struct {
	recipients map[string]*entity.Recipient
}

func NewInMemoryRecipientRepository() *InMemoryRecipientRepository {
	return &InMemoryRecipientRepository{
		recipients: make(map[string]*entity.Recipient),
	}
}

func (r *InMemoryRecipientRepository) Create(recipient *entity.Recipient) error {
	if recipient.Email == EmailThatReturnsErrorOnCreateRecipient {
		return ErrOnCreateRecipient
	}

	r.recipients[recipient.ID] = recipient

	return nil
}

func (r *InMemoryRecipientRepository) FindById(recipientId string) (*entity.Recipient, error) {
	pkg, exists := r.recipients[recipientId]

	if !exists {
		return nil, ErrRecipientNotExists
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
