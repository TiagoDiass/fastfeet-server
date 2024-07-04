package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TiagoDiass/fastfeet-server/internal/test"
	usecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/recipient"
	"github.com/stretchr/testify/require"
)

func makeRecipientHandlerSut() *RecipientHandler {
	recipientRepository := test.NewInMemoryRecipientRepository()
	createRecipientUsecase := usecase.NewCreateRecipientUsecase(recipientRepository)
	recipientHandler := NewRecipientHandler(createRecipientUsecase)

	return recipientHandler
}

func TestRecipientHandler_CreateRecipientSuccessCase(t *testing.T) {
	input := usecase.CreateRecipientInputDTO{
		Name:        "Jane Doe",
		Email:       "janedoe@example.com",
		Street:      "123 Main St",
		HouseNumber: "456",
		District:    "Downtown",
		State:       "CA",
		Zipcode:     "12345",
	}
	inputJSON, _ := json.Marshal(input)

	recipientHandler := makeRecipientHandlerSut()
	req := httptest.NewRequest(
		"POST",
		"/recipients",
		bytes.NewReader(inputJSON),
	)
	w := httptest.NewRecorder()

	recipientHandler.CreateRecipient(w, req)

	require.Equal(t, http.StatusCreated, w.Code)

	var output usecase.CreateRecipientOutputDTO
	err := json.NewDecoder(w.Body).Decode(&output)

	require.Nil(t, err)
	require.NotEmpty(t, output.ID)
	require.Equal(t, input.Name, output.Name)
	require.Equal(t, input.Email, output.Email)
	require.Equal(t, input.Street, output.Street)
	require.Equal(t, input.HouseNumber, output.HouseNumber)
	require.Equal(t, input.District, output.District)
	require.Equal(t, input.State, output.State)
	require.Equal(t, input.Zipcode, output.Zipcode)
}

func TestRecipientHandler_CreateRecipientWithInvalidJSON(t *testing.T) {
	invalidJSON := []byte(`{
		"name": "Jane Doe",
		"email": "janedoe@example.com"
	`) // Missing closing brace

	recipientHandler := makeRecipientHandlerSut()
	req := httptest.NewRequest(
		"POST",
		"/recipients",
		bytes.NewReader(invalidJSON),
	)
	w := httptest.NewRecorder()

	recipientHandler.CreateRecipient(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)

	var errorResponse Error
	err := json.NewDecoder(w.Body).Decode(&errorResponse)

	require.Nil(t, err)
	require.NotEmpty(t, errorResponse.Message)
}

func TestRecipientHandler_CreateRecipientWhenUsecaseReturnsError(t *testing.T) {
	input := usecase.CreateRecipientInputDTO{
		Name:        "Jane Doe",
		Email:       test.EmailThatReturnsErrorOnCreateRecipient,
		Street:      "123 Main St",
		HouseNumber: "456",
		District:    "Downtown",
		State:       "CA",
		Zipcode:     "12345",
	}
	inputJSON, _ := json.Marshal(input)

	recipientHandler := makeRecipientHandlerSut()
	req := httptest.NewRequest(
		"POST",
		"/recipients",
		bytes.NewReader(inputJSON),
	)
	w := httptest.NewRecorder()

	recipientHandler.CreateRecipient(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)

	var output Error
	err := json.NewDecoder(w.Body).Decode(&output)

	require.Nil(t, err)
	require.Equal(t, output.Message, test.ErrOnCreateRecipient.Error())
}
