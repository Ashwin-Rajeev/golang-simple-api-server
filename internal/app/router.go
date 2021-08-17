package app

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// InitRouter initialize a new chi router instance.
func (s *Server) InitRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Basic set of handler routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/status", s.statusHandler)

		r.Post("/category", s.addCategoryHandler)
		r.Get("/categories", s.getAllCategoryHandler)
		
		r.Post("/product", s.addProductHandler)
		r.Put("/product", s.updateProductHandler)
		r.Get("/products/{category}", s.getProductByCategoryHandler)
	})

	return r
}
