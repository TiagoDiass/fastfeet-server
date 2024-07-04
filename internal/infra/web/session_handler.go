package web

import (
	"encoding/json"
	"net/http"

	usecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/session"
)

type SessionHandler struct {
	CreateSessionUsecase *usecase.CreateSessionUsecase
}

func NewSessionHandler(createSessionUsecase *usecase.CreateSessionUsecase) *SessionHandler {
	return &SessionHandler{
		CreateSessionUsecase: createSessionUsecase,
	}
}

func (h *SessionHandler) CreateSession(w http.ResponseWriter, req *http.Request) {
	var input usecase.CreateSessionInputDTO
	err := json.NewDecoder(req.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	output, err := h.CreateSessionUsecase.Execute(input)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		error := Error{Message: "Unauthorized"}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
