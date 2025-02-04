package web

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/test"
	usecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/package"
	"github.com/go-chi/chi/v5"
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

func TestPackageHandler_CreatePackageWithInvalidJSON(t *testing.T) {
	invalidJSON := []byte(`{
		"recipient_id": "recipient-id",
		"name": "Sample Package"
	`) // Missing closing brace

	packageHandler := makePackageHandlerSut()
	token := generateToken("admin-id")

	req := httptest.NewRequest(
		"POST",
		"/packages",
		bytes.NewReader(invalidJSON),
	)

	ctx := jwtauth.NewContext(req.Context(), token, nil)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()

	packageHandler.CreatePackage(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)

	var errorResponse Error
	err := json.NewDecoder(w.Body).Decode(&errorResponse)

	require.Nil(t, err)
	require.NotEmpty(t, errorResponse.Message)
}

func TestPackageHandler_CreatePackageWhenUserIsNotAdmin(t *testing.T) {
	input := usecase.CreatePackageInputDTO{
		RecipientID: "recipient-id",
		Name:        "Sample Package",
	}
	inputJSON, _ := json.Marshal(input)

	packageHandler := makePackageHandlerSut()
	token := generateToken("deliveryman-id")

	req := httptest.NewRequest(
		"POST",
		"/packages",
		bytes.NewReader(inputJSON),
	)

	ctx := jwtauth.NewContext(req.Context(), token, nil)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()

	packageHandler.CreatePackage(w, req)

	require.Equal(t, http.StatusUnauthorized, w.Code)

	var errorResponse Error
	err := json.NewDecoder(w.Body).Decode(&errorResponse)

	require.Nil(t, err)
	require.Equal(t, errorResponse.Message, usecase.ErrUserIsNotAdmin.Error())
}

func TestPackageHandler_CreatePackageWhenRecipientDoesNotExist(t *testing.T) {
	input := usecase.CreatePackageInputDTO{
		RecipientID: "non-existent-recipient-id",
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

	require.Equal(t, http.StatusBadRequest, w.Code)

	var errorResponse Error
	err := json.NewDecoder(w.Body).Decode(&errorResponse)

	require.Nil(t, err)
	require.Equal(t, errorResponse.Message, usecase.ErrRecipientNotExists.Error())
}

func TestPackageHandler_CreatePackageWhenRepositoryReturnsAnError(t *testing.T) {
	input := usecase.CreatePackageInputDTO{
		RecipientID: "recipient-id",
		Name:        test.NameThatReturnsErrorOnCreatePackage,
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

	require.Equal(t, http.StatusInternalServerError, w.Code)

	var errorResponse Error
	err := json.NewDecoder(w.Body).Decode(&errorResponse)

	require.Nil(t, err)
	require.Equal(t, errorResponse.Message, test.ErrOnCreatePackage.Error())
}

func TestPackageHandler_ListAvailablePackagesSuccessCase(t *testing.T) {
	packageHandler := makePackageHandlerSut()

	pkg1 := entity.NewPackage("recipient-id", "Package 1")
	pkg2 := entity.NewPackage("recipient-id", "Package 2")

	packageHandler.ListAvailablePackagesUsecase.PackageRepository.Create(pkg1)
	packageHandler.ListAvailablePackagesUsecase.PackageRepository.Create(pkg2)

	req := httptest.NewRequest("GET", "/packages/available", nil)
	w := httptest.NewRecorder()

	packageHandler.ListAvailablePackages(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var output []*entity.Package
	err := json.NewDecoder(w.Body).Decode(&output)

	require.Nil(t, err)
	require.Len(t, output, 2)
	require.Equal(t, "Package 1", output[0].Name)
	require.Equal(t, "Package 2", output[1].Name)
}

func TestPackageHandler_ListAvailablePackagesWhenRepositoryReturnsAnError(t *testing.T) {
	packageHandler := makePackageHandlerSut()

	pkg := entity.NewPackage("recipient-id", test.NameThatReturnsErrorOnFindPackages)
	packageHandler.ListAvailablePackagesUsecase.PackageRepository.Create(pkg)

	req := httptest.NewRequest("GET", "/packages/available", nil)
	w := httptest.NewRecorder()

	packageHandler.ListAvailablePackages(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)

	var errorResponse Error
	err := json.NewDecoder(w.Body).Decode(&errorResponse)

	require.Nil(t, err)
	require.Equal(t, test.ErrOnFindPackages.Error(), errorResponse.Message)
}

func TestPackageHandler_ListDeliveredPackagesSuccessCase(t *testing.T) {
	packageHandler := makePackageHandlerSut()

	pkg1 := entity.NewPackage("recipient-id", "Package 1")
	pkg1.Withdraw("deliveryman-id")
	pkg1.MarkAsDelivered("fake-picture-url-1")

	pkg2 := entity.NewPackage("recipient-id", "Package 2")
	pkg2.Withdraw("deliveryman-id")
	pkg2.MarkAsDelivered("fake-picture-url-2")

	packageHandler.ListDeliveredPackagesUsecase.PackageRepository.Create(pkg1)
	packageHandler.ListDeliveredPackagesUsecase.PackageRepository.Create(pkg2)

	req := httptest.NewRequest("GET", "/packages/delivered", nil)
	w := httptest.NewRecorder()

	packageHandler.ListDeliveredPackages(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var output []*entity.Package
	err := json.NewDecoder(w.Body).Decode(&output)

	require.Nil(t, err)
	require.Len(t, output, 2)
	require.Equal(t, "Package 1", output[0].Name)
	require.Equal(t, "DELIVERED", output[0].Status)
	require.Equal(t, "Package 2", output[1].Name)
	require.Equal(t, "DELIVERED", output[1].Status)
}

func TestPackageHandler_ListDeliveredPackagesWhenRepositoryReturnsAnError(t *testing.T) {
	packageHandler := makePackageHandlerSut()

	pkg := entity.NewPackage("recipient-id", test.NameThatReturnsErrorOnFindPackages)
	pkg.Withdraw("deliveryman-id")
	pkg.MarkAsDelivered("fake-picture-url")

	packageHandler.ListDeliveredPackagesUsecase.PackageRepository.Create(pkg)

	req := httptest.NewRequest("GET", "/packages/delivered", nil)
	w := httptest.NewRecorder()

	packageHandler.ListDeliveredPackages(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)

	var errorResponse Error
	err := json.NewDecoder(w.Body).Decode(&errorResponse)

	require.Nil(t, err)
	require.Equal(t, test.ErrOnFindPackages.Error(), errorResponse.Message)
}

func TestPackageHandler_WithdrawSuccessCase(t *testing.T) {
	packageHandler := makePackageHandlerSut()

	pkg := entity.NewPackage("recipient-id", "Package 1")
	pkg.ID = "package-id"
	packageHandler.WithdrawPackageUsecase.PackageRepository.Create(pkg)

	token := generateToken("deliveryman-id")
	req := httptest.NewRequest("PATCH", "/packages/withdraw/package-id", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("packageId", "package-id")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	ctx := jwtauth.NewContext(req.Context(), token, nil)
	req = req.WithContext(ctx)
	w := httptest.NewRecorder()

	packageHandler.WithdrawPackage(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var output entity.Package
	err := json.NewDecoder(w.Body).Decode(&output)

	require.Nil(t, err)
	require.Equal(t, "ON_GOING", output.Status)
	require.Equal(t, "deliveryman-id", *output.DeliverymanId)
	require.NotNil(t, output.WithdrewAt)
}

func TestPackageHandler_WithdrawWhenPackageIdIsMissingInURLParams(t *testing.T) {
	packageHandler := makePackageHandlerSut()
	token := generateToken("deliveryman-id")

	req := httptest.NewRequest("PATCH", "/packages/withdraw/", nil)

	rctx := chi.NewRouteContext()
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	ctx := jwtauth.NewContext(req.Context(), token, nil)
	req = req.WithContext(ctx)
	w := httptest.NewRecorder()

	packageHandler.WithdrawPackage(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPackageHandler_WithdrawWhenPackageNotExists(t *testing.T) {
	packageHandler := makePackageHandlerSut()
	token := generateToken("deliveryman-id")

	req := httptest.NewRequest("PATCH", "/packages/withdraw/non-existent-package-id", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("packageId", "non-existent-package-id")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	ctx := jwtauth.NewContext(req.Context(), token, nil)
	req = req.WithContext(ctx)
	w := httptest.NewRecorder()

	packageHandler.WithdrawPackage(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)

	var errorResponse Error
	err := json.NewDecoder(w.Body).Decode(&errorResponse)

	require.Nil(t, err)
	require.Equal(t, usecase.ErrPackageNotExists.Error(), errorResponse.Message)
}

func TestPackageHandler_WithdrawWhenDeliverymanNotExists(t *testing.T) {
	packageHandler := makePackageHandlerSut()

	pkg := entity.NewPackage("recipient-id", "Package 1")
	pkg.ID = "package-id"
	packageHandler.WithdrawPackageUsecase.PackageRepository.Create(pkg)

	token := generateToken("non-existent-deliveryman-id")
	req := httptest.NewRequest("PATCH", "/packages/withdraw/package-id", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("packageId", "package-id")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	ctx := jwtauth.NewContext(req.Context(), token, nil)
	req = req.WithContext(ctx)
	w := httptest.NewRecorder()

	packageHandler.WithdrawPackage(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)

	var errorResponse Error
	err := json.NewDecoder(w.Body).Decode(&errorResponse)

	require.Nil(t, err)
	require.Equal(t, usecase.ErrDeliverymanNotExists.Error(), errorResponse.Message)
}

func TestPackageHandler_WithdrawWhenPackageWasAlreadyWithdrew(t *testing.T) {
	packageHandler := makePackageHandlerSut()

	pkg := entity.NewPackage("recipient-id", "Package 1")
	pkg.ID = "package-id"
	pkg.Withdraw("other-deliveryman-id")
	packageHandler.WithdrawPackageUsecase.PackageRepository.Create(pkg)

	token := generateToken("deliveryman-id")
	req := httptest.NewRequest("PATCH", "/packages/withdraw/package-id", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("packageId", "package-id")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	ctx := jwtauth.NewContext(req.Context(), token, nil)
	req = req.WithContext(ctx)
	w := httptest.NewRecorder()

	packageHandler.WithdrawPackage(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)

	var errorResponse Error
	err := json.NewDecoder(w.Body).Decode(&errorResponse)

	require.Nil(t, err)
	require.Equal(t, usecase.ErrPackageWasAlreadyWithdrew.Error(), errorResponse.Message)
}

func TestPackageHandler_WithdrawWhenRepositoryReturnsAnError(t *testing.T) {
	packageHandler := makePackageHandlerSut()

	pkg := entity.NewPackage("recipient-id", test.NameThatReturnsErrorOnUpdatePackage)
	pkg.ID = "package-id"
	packageHandler.WithdrawPackageUsecase.PackageRepository.Create(pkg)

	token := generateToken("deliveryman-id")
	req := httptest.NewRequest("PATCH", "/packages/withdraw/package-id", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("packageId", "package-id")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	ctx := jwtauth.NewContext(req.Context(), token, nil)
	req = req.WithContext(ctx)
	w := httptest.NewRecorder()

	packageHandler.WithdrawPackage(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)

	var errorResponse Error
	err := json.NewDecoder(w.Body).Decode(&errorResponse)

	require.Nil(t, err)
	require.Equal(t, test.ErrOnUpdatePackage.Error(), errorResponse.Message)
}
