package main

import (
	"net/http"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/infra/web"
	"github.com/TiagoDiass/fastfeet-server/internal/repository"
	repositoryimpl "github.com/TiagoDiass/fastfeet-server/internal/repository/repository_impl"
	packageUsecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/package"
	recipientUsecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/recipient"
	userUsecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repositories struct {
	RecipientRepository repository.RecipientRepository
	UserRepository      repository.UserRepository
	PackageRepository   repository.PackageRepository
}

type Usecases struct {
	CreateUserUsecase              *userUsecase.CreateUserUsecase
	CreateRecipientUsecase         *recipientUsecase.CreateRecipientUsecase
	CreatePackageUsecase           *packageUsecase.CreatePackageUsecase
	ListAvailablePackagesUsecase   *packageUsecase.ListAvailablePackagesUsecase
	ListDeliveredPackagesUsecase   *packageUsecase.ListDeliveredPackagesUsecase
	WithdrawPackageUsecase         *packageUsecase.WithdrawPackageUsecase
	ConfirmDeliveredPackageUsecase *packageUsecase.ConfirmDeliveredPackageUsecase
}

type Handlers struct {
	RecipientHandler *web.RecipientHandler
	UserHandler      *web.UserHandler
	PackageHandler   *web.PackageHandler
}

func main() {
	// cfg, err := configs.LoadConfig(".")

	// if err != nil {
	// 	panic(err)
	// }

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Recipient{}, &entity.User{}, &entity.Package{})

	repositories := createRepositories(db)
	usecases := createUsecases(repositories)
	handlers := createHandlers(usecases)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	initializeRoutes(router, handlers)

	http.ListenAndServe(":8000", router)
}

func createRepositories(db *gorm.DB) *Repositories {
	recipientRepository := repositoryimpl.NewGormRecipientRepository(db)
	userRepository := repositoryimpl.NewGormUserRepository(db)
	packageRepository := repositoryimpl.NewGormPackageRepository(db)

	return &Repositories{
		RecipientRepository: recipientRepository,
		UserRepository:      userRepository,
		PackageRepository:   packageRepository,
	}
}

func createUsecases(repositories *Repositories) *Usecases {
	createUserUsecase := userUsecase.NewCreateUserUsecase(repositories.UserRepository)
	createRecipientUsecase := recipientUsecase.NewCreateRecipientUsecase(repositories.RecipientRepository)
	createPackageUsecase := packageUsecase.NewCreatePackageUsecase(
		repositories.PackageRepository,
		repositories.UserRepository,
		repositories.RecipientRepository,
	)
	listAvailablePackagesUsecase := packageUsecase.NewListAvailablePackagesUsecase(repositories.PackageRepository)
	listDeliveredPackagesUsecase := packageUsecase.NewListDeliveredPackagesUsecase(repositories.PackageRepository)
	withdrawPackageUsecase := packageUsecase.NewWithdrawPackageUsecase(repositories.PackageRepository, repositories.UserRepository)
	confirmDeliveredPackageUsecase := packageUsecase.NewConfirmDeliveredPackageUsecase(repositories.PackageRepository)

	return &Usecases{
		CreateUserUsecase:              createUserUsecase,
		CreateRecipientUsecase:         createRecipientUsecase,
		CreatePackageUsecase:           createPackageUsecase,
		ListAvailablePackagesUsecase:   listAvailablePackagesUsecase,
		ListDeliveredPackagesUsecase:   listDeliveredPackagesUsecase,
		WithdrawPackageUsecase:         withdrawPackageUsecase,
		ConfirmDeliveredPackageUsecase: confirmDeliveredPackageUsecase,
	}
}

func createHandlers(usecases *Usecases) *Handlers {
	recipientHandler := web.NewRecipientHandler(usecases.CreateRecipientUsecase)
	userHandler := web.NewUserHandler(usecases.CreateUserUsecase)
	packageHandler := web.NewPackageHandler(
		usecases.CreatePackageUsecase,
		usecases.ListAvailablePackagesUsecase,
		usecases.ListDeliveredPackagesUsecase,
		usecases.WithdrawPackageUsecase,
		usecases.ConfirmDeliveredPackageUsecase,
	)

	return &Handlers{
		RecipientHandler: recipientHandler,
		UserHandler:      userHandler,
		PackageHandler:   packageHandler,
	}
}

func initializeRoutes(router *chi.Mux, handlers *Handlers) {
	router.Post("/recipients", handlers.RecipientHandler.CreateRecipient)
	router.Post("/users", handlers.UserHandler.CreateUser)
	router.Post("/packages", handlers.PackageHandler.CreatePackage)
	router.Get("/packages/available", handlers.PackageHandler.ListAvailablePackages)
	router.Get("/packages/delivered", handlers.PackageHandler.ListDeliveredPackages)
	router.Patch("/packages/withdraw/{packageId}", handlers.PackageHandler.WithdrawPackage)
	router.Patch("/packages/confirm-delivery/{packageId}", handlers.PackageHandler.ConfirmDeliveredPackage)
}
