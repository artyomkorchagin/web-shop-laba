package main

import (
	"artyomkorchagin/web-shop/internal/app/user"
	"artyomkorchagin/web-shop/internal/config"
	v1 "artyomkorchagin/web-shop/internal/handlers/v1"
	mssqlUser "artyomkorchagin/web-shop/internal/mssql/user"
	psqlUser "artyomkorchagin/web-shop/internal/postgresql/user"
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/microsoft/go-mssqldb"
)

func main() {

	var configFlag = flag.String("dbms", "psql", "name of DBMS (e.g psql, mssql etc.)")

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

	handler := initHandler(db)

	r := handler.InitRoutes()

	r.Run(":3000")
}

func initHandler(db *sql.DB) *v1.Handler {
	var userRepo user.ReadWriter
	switch fmt.Sprintf("%v", db.Driver()) {
	case "pgx":
		{
			userRepo = psqlUser.NewUserRepository(db)
		}
	case "mssql":
		{
			userRepo = mssqlUser.NewUserRepository(db)
		}
	default:
		{
			log.Fatal("unknown driver", fmt.Sprintf("%v", db.Driver()))
			return &v1.Handler{}
		}
	}
	userService := user.NewService(userRepo)
	svc := v1.NewAllSercivces(userService, nil)
	return v1.NewHandler(svc)
}
