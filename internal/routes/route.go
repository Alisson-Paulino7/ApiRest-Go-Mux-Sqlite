package routes

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	
	"github.com/alisson-paulino7/ApiRest-Go-Mux-Sqlite/api"
	// httpSwagger "github.com/swaggo/http-swagger"
)

func LoadRoutes() {
	router := mux.NewRouter()

	// router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler) - Tentativa falha
	router.HandleFunc("/", api.Home).Methods("GET")
	router.HandleFunc("/cadastro", api.CadastrarUser).Methods("POST")
	router.HandleFunc("/cadastro/", api.BuscarAllUser).Methods("GET")
	router.HandleFunc("/cadastro/{id}", api.BuscarOneUser).Methods("GET")
	router.HandleFunc("/cadastro/{id}", api.DeletarUser).Methods("DELETE")
	router.HandleFunc("/cadastro/{id}", api.AtualizarUser).Methods("PUT")
	router.HandleFunc("/totalUsers", api.TotalUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}