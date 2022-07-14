package models

import (
	"fmt"
	"genealog/bd"
	"genealog/utilities"
	"strconv"
)

//struct para devolver multiples indicios
type Indicio struct {
	Ind_entidad_movimiento string `json:"ind_entidad_movimiento"`
	Ind_fuente             string `json:"ind_fuente"`
	Ind_id                 string `json:"ind_id"`
	Ind_objeto_movimiento  string `json:"ind_objeto_movimiento"`
	Ind_orden              string `json:"ind_orden"`
	Ind_periodo            string `json:"ind_periodo"`
	Ind_tipo_movimiento    string `json:"ind_tipo_movimiento"`
	Ind_anexos             string `json:"ind_anexos"`
}

//struct para devolver un solo indicio
type IndicioI struct {
	Ind_entidad            string `json:"ind_entidad"`
	Ind_entidad_movimiento string `json:"ind_entidad_movimiento"`
	Ind_fuente             string `json:"ind_fuente"`
	Ind_id                 string `json:"ind_id"`
	Ind_objeto_movimiento  string `json:"ind_objeto_movimiento"`
	Ind_orden              string `json:"ind_orden"`
	Ind_periodo            string `json:"ind_periodo"`
	Ind_tipo_movimiento    string `json:"ind_tipo_movimiento"`
}
type IndicioPost struct {
	Id        int64
	Entidades []string
	SecGen    int64
	EntMov    string
	TipoMov   string
	Periodo   string
	ObjMov    string
	Fuente    string
}

type Indicios []interface{}

//`json:"ind_anexos,omitempty"`  controla campos vacios

func GetIndiciosModel(id string) (Indicios, error) {

	rows, err := bd.QueryConsulta("SELECT ind_entidad_movimiento,ind_fuente, ind_id,ind_objeto_movimiento ,ind_orden,ind_periodo,ind_tipo_movimiento,ind_anexos from indicios inner join rel_entidad_indicio on ind_id=rei_id_indicio where rei_id_entidad=" + id + " order by ind_periodo,ind_orden")
	indicios := Indicios{}
	for rows.Next() {
		indicio := Indicio{}

		rows.Scan(&indicio.Ind_entidad_movimiento, &indicio.Ind_fuente, &indicio.Ind_id, &indicio.Ind_objeto_movimiento, &indicio.Ind_orden, &indicio.Ind_periodo, &indicio.Ind_tipo_movimiento, &indicio.Ind_anexos)
		indicios = append(indicios, indicio)

	}
	return indicios, err
}

func GetIndicioModel(id string) (Indicios, error) {
	rows, err := bd.QueryConsulta("select GROUP_CONCAT(rei_id_entidad SEPARATOR ';') as ind_entidad, ind_entidad_movimiento,ind_fuente,  ind_id,ind_objeto_movimiento,ind_orden, ind_periodo,ind_tipo_movimiento from indicios inner join rel_entidad_indicio on ind_id=rei_id_indicio where rei_id_indicio=" + id + " order by ind_periodo,ind_orden")
	indicios := Indicios{}

	for rows.Next() {
		indicioI := IndicioI{}
		rows.Scan(&indicioI.Ind_entidad, &indicioI.Ind_entidad_movimiento, &indicioI.Ind_fuente, &indicioI.Ind_id, &indicioI.Ind_objeto_movimiento, &indicioI.Ind_orden, &indicioI.Ind_periodo, &indicioI.Ind_tipo_movimiento)
		indicios = append(indicios, indicioI)
	}
	return indicios, err
}
func DeleteIndicioModel(id string) (int64, error, error) {
	resultado, err := bd.ExecDelete("indicios", "ind_id="+id)
	resultado2, err2 := bd.ExecDelete("rel_entidad_indicio", "rei_id_indicio="+id)
	return resultado + resultado2, err, err2
}
func PostIndicioModel(objetoIndicio IndicioPost) {

	valores := strconv.Itoa(int(objetoIndicio.SecGen)) + ",\"" + objetoIndicio.EntMov + "\"" + ",\"" + objetoIndicio.TipoMov + "\"" + ",\"" + objetoIndicio.Periodo + "\"" + ",\"" + objetoIndicio.ObjMov + "\"" + ",\"" + objetoIndicio.Fuente + "\""

	idRegistroNuevo, err := bd.ExecInserta2("indicios", "ind_sec_gen,ind_entidad_movimiento,ind_tipo_movimiento,ind_periodo,ind_objeto_movimiento,ind_fuente", valores)
	//obtener id y asociarlo con el arreglo entidades package utilities
	cadenaEntInd := utilities.GeneraRelEntidadIndicioPares(objetoIndicio.Entidades, idRegistroNuevo)

	fmt.Println(err)
	fmt.Println(cadenaEntInd)
}
