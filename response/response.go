package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/* Respuesta en formato json al usuario*/
func RespuestaUsuario(res http.ResponseWriter, datosUsuario interface{}) {
	res.Header().Set("Content-Type", "application/json")

	output, _ := json.Marshal(datosUsuario)
	fmt.Fprintln(res, string(output))
}
