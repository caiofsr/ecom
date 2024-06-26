package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/caiofsr/ecom/service/cart"
	"github.com/caiofsr/ecom/service/order"
	"github.com/caiofsr/ecom/service/product"
	"github.com/caiofsr/ecom/service/user"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	orderStore := order.NewStore(s.db)
	productStore := product.NewStore(s.db)

	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	log.Println("Listening on ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
