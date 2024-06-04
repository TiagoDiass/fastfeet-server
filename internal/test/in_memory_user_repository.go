package test

import "github.com/TiagoDiass/fastfeet-server/internal/entity"

type InMemoryUserRepository struct {
	users map[string]*entity.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*entity.User),
	}
}

func (r *InMemoryUserRepository) Create(user *entity.User) error {
	r.users[user.ID] = user

	return nil
}

func (r *InMemoryUserRepository) FindById(userId string) (*entity.User, error) {
	user, exists := r.users[userId]

	if !exists {
		return nil, nil
	}

	return user, nil
}

func (r *InMemoryUserRepository) Delete(userId string) error {
	delete(r.users, userId)
	return nil
}
