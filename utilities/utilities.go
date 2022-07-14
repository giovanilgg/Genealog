package utilities

import (
	"strconv"
	"strings"
)

func GeneraRelEntidadIndicioPares(arrayEntidades []string, id int64) string {
	var entidadesRelacion string
	for i := 0; i < len(arrayEntidades); i++ {
		valores := "(" + "\"" + strconv.Itoa(int(id)) + "\"" + "," + "\"" + arrayEntidades[i] + "\"" + ")"
		entidadesRelacion += "," + valores
	}
	salida := strings.TrimLeft(entidadesRelacion, ",")
	return salida
}
