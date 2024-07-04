package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TiagoDiass/fastfeet-server/internal/test"
	usecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/user"
	"github.com/stretchr/testify/require"
)

func makeUserHandlerSut() *UserHandler {
	userRepository := test.NewInMemoryUserRepository()
	createUserUsecase := usecase.NewCreateUserUsecase(userRepository)
	userHandler := NewUserHandler(createUserUsecase)

	return userHandler
}

func TestUserHandler_CreateUserSuccessCase(t *testing.T) {
	input := usecase.CreateUserInputDTO{
		Document: "87847048027",
		Password: "beautiful-password",
		Name:     "John Doe",
		Email:    "johndoe@gmail.com",
		Phone:    "19912341234",
		Role:     "admin",
	}
	inputJSON, _ := json.Marshal(input)

	userHandler := makeUserHandlerSut()
	req := httptest.NewRequest(
		"POST",
		"/users",
		bytes.NewReader(inputJSON),
	)
	w := httptest.NewRecorder()

	userHandler.CreateUser(w, req)

	require.Equal(t, http.StatusCreated, w.Code)

	var output usecase.CreateUserOutputDTO
	err := json.NewDecoder(w.Body).Decode(&output)

	require.Nil(t, err)
	require.NotEmpty(t, output.ID)
	require.Equal(t, input.Document, output.Document)
	require.Equal(t, input.Name, output.Name)
	require.Equal(t, input.Email, output.Email)
	require.Equal(t, input.Phone, output.Phone)
	require.Equal(t, input.Role, output.Role)
}

func TestUserHandler_CreateUserWithInvalidJSON(t *testing.T) {
	invalidJSON := []byte(`{
		"document": "87847048027", 
		"password": "beautiful-password"
	`) // Missing closing brace

	userHandler := makeUserHandlerSut()
	req := httptest.NewRequest(
		"POST",
		"/users",
		bytes.NewReader(invalidJSON),
	)
	w := httptest.NewRecorder()

	userHandler.CreateUser(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)

	var errorResponse Error
	err := json.NewDecoder(w.Body).Decode(&errorResponse)

	require.Nil(t, err)
	require.NotEmpty(t, errorResponse.Message)
}

func TestUserHandler_CreateUserWhenUsecaseReturnsError(t *testing.T) {
	input := usecase.CreateUserInputDTO{
		Document: test.DocumentThatReturnsErrorOnCreateUser,
		Password: "beautiful-password",
		Name:     "John Doe",
		Email:    "johndoe@gmail.com",
		Phone:    "19912341234",
		Role:     "admin",
	}
	inputJSON, _ := json.Marshal(input)

	userHandler := makeUserHandlerSut()
	req := httptest.NewRequest(
		"POST",
		"/users",
		bytes.NewReader(inputJSON),
	)
	w := httptest.NewRecorder()

	userHandler.CreateUser(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)

	var output Error
	err := json.NewDecoder(w.Body).Decode(&output)

	require.Nil(t, err)
	require.Equal(t, output.Message, test.ErrOnCreateUser.Error())
}
