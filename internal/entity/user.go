package entity

import "github.com/google/uuid"

type Role uint8

const (
	Admin Role = iota
	DeliveryMan
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "admin"
	case DeliveryMan:
		return "deliveryman"
	}

	return "unknown"
}

type User struct {
	ID       string
	Document string
	Password string
	Name     string
	Email    string
	Phone    string
	Role     Role
}

func NewUser(document, password, name, email, phone string, role Role) User {
	return User{
		ID:       uuid.NewString(),
		Document: document,
		Password: password,
		Name:     name,
		Email:    email,
		Phone:    phone,
	}
}
