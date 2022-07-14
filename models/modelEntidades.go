package models

import (
	"fmt"
	"genealog/bd"
	"time"
)

type Entidad struct {
	Ent_id       string `json:"ent_id"`
	Ent_nombre   string `json:"ent_nombre"`
	Ent_vigencia string `json:"ent_vigencia"`
	Indicios     string `json:"indicios"`
}

type Entidades []Entidad

//Consulta completa=>"SELECT ent_id,ent_nombre,ent_vigencia,count(rei_id_indicio) as indicios from entidades e left join rel_entidad_indicio on ent_id=rei_id_entidad group by ent_id order by ent_nombre "

func GetEntidadesModel() (Entidades, error) {
	entidades := Entidades{}
	inicioTiempo := time.Now()
	//Consulta con vista
	rows, err := bd.QueryConsulta("select ent_id,ent_nombre,ent_vigencia,indicios from entidadesv ")
	tiempoTranscurrido := time.Since(inicioTiempo)
	fmt.Println(tiempoTranscurrido)
	for rows.Next() {
		entidad := Entidad{}

		rows.Scan(&entidad.Ent_id, &entidad.Ent_nombre, &entidad.Ent_vigencia, &entidad.Indicios)
		entidades = append(entidades, entidad)

	}
	return entidades, err
}
