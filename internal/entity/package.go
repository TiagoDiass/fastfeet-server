package entity

import (
	"time"

	"github.com/google/uuid"
)

type Package struct {
	ID            string    `json:"id"`
	RecipientId   string    `json:"recipient_id"`
	DeliverymanId string    `json:"deliveryman_id"`
	Name          string    `json:"name"`
	Status        string    `json:"status"` // WAITING_WITHDRAW | ON_GOING | DELIVERED | RETURNED
	PostedAt      time.Time `json:"posted_at"`

	// optional / nullable fields
	DeliveredPicture *string    `json:"delivered_picture"`
	WithdrewAt       *time.Time `json:"withdrew_at"`
	DeliveredAt      *time.Time `json:"delivered_at"`
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

func (p *Package) Withdraw() *Package {
	now := time.Now()

	p.WithdrewAt = &now
	p.Status = "ON_GOING"

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
