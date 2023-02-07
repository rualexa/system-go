package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func conexionDB() (*sql.DB, error) {
	// Create the database handle, confirm driver is present
	Driver := "mysql"
	Usuario := "root"
	Password := "rbn208**"
	namedb := "sistema"

	conexion, err := sql.Open(Driver, Usuario+":"+Password+"@tcp(localhost:3306)/"+namedb)
	if err != nil {
		return nil, err
	}
	err = conexion.Ping()

	if err != nil {
		return nil, err
	}

	return conexion, nil

}
func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)
	log.Println("servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request) {

	conexionEstrablecidad, _ := conexionDB()

	insertarRegistro, err := conexionEstrablecidad.Prepare("INSERT INTO empleados(nombre, correo) VALUES('Rubens Rangel','rualexander12@gmail.com')")

	if err != nil {
		panic(err.Error())
	}

	insertarRegistro.Exec()

	conexionEstrablecidad.Close()

	//fmt.Fprintf(w, "hello world")
	plantillas.ExecuteTemplate(w, "inicio", nil)
}

func Crear(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hello world")
	plantillas.ExecuteTemplate(w, "crear", nil)
}
