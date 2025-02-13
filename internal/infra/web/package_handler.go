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
	claims, err := GetClaimsFromContext(req.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	var input usecase.CreatePackageInputDTO
	err = json.NewDecoder(req.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	input.UserID = claims.User.ID
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
		errorResponse := Error{Message: "packageId is required"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	input.PackageID = packageId

	claims, err := GetClaimsFromContext(req.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	input.DeliverymanID = claims.User.ID

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

	claims, err := GetClaimsFromContext(req.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	err = json.NewDecoder(req.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	input.PackageID = packageId
	input.DeliverymanID = claims.User.ID

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
