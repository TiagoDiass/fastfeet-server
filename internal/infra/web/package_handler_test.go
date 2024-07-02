package web

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/TiagoDiass/fastfeet-server/internal/repository/test"
// 	"github.com/TiagoDiass/fastfeet-server/internal/usecase/package"
// 	"github.com/stretchr/testify/require"
// )

// func TestListAvailablePackagesHandler(t *testing.T) {
// 	// Setup
// 	packageRepo := test.NewInMemoryPackageRepository()
// 	listAvailablePackagesUsecase := usecase.NewListAvailablePackagesUsecase(packageRepo)
// 	packageHandler := NewPackageHandler(nil, listAvailablePackagesUsecase)

// 	// Add necessary data to repositories
// 	packageRepo.Create(&entity.Package{
// 		ID:            "pkg1",
// 		RecipientId:   "rec1",
// 		DeliverymanId: "del1",
// 		Name:          "Package 1",
// 		Status:        "WAITING_WITHDRAW",
// 	})

// 	// Create request
// 	req := httptest.NewRequest("GET", "/packages/available", nil)
// 	w := httptest.NewRecorder()

// 	// Call handler
// 	packageHandler.ListAvailablePackages(w, req)

// 	// Check response
// 	require.Equal(t, http.StatusOK, w.Code)

// 	var output []usecase.ListAvailablePackagesOutputDTO
// 	err := json.NewDecoder(w.Body).Decode(&output)
// 	require.Nil(t, err)
// 	require.Len(t, output, 1)
// 	require.Equal(t, "pkg1", output[0].ID)
// 	require.Equal(t, "Package 1", output[0].Name)
// 	require.Equal(t, "WAITING_WITHDRAW", output[0].Status)
// }
