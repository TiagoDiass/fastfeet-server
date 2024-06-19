package usecase

import "errors"

var (
	ErrUserIsNotAdmin            = errors.New("unauthorized: only admins can create packages")
	ErrRecipientNotExists        = errors.New("recipient does not exist")
	ErrDeliverymanNotExists      = errors.New("deliveryman does not exist")
	ErrPackageNotExists          = errors.New("package does not exist")
	ErrPackageWasAlreadyWithdrew = errors.New("package was already withdrawn")
	ErrPackageCannotBeDelivered  = errors.New("package cannot be delivered because its status is not waiting for withdraw")
	ErrDifferentDeliveryman      = errors.New("package delivery belongs to another deliveryman")
)
