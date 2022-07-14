package bd

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var BD *sql.DB

//Abrir y cerrar conexion
func Conectar() {
	url := "root:123@tcp(localhost:3306)/genealogias"
	conexion, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	} else {
		BD = conexion
	}
}
func Close() {
	BD.Close()
}

//Funcion ver tabla
func QueryVerTabla(tabla string, campos string, condicion string) (*sql.Rows, error) {
	sql := "SELECT " + campos + " FROM " + tabla + " WHERE " + condicion
	Conectar()
	rows, err := BD.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	Close()
	return rows, err

}

//Funcion ver SQL
func QueryConsulta(consulta string) (*sql.Rows, error) {
	Conectar()
	rows, err := BD.Query(consulta)
	if err != nil {
		fmt.Println(err)
	}
	Close()
	return rows, err

}

//Funcion inserta
func ExecInserta(tabla string, campos string, valores string) (int64, error) {
	Conectar()
	sql := "INSERT INTO " + tabla + " (" + campos + ") VALUES(" + valores + ")"
	resultado, err := BD.Exec(sql)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	Close()
	return resultado.RowsAffected()
}

//Funcion actualiza
func ExecActualiza(tabla string, valores string, condicion string) (int64, error) {
	Conectar()
	sql := "UPDATE " + tabla + " SET " + valores + " WHERE " + condicion
	resultado, err := BD.Exec(sql)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	Close()
	return resultado.RowsAffected()
}

//Funcion inserta2
func ExecInserta2(tabla string, campos string, valores string) (int64, error) {
	Conectar()
	sql := "INSERT INTO " + tabla + " (" + campos + ") VALUES(" + valores + ")"
	fmt.Println(sql)
	resultado, err := BD.Exec(sql)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	Close()
	return resultado.LastInsertId()
}

//Funcion borrar
func ExecDelete(tabla string, condicion string) (int64, error) {
	sql := "DELETE FROM " + tabla + " WHERE " + condicion
	Conectar()
	resultado, err := BD.Exec(sql)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	Close()

	return resultado.RowsAffected()
}
