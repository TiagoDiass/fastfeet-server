package repository

import "github.com/TiagoDiass/fastfeet-server/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
	// Update(user *entity.User) error
	FindById(id string) (*entity.User, error)
	FindByDocument(document string) (*entity.User, error)
	Delete(id string) error
}
