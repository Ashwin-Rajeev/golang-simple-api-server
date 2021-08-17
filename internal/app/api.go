package app

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// statusHandler is the :GET API for get DB status.
func (s *Server) statusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := s.DBGetStatus()
	if err != nil {
		send(w, http.StatusInternalServerError, nil, err)
		return
	}
	send(w, http.StatusOK, status, nil)
}

// addCategoryHandler is the :POST API for add category.
func (s *Server) addCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		send(w, http.StatusBadRequest, nil, err)
		return
	}
	id, err := s.DBAddCategory(category)
	if err != nil {
		send(w, http.StatusInternalServerError, nil, err)
		return
	}
	send(w, http.StatusCreated, id, nil)
}

// addProductHandler is the :POST API for add product.
func (s *Server) addProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		send(w, http.StatusBadRequest, nil, err)
		return
	}
	id, err := s.DBAddProduct(product)
	if err != nil {
		send(w, http.StatusInternalServerError, nil, err)
		return
	}
	send(w, http.StatusCreated, id, nil)
}

// updateProductHandler is the :PUT API for update product.
func (s *Server) updateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		send(w, http.StatusBadRequest, nil, err)
		return
	}
	id, err := s.DBUpdateProduct(product)
	if err != nil {
		send(w, http.StatusInternalServerError, nil, err)
		return
	}
	send(w, http.StatusCreated, id, nil)
}

// getProductByCategoryHandler is the :GET API for get product.
func (s *Server) getProductByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "category")
	products, err := s.DBGetProductByCategory(categoryID)
	if err != nil {
		send(w, http.StatusInternalServerError, nil, err)
		return
	}
	send(w, http.StatusCreated, products, nil)
}

// getAllCategoryHandler is the :GET API for get all categories.
func (s *Server) getAllCategoryHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := s.DBGetAllCategory()
	if err != nil {
		send(w, http.StatusInternalServerError, nil, err)
		return
	}
	send(w, http.StatusOK, categories, nil)
}
