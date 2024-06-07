package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAddress(t *testing.T) {
	address := NewAddress("Main St", "123", "Downtown", "CA", "12345")

	require.Equal(t, address.Street, "Main St")
	require.Equal(t, address.HouseNumber, "123")
	require.Equal(t, address.District, "Downtown")
	require.Equal(t, address.State, "CA")
	require.Equal(t, address.Zipcode, "12345")
}

func TestNewRecipient(t *testing.T) {
	address := NewAddress("Main St", "123", "Downtown", "CA", "12345")
	recipient := NewRecipient("John Doe", "john@example.com", address)

	require.NotEmpty(t, recipient.ID)
	require.Equal(t, recipient.Name, "John Doe")
	require.Equal(t, recipient.Address.Street, "Main St")
	require.Equal(t, recipient.Address.HouseNumber, "123")
}
