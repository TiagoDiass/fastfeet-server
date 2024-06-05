package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id"`
	Document  string    `json:"document"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"` // "admin" or "deliveryman"
	CreatedAt time.Time `json:"createdAt"`
}

func NewUser(document, password, name, email, phone, role string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &User{
		ID:        uuid.NewString(),
		Document:  document,
		Password:  string(hash),
		Name:      name,
		Email:     email,
		Phone:     phone,
		Role:      role,
		CreatedAt: time.Now(),
	}

	return user, nil
}

func (u *User) ValidatePassword(password string) (passwordsMatch bool) {
	passwordInBytes := []byte(password)
	userPasswordInBytes := []byte(u.Password)

	err := bcrypt.CompareHashAndPassword(userPasswordInBytes, passwordInBytes)

	return err == nil
}
