package main

import (
	"encoding/json"
	"net/http"

	"github.com/TiagoDiass/fastfeet-server/internal/infra/web"
	"github.com/TiagoDiass/fastfeet-server/internal/test"
	recipientUsecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/recipient"
	userUsecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// cfg, err := configs.LoadConfig(".")

	// if err != nil {
	// 	panic(err)
	// }

	recipientRepository := test.NewInMemoryRecipientRepository()
	userRepository := test.NewInMemoryUserRepository()

	createUserUsecase := userUsecase.NewCreateUserUsecase(userRepository)
	createRecipientUsecase := recipientUsecase.NewCreateRecipientUsecase(recipientRepository)

	recipientHandler := web.NewRecipientHandler(createRecipientUsecase)
	userHandler := web.NewUserHandler(createUserUsecase)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/recipients", func(w http.ResponseWriter, r *http.Request) {
		recipients, _ := recipientRepository.FindAll()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(recipients)
		return
	})
	router.Post("/recipients", recipientHandler.CreateRecipient)

	router.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		users, _ := userRepository.FindAll()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
		return
	})
	router.Post("/users", userHandler.CreateUser)

	http.ListenAndServe(":8000", router)
}
