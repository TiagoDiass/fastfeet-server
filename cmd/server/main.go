package main

import (
	"encoding/json"
	"net/http"

	"github.com/TiagoDiass/fastfeet-server/internal/infra/web"
	"github.com/TiagoDiass/fastfeet-server/internal/test"
	recipientUsecase "github.com/TiagoDiass/fastfeet-server/internal/usecase/recipient"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// cfg, err := configs.LoadConfig(".")

	// if err != nil {
	// 	panic(err)
	// }

	recipientRepository := test.NewInMemoryRecipientRepository()
	createRecipientUsecase := recipientUsecase.NewCreateRecipientUsecase(recipientRepository)
	recipientHandler := web.NewRecipientHandler(createRecipientUsecase)

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

	http.ListenAndServe(":8000", router)
}
