package main

import (
	"artyomkorchagin/web-shop/internal/app/user"
	v1 "artyomkorchagin/web-shop/internal/handlers/v1"
	mssqlUser "artyomkorchagin/web-shop/internal/mssql/user"
	"database/sql"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {

	// cfg, err := config.Load("../config/config.yaml")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	db, err := sql.Open("mssql", "server=localhost;user id=admin;database=social;password=123412341234")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to the database!")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := mssqlUser.NewUserRepository(db)

	userService := user.NewService(userRepo)

	svc := v1.NewAllSercivces(userService, nil)

	handler := v1.NewHandler(svc)

	r := handler.InitRoutes()

	r.Run(":3000")
}
