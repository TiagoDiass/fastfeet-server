package web

import (
	"encoding/json"
	"net/http"

	usecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/user"
)

type UserHandler struct {
	CreateUserUsecase *usecase.CreateUserUsecase
}

func NewUserHandler(createUserUsecase *usecase.CreateUserUsecase) *UserHandler {
	return &UserHandler{
		CreateUserUsecase: createUserUsecase,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, req *http.Request) {
	var input usecase.CreateUserInputDTO
	err := json.NewDecoder(req.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	output, err := h.CreateUserUsecase.Execute(input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
