package routes

import (
	"pet-dex-backend/v2/api/controllers"

	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/api/ongs", controllers.HandleCreateOng)

	return router
}
