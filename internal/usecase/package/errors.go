package usecase

import "errors"

var (
	ErrUserIsNotAdmin            = errors.New("unauthorized: only admins can create packages")
	ErrRecipientNotExists        = errors.New("recipient does not exist")
	ErrDeliverymanNotExists      = errors.New("deliveryman does not exist")
	ErrPackageNotExists          = errors.New("package does not exist")
	ErrPackageWasAlreadyWithdrew = errors.New("package was already withdrawn")
)
