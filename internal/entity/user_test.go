package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser(
		"fake-document",
		"fake-password",
		"fake-name",
		"email@example.com",
		"fake-phone",
		"admin",
	)

	require.Nil(t, err)
	require.NotEmpty(t, user.ID)
	require.Equal(t, user.Document, "fake-document")
	require.NotEmpty(t, user.Password)
	require.Equal(t, user.Name, "fake-name")
	require.Equal(t, user.Email, "email@example.com")
	require.Equal(t, user.Phone, "fake-phone")
	require.Equal(t, user.Role, "admin")
	require.NotEmpty(t, user.CreatedAt)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, _ := NewUser(
		"fake-document",
		"fake-password",
		"fake-name",
		"email@example.com",
		"fake-phone",
		"admin",
	)

	passwordsMatch := user.ValidatePassword("wrong-password")
	require.Equal(t, passwordsMatch, false)

	passwordsMatch = user.ValidatePassword("fake-password")
	require.Equal(t, passwordsMatch, true)
}
