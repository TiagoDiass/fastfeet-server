package entity

import (
	"time"

	"github.com/google/uuid"
)

type Package struct {
	ID          string    `json:"id"`
	RecipientId string    `json:"recipient_id"`
	Name        string    `json:"name"`
	Status      string    `json:"status"` // WAITING_WITHDRAW | ON_GOING | DELIVERED | RETURNED
	PostedAt    time.Time `json:"posted_at"`

	// optional / nullable fields
	DeliverymanId    *string    `json:"deliveryman_id"`
	DeliveredPicture *string    `json:"delivered_picture"`
	WithdrewAt       *time.Time `json:"withdrew_at"`
	DeliveredAt      *time.Time `json:"delivered_at"`
}

func NewPackage(recipientId, name string) *Package {
	p := &Package{
		ID:               uuid.NewString(),
		RecipientId:      recipientId,
		Name:             name,
		Status:           "WAITING_WITHDRAW",
		PostedAt:         time.Now(),
		DeliverymanId:    nil,
		DeliveredPicture: nil,
		WithdrewAt:       nil,
		DeliveredAt:      nil,
	}

	return p
}

func (p *Package) Withdraw(deliverymanId string) *Package {
	now := time.Now()

	p.WithdrewAt = &now
	p.DeliverymanId = &deliverymanId
	p.Status = "ON_GOING"

	return p
}

func (p *Package) MarkAsDelivered(deliveredPictureUrl string) *Package {
	now := time.Now()

	p.DeliveredAt = &now
	p.DeliveredPicture = &deliveredPictureUrl
	p.Status = "DELIVERED"

	return p
}
