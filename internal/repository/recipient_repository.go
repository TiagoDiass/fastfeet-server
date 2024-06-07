package repository

import "github.com/TiagoDiass/fastfeet-server/internal/entity"

type RecipientRepository interface {
	Create(recipient *entity.Recipient) error
	FindById(id string) (*entity.Recipient, error)
}
