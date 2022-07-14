package rutas

import (
	"genealog/handler"

	"github.com/gorilla/mux"
)

func Routing() *mux.Router {

	Enrutador := mux.NewRouter()

	Enrutador.HandleFunc("/entidades", handler.GetEntidades).Methods("GET")
	Enrutador.HandleFunc("/indicios/{id:[0-9]+}", handler.GetIndicios).Methods("GET")
	Enrutador.HandleFunc("/indicio/{id:[0-9]+}", handler.GetIndicio).Methods("GET")
	Enrutador.HandleFunc("/indicio/{id:[0-9]+}", handler.DeleteIndicio).Methods("DELETE")
	Enrutador.HandleFunc("/indicio", handler.PostIndicio).Methods("POST")
	Enrutador.HandleFunc("/indicio", handler.PutIndicio).Methods("PUT")
	return Enrutador
}
