package api

import (
	"database/sql"
	"log"
	"net/http"
	"tutorial/tiago-complete-backend/service/product"
	"tutorial/tiago-complete-backend/service/user"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: address,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	r := mux.NewRouter()
	subrouter := r.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, r)
}
