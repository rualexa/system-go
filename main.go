package main

import (
	"database/sql"
	"fmt"
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

type Empleado struct {
	Id     int
	Nombre string
	Correo string
}

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insertar", Insertar)
	http.HandleFunc("/borrar", Borrar)
	http.HandleFunc("/editar", Editar)

	log.Println("servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request) {

	conexionEstrablecidad, _ := conexionDB()

	Registros, err := conexionEstrablecidad.Query("SELECT * FROM empleados")

	if err != nil {
		panic(err.Error())
	}

	empleado := Empleado{}
	listaEmpleados := []Empleado{}

	for Registros.Next() {
		var id int
		var correo, nombre string
		err = Registros.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Correo = correo
		empleado.Nombre = nombre
		listaEmpleados = append(listaEmpleados, empleado)
	}
	fmt.Println(listaEmpleados)

	conexionEstrablecidad.Close()

	//fmt.Fprintf(w, "hello world")
	plantillas.ExecuteTemplate(w, "inicio", listaEmpleados)
}

func Crear(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hello world")
	plantillas.ExecuteTemplate(w, "crear", nil)
}

func Insertar(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		name := r.FormValue("nombre")
		email := r.FormValue("correo")

		conexionEstrablecidad, _ := conexionDB()

		insertarRegistro, err := conexionEstrablecidad.Prepare("INSERT INTO empleados(nombre, correo) VALUES(?,?)")

		if err != nil {
			panic(err.Error())
		}

		insertarRegistro.Exec(name, email)
		conexionEstrablecidad.Close()
	}

	http.Redirect(w, r, "/", 301)

}

func Borrar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	fmt.Println(idEmpleado)

	conexionEstrablecidad, _ := conexionDB()

	borrarRegistro, err := conexionEstrablecidad.Prepare("DELETE FROM empleados WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	borrarRegistro.Exec(idEmpleado)
	conexionEstrablecidad.Close()

	http.Redirect(w, r, "/", 301)
}

func Editar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	if r.Method == "POST" {

		conexionEstrablecidad, _ := conexionDB()

		Registro, err := conexionEstrablecidad.Query("SELECT * FROM empleados WHERE id=?", idEmpleado)

		if err != nil {
			panic(err.Error())
		}

		empleado := Empleado{}

		for Registro.Next() {
			var id int
			var correo, nombre string
			err = Registro.Scan(&id, &nombre, &correo)
			if err != nil {
				panic(err.Error())
			}
			empleado.Id = id
			empleado.Correo = correo
			empleado.Nombre = nombre
		}

		conexionEstrablecidad.Close()
	}

	http.Redirect(w, r, "/", 301)
}
