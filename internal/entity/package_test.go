package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewPackage(t *testing.T) {
	pkg := NewPackage(
		"fake-recipient-id",
		"fake-deliveryman-id",
		"Package from Store_01 to John Doe",
		"WAITING_WITHDRAW",
	)

	require.Equal(t, pkg.Name, "Package from Store_01 to John Doe")
	require.Equal(t, pkg.RecipientId, "fake-recipient-id")
	require.Equal(t, pkg.DeliverymanId, "fake-deliveryman-id")
	require.Equal(t, pkg.Status, "WAITING_WITHDRAW")
}

func TestNewPackageWithOptionalFields(t *testing.T) {
	fakeTime := time.Now().Add(time.Hour * 24 * 10)
	fakeDeliveredPicture := "fake-delivered-picture-url"

	pkg := NewPackage(
		"fake-recipient-id",
		"fake-deliveryman-id",
		"Package from Store_01 to John Doe",
		"WAITING_WITHDRAW",
	).
		WithDeliveredPicture(&fakeDeliveredPicture).
		WithWithdrewAt(&fakeTime).
		WithDeliveredAt(&fakeTime)

	require.Equal(t, pkg.Name, "Package from Store_01 to John Doe")
	require.Equal(t, pkg.RecipientId, "fake-recipient-id")
	require.Equal(t, pkg.DeliverymanId, "fake-deliveryman-id")
	require.Equal(t, pkg.Status, "WAITING_WITHDRAW")
	require.Equal(t, *pkg.DeliveredPicture, "fake-delivered-picture-url")
	require.Equal(t, *pkg.WithdrewAt, fakeTime)
	require.Equal(t, *pkg.DeliveredAt, fakeTime)
}

func TestPackage_Withdraw(t *testing.T) {
	pkg := NewPackage(
		"fake-recipient-id",
		"fake-deliveryman-id",
		"Package from Store_01 to John Doe",
		"WAITING_WITHDRAW",
	)

	require.Equal(t, pkg.Status, "WAITING_WITHDRAW")
	require.Empty(t, pkg.WithdrewAt)

	pkg = pkg.Withdraw()

	require.Equal(t, pkg.Status, "ON_GOING")
	require.NotEmpty(t, pkg.WithdrewAt)
}
