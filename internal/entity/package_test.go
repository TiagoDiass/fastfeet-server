package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPackage(t *testing.T) {
	pkg := NewPackage(
		"fake-recipient-id",
		"Package from Store_01 to John Doe",
	)

	require.NotEmpty(t, pkg.ID)
	require.Equal(t, pkg.Name, "Package from Store_01 to John Doe")
	require.Equal(t, pkg.RecipientId, "fake-recipient-id")
	require.Equal(t, pkg.Status, "WAITING_WITHDRAW")
	require.NotNil(t, pkg.PostedAt)
	require.Nil(t, pkg.DeliverymanId)
	require.Nil(t, pkg.DeliveredPicture)
	require.Nil(t, pkg.WithdrewAt)
	require.Nil(t, pkg.DeliveredAt)
}

func TestPackage_Withdraw(t *testing.T) {
	pkg := NewPackage(
		"fake-recipient-id",
		"Package from Store_01 to John Doe",
	)

	require.Equal(t, pkg.Status, "WAITING_WITHDRAW")
	require.Nil(t, pkg.WithdrewAt)
	require.Nil(t, pkg.DeliverymanId)

	pkg = pkg.Withdraw("fake-deliveryman-id")

	require.Equal(t, pkg.Status, "ON_GOING")
	require.Equal(t, *pkg.DeliverymanId, "fake-deliveryman-id")
	require.NotNil(t, pkg.WithdrewAt)
}
