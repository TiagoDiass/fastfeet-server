package test

import (
	"errors"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
)

var (
	ErrRecipientDoesNotExist = errors.New("recipient does not exist")
)

type InMemoryRecipientRepository struct {
	packages map[string]*entity.Recipient
}

func NewInMemoryRecipientRepository() *InMemoryRecipientRepository {
	return &InMemoryRecipientRepository{
		packages: make(map[string]*entity.Recipient),
	}
}

func (r *InMemoryRecipientRepository) Create(recipientId *entity.Recipient) error {
	r.packages[recipientId.ID] = recipientId

	return nil
}

func (r *InMemoryRecipientRepository) FindById(recipientId string) (*entity.Recipient, error) {
	pkg, exists := r.packages[recipientId]

	if !exists {
		return nil, ErrRecipientDoesNotExist
	}

	return pkg, nil
}
