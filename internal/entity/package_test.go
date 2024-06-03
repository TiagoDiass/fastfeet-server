package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewPackage(t *testing.T) {
	p := NewPackage(
		"fake-recipient-id",
		"fake-deliveryman-id",
		"Package from Store_01 to John Doe",
		"WAITING_WITHDRAW",
	)

	require.Equal(t, p.Name, "Package from Store_01 to John Doe")
	require.Equal(t, p.RecipientId, "fake-recipient-id")
	require.Equal(t, p.DeliverymanId, "fake-deliveryman-id")
	require.Equal(t, p.Status, "WAITING_WITHDRAW")
}

func TestNewPackageWithOptionalFields(t *testing.T) {
	fakeTime := time.Now().Add(time.Hour * 24 * 10)

	p := NewPackage(
		"fake-recipient-id",
		"fake-deliveryman-id",
		"Package from Store_01 to John Doe",
		"WAITING_WITHDRAW",
	).
		WithDeliveredPicture("fake-delivered-picture-url").
		WithWithdrewAt(fakeTime).
		WithDeliveredAt(fakeTime)

	require.Equal(t, p.Name, "Package from Store_01 to John Doe")
	require.Equal(t, p.RecipientId, "fake-recipient-id")
	require.Equal(t, p.DeliverymanId, "fake-deliveryman-id")
	require.Equal(t, p.Status, "WAITING_WITHDRAW")
	require.Equal(t, p.DeliveredPicture, "fake-delivered-picture-url")
	require.Equal(t, p.WithdrewAt, fakeTime)
	require.Equal(t, p.DeliveredAt, fakeTime)
}
