package routes

import (
	ongcontroller "pet-dex-backend/v2/api/controllers/ong"

	"github.com/go-chi/chi"
)

func InitRouter(c *chi.Mux) {
	c.Route("/api", func(r chi.Router) {
		r.Route("/pet", func(r chi.Router) {
		})
		r.Route("/ong", func(r chi.Router) {
			r.Post("/", ongcontroller.CreateOng)
		})
	})

}
