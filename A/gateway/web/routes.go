package web

import (
	"gateway/web/controllers/order"
	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {
	r.Route("/api", func(r chi.Router) {

		// Orders
		r.Route("/orders", func(r chi.Router) {
			ordersController := order.NewOrder()
			r.Post("/", ordersController.Create)
		})

	})
}
