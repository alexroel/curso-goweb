package main

import (
	"dbmysql/db"
	"dbmysql/models"
	"fmt"
)

func main() {

	//Verificar conexion
	db.Connect()
	db.Ping()
	fmt.Println(db.ExistsTable("users"))

	//Clear tabla
	db.CreateTable(models.UserSchema, "users")
	//db.TruncateTable("users")
	//user := models.NewUser("roel", "roel123", "roel@gmail.com")
	//fmt.Println(user)
	//user.Username = "Alex"
	//user.Save()

	user := models.ListUsers()
	//user.Username = "roel"
	//user.Password = "roel"
	//user.Email = "roel@gmail.com"
	//user.Save()
	//user.Delete()
	fmt.Println(user)

	//Cerrar la conexion
	db.Close()

}
