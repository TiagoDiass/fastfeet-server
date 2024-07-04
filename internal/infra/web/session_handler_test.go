package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/test"
	usecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/session"
	"github.com/go-chi/jwtauth"
	"github.com/stretchr/testify/require"
)

func makeSessionHandlerSut() *SessionHandler {
	userRepository := test.NewInMemoryUserRepository()
	jwt := jwtauth.New("HS256", []byte("secret"), nil)
	expiresIn := 600

	user, _ := entity.NewUser(
		"87847048027",
		"beautiful-password",
		"John Doe",
		"john@example.com",
		"fake-phone",
		"admin",
	)
	userRepository.Create(user)

	createSessionUsecase := usecase.NewCreateSessionUsecase(
		userRepository,
		jwt,
		expiresIn,
	)
	sessionHandler := NewSessionHandler(createSessionUsecase)

	return sessionHandler
}

func TestSessionHandler_CreateSessionSuccessCase(t *testing.T) {
	input := usecase.CreateSessionInputDTO{
		Document: "87847048027",
		Password: "beautiful-password",
	}
	inputJSON, _ := json.Marshal(input)

	sessionHandler := makeSessionHandlerSut()
	req := httptest.NewRequest(
		"POST",
		"/session",
		bytes.NewReader(inputJSON),
	)
	w := httptest.NewRecorder()

	sessionHandler.CreateSession(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var output usecase.CreateSessionOutputDTO
	err := json.NewDecoder(w.Body).Decode(&output)

	require.Nil(t, err)
	require.NotEmpty(t, output.AccessToken)
}

// func TestSessionHandler_CreateSessionWithInvalidJSON(t *testing.T) {
// 	invalidJSON := []byte(`{
// 		"document": "87847048027",
// 		"password": "beautiful-password"
// 	`) // Missing closing brace

// 	sessionHandler := makeSessionHandlerSut()
// 	req := httptest.NewRequest(
// 		"POST",
// 		"/sessions",
// 		bytes.NewReader(invalidJSON),
// 	)
// 	w := httptest.NewRecorder()

// 	sessionHandler.CreateSession(w, req)

// 	require.Equal(t, http.StatusBadRequest, w.Code)

// 	var errorResponse Error
// 	err := json.NewDecoder(w.Body).Decode(&errorResponse)

// 	require.Nil(t, err)
// 	require.NotEmpty(t, errorResponse.Message)
// }

// func TestSessionHandler_CreateSessionWhenUserNotExists(t *testing.T) {
// 	input := usecase.CreateSessionInputDTO{
// 		Document: "wrong-document",
// 		Password: "any-password",
// 	}
// 	inputJSON, _ := json.Marshal(input)

// 	sessionHandler := makeSessionHandlerSut()
// 	req := httptest.NewRequest(
// 		"POST",
// 		"/sessions",
// 		bytes.NewReader(inputJSON),
// 	)
// 	w := httptest.NewRecorder()

// 	sessionHandler.CreateSession(w, req)

// 	require.Equal(t, http.StatusUnauthorized, w.Code)

// 	var errorResponse Error
// 	err := json.NewDecoder(w.Body).Decode(&errorResponse)

// 	require.Nil(t, err)
// 	require.Equal(t, errorResponse.Message, "Unauthorized")
// }

// func TestSessionHandler_CreateSessionWithWrongPassword(t *testing.T) {
// 	input := usecase.CreateSessionInputDTO{
// 		Document: "87847048027",
// 		Password: "wrong-password",
// 	}
// 	inputJSON, _ := json.Marshal(input)

// 	sessionHandler := makeSessionHandlerSut()
// 	req := httptest.NewRequest(
// 		"POST",
// 		"/sessions",
// 		bytes.NewReader(inputJSON),
// 	)
// 	w := httptest.NewRecorder()

// 	sessionHandler.CreateSession(w, req)

// 	require.Equal(t, http.StatusUnauthorized, w.Code)

// 	var errorResponse Error
// 	err := json.NewDecoder(w.Body).Decode(&errorResponse)

// 	require.Nil(t, err)
// 	require.Equal(t, errorResponse.Message, "Unauthorized")
// }
