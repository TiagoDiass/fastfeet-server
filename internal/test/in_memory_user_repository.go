package test

import (
	"errors"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
)

var (
	ErrUserNotExists                     = errors.New("user does not exist")
	ErrDeliverymanNotExists              = errors.New("user does not exist")
	ErrOnCreateUser                      = errors.New("mocked error while creating user")
	DocumentThatReturnsErrorOnCreateUser = "52780765003"
)

type InMemoryUserRepository struct {
	users                  map[string]*entity.User
	usersIndexedByDocument map[string]*entity.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:                  make(map[string]*entity.User),
		usersIndexedByDocument: make(map[string]*entity.User),
	}
}

func (r *InMemoryUserRepository) Create(user *entity.User) error {
	if user.Document == DocumentThatReturnsErrorOnCreateUser {
		return ErrOnCreateUser
	}

	r.users[user.ID] = user
	r.usersIndexedByDocument[user.Document] = user

	return nil
}

func (r *InMemoryUserRepository) FindById(userId string) (*entity.User, error) {
	user, exists := r.users[userId]

	if !exists {
		return nil, ErrUserNotExists
	}

	return user, nil
}

func (r *InMemoryUserRepository) Delete(userId string) error {
	delete(r.users, userId)
	return nil
}

func (r *InMemoryUserRepository) FindByDocument(document string) (*entity.User, error) {
	user, exists := r.usersIndexedByDocument[document]

	if !exists {
		return nil, ErrUserNotExists
	}

	return user, nil
}

func (r *InMemoryUserRepository) FindAll() ([]*entity.User, error) {
	var users []*entity.User

	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

func (r *InMemoryUserRepository) FindDeliverymanById(deliverymanId string) (*entity.User, error) {
	user, exists := r.users[deliverymanId]

	if !exists || user.Role != "deliveryman" {
		return nil, ErrDeliverymanNotExists
	}

	return user, nil
}
