package entity

import "github.com/google/uuid"

type Address struct {
	Street      string
	HouseNumber string
	District    string
	State       string
	Zipcode     string
}

type Recipient struct {
	ID   string
	Name string
	Address
}

func NewAddress(
	street,
	houseNumber,
	district,
	state,
	zipcode string,
) Address {
	return Address{
		Street:      street,
		HouseNumber: houseNumber,
		District:    district,
		State:       state,
		Zipcode:     zipcode,
	}
}

func NewRecipient(name string, address Address) Recipient {
	return Recipient{
		ID:      uuid.NewString(),
		Name:    name,
		Address: address,
	}
}

func NewRecipientWithID(
	id,
	name string,
	address Address,
) Recipient {
	recipient := NewRecipient(name, address)
	recipient.ID = id

	return recipient
}
