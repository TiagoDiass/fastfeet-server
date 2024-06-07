package entity

import (
	"time"

	"github.com/google/uuid"
)

type Package struct {
	ID            string
	RecipientId   string
	DeliverymanId string
	Name          string
	Status        string // WAITING_WITHDRAW | ON_GOING | DELIVERED | RETURNED
	PostedAt      time.Time

	// optional / nullable fields
	DeliveredPicture *string
	WithdrewAt       *time.Time
	DeliveredAt      *time.Time
}

func NewPackage(recipientId, deliverymanId, name, status string) *Package {
	p := &Package{
		ID:               uuid.NewString(),
		RecipientId:      recipientId,
		DeliverymanId:    deliverymanId,
		Name:             name,
		Status:           status,
		PostedAt:         time.Now(),
		DeliveredPicture: nil,
		WithdrewAt:       nil,
		DeliveredAt:      nil,
	}

	return p
}

func (p *Package) WithDeliveredPicture(deliveredPictureUrl *string) *Package {
	p.DeliveredPicture = deliveredPictureUrl

	return p
}

func (p *Package) WithWithdrewAt(withdrewAt *time.Time) *Package {
	p.WithdrewAt = withdrewAt

	return p
}

func (p *Package) WithDeliveredAt(deliveredAt *time.Time) *Package {
	p.DeliveredAt = deliveredAt

	return p
}
