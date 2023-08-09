package main

import (
	"fmt"
	"react_go_echo/db"
	"react_go_echo/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)

	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
