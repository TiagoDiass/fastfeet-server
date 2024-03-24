package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	user := NewUser(
		"fake-document",
		"fake-password",
		"fake-name",
		"email@example.com",
		"fake-phone",
		Admin,
	)

	require.NotEmpty(t, user.ID)
	require.Equal(t, user.Document, "fake-document")
	require.Equal(t, user.Password, "fake-password")
	require.Equal(t, user.Name, "fake-name")
	require.Equal(t, user.Email, "email@example.com")
	require.Equal(t, user.Phone, "fake-phone")
	require.Equal(t, user.Role, Admin)
	require.Equal(t, user.Role.String(), "admin")
	require.NotEmpty(t, user.CreatedAt)
}

func TestNewExistingUser(t *testing.T) {
	now := time.Now()

	user := NewExistingUser(
		"fake-id",
		"fake-document",
		"fake-password",
		"fake-name",
		"email@example.com",
		"fake-phone",
		Admin,
		now,
	)

	require.Equal(t, user.ID, "fake-id")
	require.Equal(t, user.CreatedAt, now)
}

func TestRoleEnum(t *testing.T) {
	unknownRole := Role(2)

	require.Equal(t, Admin.String(), "admin")
	require.Equal(t, DeliveryMan.String(), "deliveryman")
	require.Equal(t, unknownRole.String(), "unknown")
}
