package handler

import (
	"encoding/json"
	"fmt"
	"genealog/models"
	"genealog/response"
	"net/http"

	"github.com/gorilla/mux"
)

func GetEntidades(res http.ResponseWriter, req *http.Request) {

	if entidades, err := models.GetEntidadesModel(); err != nil {
		fmt.Println(err)

	} else {
		response.RespuestaUsuario(res, entidades)
	}

}
func GetIndicios(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	if indicios, err := models.GetIndiciosModel(id); err != nil {
		fmt.Println(err)

	} else {
		response.RespuestaUsuario(res, indicios)
	}

}
func GetIndicio(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	if indicio, err := models.GetIndicioModel(id); err != nil {
		fmt.Println(err)
	} else {
		response.RespuestaUsuario(res, indicio)
	}

}
func DeleteIndicio(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	res.Header().Set("Content-Type", "application/json")
	if result, err, err2 := models.DeleteIndicioModel(id); err != nil || err2 != nil {
		fmt.Println(err)
		fmt.Println(err2)

	} else {

		fmt.Fprintln(res, result)

	}
}
func PostIndicio(res http.ResponseWriter, req *http.Request) {

	indicio := models.IndicioPost{}

	err := json.NewDecoder(req.Body).Decode(&indicio)
	if err != nil {
		json.NewEncoder(res).Encode("Cuerpo de petición no válido ejem:{int,[string],int,string,string,string,strig,string}")
		return
	}
	models.PostIndicioModel(indicio)

}
func PutIndicio(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Actualizando registro")
}
