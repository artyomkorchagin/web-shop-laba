package main

import (
	"artyomkorchagin/web-shop/internal/app/user"
	"artyomkorchagin/web-shop/internal/config"
	v1 "artyomkorchagin/web-shop/internal/handlers/v1"
	mssqlUser "artyomkorchagin/web-shop/internal/mssql/user"
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/microsoft/go-mssqldb"
)

func main() {

	var configFlag = flag.String("dbms", "mssql", "name of DBMS (e.g pgsql, mssql etc.)")

	flag.Parse()

	cfg, err := config.Load(fmt.Sprintf("../config/%s.yaml", *configFlag))
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open(cfg.Driver, cfg.GetDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to the database!")

	defer db.Close()

	fmt.Println(db.Driver())

	userRepo := mssqlUser.NewUserRepository(db)

	userService := user.NewService(userRepo)

	svc := v1.NewAllSercivces(userService, nil)

	handler := v1.NewHandler(svc)

	r := handler.InitRoutes()

	r.Run(":3000")
}

func initServices(db sql.DB) {

}
