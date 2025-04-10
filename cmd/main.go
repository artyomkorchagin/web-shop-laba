package main

import (
	"database/sql"
	"log"
	"socialsecurity/internal/app/application"
	"socialsecurity/internal/app/user"
	v1 "socialsecurity/internal/handlers/v1"
	mssqlApplication "socialsecurity/internal/mssql/application"
	mssqlUser "socialsecurity/internal/mssql/user"

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

	applicationRepo := mssqlApplication.NewApplicationRepository(db)

	applicationService := application.NewService(applicationRepo)

	svc := v1.NewAllSercivces(userService, applicationService)

	handler := v1.NewHandler(svc)

	r := handler.InitRoutes()

	r.Run(":3000")
}
