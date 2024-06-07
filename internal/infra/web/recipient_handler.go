package web

import (
	"encoding/json"
	"net/http"

	usecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/recipient"
)

type RecipientHandler struct {
	CreateRecipientUsecase *usecase.CreateRecipientUsecase
}

func NewRecipientHandler(createRecipientUsecase *usecase.CreateRecipientUsecase) *RecipientHandler {
	return &RecipientHandler{
		CreateRecipientUsecase: createRecipientUsecase,
	}
}

func (h *RecipientHandler) CreateRecipient(w http.ResponseWriter, req *http.Request) {
	var input usecase.CreateRecipientInputDTO
	err := json.NewDecoder(req.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	output, err := h.CreateRecipientUsecase.Execute(input)

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
