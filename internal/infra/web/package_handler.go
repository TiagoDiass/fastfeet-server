package web

import (
	"encoding/json"
	"net/http"

	usecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/package"
	"github.com/go-chi/chi/v5"
)

type PackageHandler struct {
	CreatePackageUsecase           *usecase.CreatePackageUsecase
	ListAvailablePackagesUsecase   *usecase.ListAvailablePackagesUsecase
	ListDeliveredPackagesUsecase   *usecase.ListDeliveredPackagesUsecase
	WithdrawPackageUsecase         *usecase.WithdrawPackageUsecase
	ConfirmDeliveredPackageUsecase *usecase.ConfirmDeliveredPackageUsecase
}

func NewPackageHandler(
	createPackageUsecase *usecase.CreatePackageUsecase,
	listAvailablePackagesUsecase *usecase.ListAvailablePackagesUsecase,
	listDeliveredPackagesUsecase *usecase.ListDeliveredPackagesUsecase,
	withdrawPackageUsecase *usecase.WithdrawPackageUsecase,
	confirmDeliveredPackageUsecase *usecase.ConfirmDeliveredPackageUsecase,
) *PackageHandler {
	return &PackageHandler{
		CreatePackageUsecase:           createPackageUsecase,
		ListAvailablePackagesUsecase:   listAvailablePackagesUsecase,
		ListDeliveredPackagesUsecase:   listDeliveredPackagesUsecase,
		WithdrawPackageUsecase:         withdrawPackageUsecase,
		ConfirmDeliveredPackageUsecase: confirmDeliveredPackageUsecase,
	}
}

func (h *PackageHandler) CreatePackage(w http.ResponseWriter, req *http.Request) {
	// TODO: refactor later to get UserID on request headers or context, idk

	var input usecase.CreatePackageInputDTO
	err := json.NewDecoder(req.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	output, err := h.CreatePackageUsecase.Execute(input)

	if err != nil {
		switch err {
		case usecase.ErrUserIsNotAdmin:
			w.WriteHeader(http.StatusUnauthorized)
		case usecase.ErrRecipientNotExists, usecase.ErrDeliverymanNotExists:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (h *PackageHandler) ListAvailablePackages(w http.ResponseWriter, req *http.Request) {
	output, err := h.ListAvailablePackagesUsecase.Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *PackageHandler) ListDeliveredPackages(w http.ResponseWriter, req *http.Request) {
	output, err := h.ListDeliveredPackagesUsecase.Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *PackageHandler) WithdrawPackage(w http.ResponseWriter, req *http.Request) {
	var input usecase.WithdrawPackageInputDTO
	packageId := chi.URLParam(req, "packageId")

	if packageId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input.PackageID = packageId

	err := json.NewDecoder(req.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	output, err := h.WithdrawPackageUsecase.Execute(input)

	if err != nil {
		switch err {
		case usecase.ErrPackageNotExists, usecase.ErrDeliverymanNotExists:
			w.WriteHeader(http.StatusNotFound)
		case usecase.ErrPackageWasAlreadyWithdrew:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *PackageHandler) ConfirmDeliveredPackage(w http.ResponseWriter, req *http.Request) {
	var input usecase.ConfirmDeliveredPackageInputDTO
	packageId := chi.URLParam(req, "packageId")

	if packageId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input.PackageID = packageId

	err := json.NewDecoder(req.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	output, err := h.ConfirmDeliveredPackageUsecase.Execute(input)

	if err != nil {
		switch err {
		case usecase.ErrPackageNotExists:
			w.WriteHeader(http.StatusNotFound)
		case usecase.ErrPackageCannotBeDelivered, usecase.ErrDifferentDeliveryman:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
