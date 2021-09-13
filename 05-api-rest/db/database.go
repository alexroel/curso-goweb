package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(host:port)/database?charset=utf8
const url = "root:123456@tcp(localhost:3306)/goweb_db"

//Guarda la conexion
var db *sql.DB

//Realizar lac conexion
func Connect() {
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion exitosa")
	db = conection
}

//Cerrar la Conexion
func Close() {
	db.Close()
}

//Verificar la conexion
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//Ferificar si una tabla existe o no
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows.Next()
}

//Crear una tabla en la base de datos
func CreateTable(schema, name string) {

	if !ExistsTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}

//Eliminara Tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	db.Exec(sql)
}

//Polimorfismo a Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

//Polimorfismo a Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}

	return rows, nil
}
