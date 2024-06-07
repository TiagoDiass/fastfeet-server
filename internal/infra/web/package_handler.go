package web

import (
	"encoding/json"
	"net/http"

	usecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/package"
)

type PackageHandler struct {
	CreatePackageUsecase         *usecase.CreatePackageUsecase
	ListAvailablePackagesUsecase *usecase.ListAvailablePackagesUsecase
}

func NewPackageHandler(
	createPackageUsecase *usecase.CreatePackageUsecase,
	listAvailablePackagesUsecase *usecase.ListAvailablePackagesUsecase,
) *PackageHandler {
	return &PackageHandler{
		CreatePackageUsecase:         createPackageUsecase,
		ListAvailablePackagesUsecase: listAvailablePackagesUsecase,
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
