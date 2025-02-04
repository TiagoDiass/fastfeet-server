package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/test"
	usecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/package"
	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwt"

	"github.com/stretchr/testify/require"
)

func makePackageHandlerSut() *PackageHandler {
	userRepository := test.NewInMemoryUserRepository()
	recipientRepository := test.NewInMemoryRecipientRepository()
	packageRepository := test.NewInMemoryPackageRepository()

	// Create users and recipient
	admin, _ := entity.NewUser(
		"33872158007",
		"beautiful-password",
		"Admin",
		"admin@example.com",
		"fake-phone",
		"admin",
	)
	admin.ID = "admin-id"
	userRepository.Create(admin)

	deliveryman, _ := entity.NewUser(
		"36799751044",
		"beautiful-password",
		"Deliveryman",
		"deliveryman@example.com",
		"fake-phone",
		"deliveryman",
	)
	deliveryman.ID = "deliveryman-id"
	userRepository.Create(deliveryman)

	address := entity.NewAddress("Main St", "123", "Downtown", "CA", "12345")
	recipient := entity.NewRecipient(
		"Jane Doe",
		"jane@example.com",
		address,
	)
	recipient.ID = "recipient-id"
	recipientRepository.Create(recipient)

	createPackageUsecase := usecase.NewCreatePackageUsecase(
		packageRepository,
		userRepository,
		recipientRepository,
	)
	listAvailablePackagesUsecase := usecase.NewListAvailablePackagesUsecase(packageRepository)
	listDeliveredPackagesUsecase := usecase.NewListDeliveredPackagesUsecase(packageRepository)
	withdrawPackageUsecase := usecase.NewWithdrawPackageUsecase(packageRepository, userRepository)
	confirmDeliveredPackageUsecase := usecase.NewConfirmDeliveredPackageUsecase(packageRepository)

	packageHandler := NewPackageHandler(
		createPackageUsecase,
		listAvailablePackagesUsecase,
		listDeliveredPackagesUsecase,
		withdrawPackageUsecase,
		confirmDeliveredPackageUsecase,
	)

	return packageHandler
}

func generateToken(userID string) jwt.Token {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	claims := map[string]interface{}{
		"sub": userID,
		"user": map[string]interface{}{
			"id": userID,
		},
	}
	_, tokenString, _ := tokenAuth.Encode(claims)
	token, _ := tokenAuth.Decode(tokenString)

	return token
}

func TestPackageHandler_CreatePackageSuccessCase(t *testing.T) {
	input := usecase.CreatePackageInputDTO{
		RecipientID: "recipient-id",
		Name:        "Sample Package",
	}
	inputJSON, _ := json.Marshal(input)

	packageHandler := makePackageHandlerSut()

	token := generateToken("admin-id")

	req := httptest.NewRequest(
		"POST",
		"/packages",
		bytes.NewReader(inputJSON),
	)

	ctx := jwtauth.NewContext(req.Context(), token, nil)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()

	packageHandler.CreatePackage(w, req)

	require.Equal(t, http.StatusCreated, w.Code)

	var output entity.Package
	err := json.NewDecoder(w.Body).Decode(&output)

	require.Nil(t, err)
	require.NotEmpty(t, output.ID)
	require.Equal(t, input.RecipientID, output.RecipientId)
	require.Equal(t, input.Name, output.Name)
}

// func TestPackageHandler_CreatePackageWithInvalidJSON(t *testing.T) {
// 	invalidJSON := []byte(`{
// 		"recipient_id": "recipient-id",
// 		"name": "Sample Package"
// 	`) // Missing closing brace

// 	packageHandler := makePackageHandlerSut()
// 	token := generateToken("admin-id")

// 	// Create a context with the token
// 	ctx := context.WithValue(context.Background(), jwtauth.TokenCtxKey, token)

// 	req := httptest.NewRequest(
// 		"POST",
// 		"/packages",
// 		bytes.NewReader(invalidJSON),
// 	)
// 	req = req.WithContext(ctx)
// 	w := httptest.NewRecorder()

// 	packageHandler.CreatePackage(w, req)

// 	require.Equal(t, http.StatusBadRequest, w.Code)

// 	var errorResponse Error
// 	err := json.NewDecoder(w.Body).Decode(&errorResponse)

// 	require.Nil(t, err)
// 	require.NotEmpty(t, errorResponse.Message)
// }

// func TestPackageHandler_CreatePackageWhenUserIsNotAdmin(t *testing.T) {
// 	input := usecase.CreatePackageInputDTO{
// 		RecipientID: "recipient-id",
// 		Name:        "Sample Package",
// 	}
// 	inputJSON, _ := json.Marshal(input)

// 	packageHandler := makePackageHandlerSut()
// 	token := generateToken("deliveryman-id")

// 	// Create a context with the token
// 	ctx := context.WithValue(context.Background(), jwtauth.TokenCtxKey, token)

// 	req := httptest.NewRequest(
// 		"POST",
// 		"/packages",
// 		bytes.NewReader(inputJSON),
// 	)
// 	req = req.WithContext(ctx)
// 	w := httptest.NewRecorder()

// 	packageHandler.CreatePackage(w, req)

// 	require.Equal(t, http.StatusUnauthorized, w.Code)

// 	var errorResponse Error
// 	err := json.NewDecoder(w.Body).Decode(&errorResponse)

// 	require.Nil(t, err)
// 	require.Equal(t, errorResponse.Message, usecase.ErrUserIsNotAdmin.Error())
// }

// func TestPackageHandler_CreatePackageWhenRecipientDoesNotExist(t *testing.T) {
// 	input := usecase.CreatePackageInputDTO{
// 		RecipientID: "non-existent-recipient-id",
// 		Name:        "Sample Package",
// 	}
// 	inputJSON, _ := json.Marshal(input)

// 	packageHandler := makePackageHandlerSut()
// 	token := generateToken("admin-id")

// 	// Create a context with the token
// 	ctx := context.WithValue(context.Background(), jwtauth.TokenCtxKey, token)

// 	req := httptest.NewRequest(
// 		"POST",
// 		"/packages",
// 		bytes.NewReader(inputJSON),
// 	)
// 	req = req.WithContext(ctx)
// 	w := httptest.NewRecorder()

// 	packageHandler.CreatePackage(w, req)

// 	require.Equal(t, http.StatusBadRequest, w.Code)

// 	var errorResponse Error
// 	err := json.NewDecoder(w.Body).Decode(&errorResponse)

// 	require.Nil(t, err)
// 	require.Equal(t, errorResponse.Message, usecase.ErrRecipientNotExists.Error())
// }

// func TestPackageHandler_CreatePackageWhenRepositoryReturnsAnError(t *testing.T) {
// 	input := usecase.CreatePackageInputDTO{
// 		RecipientID: "recipient-id",
// 		Name:        test.NameThatReturnsErrorOnCreatePackage,
// 	}
// 	inputJSON, _ := json.Marshal(input)

// 	packageHandler := makePackageHandlerSut()
// 	token := generateToken("admin-id")

// 	// Create a context with the token
// 	ctx := context.WithValue(context.Background(), jwtauth.TokenCtxKey, token)

// 	req := httptest.NewRequest(
// 		"POST",
// 		"/packages",
// 		bytes.NewReader(inputJSON),
// 	)
// 	req = req.WithContext(ctx)
// 	w := httptest.NewRecorder()

// 	packageHandler.CreatePackage(w, req)

// 	require.Equal(t, http.StatusInternalServerError, w.Code)

// 	var errorResponse Error
// 	err := json.NewDecoder(w.Body).Decode(&errorResponse)

// 	require.Nil(t, err)
// 	require.Equal(t, errorResponse.Message, test.ErrOnCreatePackage.Error())
// }
