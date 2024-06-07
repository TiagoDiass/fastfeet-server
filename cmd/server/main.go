package main

import (
	"net/http"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"github.com/TiagoDiass/fastfeet-server/internal/infra/web"
	repositoryimpl "github.com/TiagoDiass/fastfeet-server/internal/repository/repository_impl"
	packageUsecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/package"
	recipientUsecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/recipient"
	userUsecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

	recipientRepository := repositoryimpl.NewGormRecipientRepository(db)
	userRepository := repositoryimpl.NewGormUserRepository(db)
	packageRepository := repositoryimpl.NewGormPackageRepository(db)

	createUserUsecase := userUsecase.NewCreateUserUsecase(userRepository)
	createRecipientUsecase := recipientUsecase.NewCreateRecipientUsecase(recipientRepository)
	createPackageUsecase := packageUsecase.NewCreatePackageUsecase(
		packageRepository,
		userRepository,
		recipientRepository,
	)
	listAvailablePackagesUsecase := packageUsecase.NewListAvailablePackagesUsecase(packageRepository)

	recipientHandler := web.NewRecipientHandler(createRecipientUsecase)
	userHandler := web.NewUserHandler(createUserUsecase)
	packageHandler := web.NewPackageHandler(createPackageUsecase, listAvailablePackagesUsecase)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/recipients", recipientHandler.CreateRecipient)
	router.Post("/users", userHandler.CreateUser)
	router.Post("/packages", packageHandler.CreatePackage)
	router.Get("/packages/available", packageHandler.ListAvailablePackages)

	http.ListenAndServe(":8000", router)
}
