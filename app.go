package main

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/rhperera/go-base/config"
	"github.com/rhperera/go-base/server"
	"github.com/rhperera/go-base/user"
)

func main()  {
	config.Init()
	server.Init()
	server.InitAPI()

	db, _ := sql.Open("mssql","")

	userMSSqlRepo := user.NewMSSqlRepo(db)
	userService := user.NewService(userMSSqlRepo)
	user.NewHandler(userService)

	server.Connect("8080")
}


