package main

import (
	"artyomkorchagin/web-shop/internal/config"
	v1 "artyomkorchagin/web-shop/internal/handlers/v1"
	"artyomkorchagin/web-shop/internal/services/category"
	"artyomkorchagin/web-shop/internal/services/product"
	"artyomkorchagin/web-shop/internal/services/user"
	mssqlUser "artyomkorchagin/web-shop/internal/storage/mssql/user"
	psqlCategory "artyomkorchagin/web-shop/internal/storage/postgresql/category"
	psqlProduct "artyomkorchagin/web-shop/internal/storage/postgresql/product"
	psqlUser "artyomkorchagin/web-shop/internal/storage/postgresql/user"
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

	handler := initHandler(db, cfg.Driver)

	r := handler.InitRoutes()

	r.Run(":3000")
}

func initHandler(db *sql.DB, driver string) *v1.Handler {
	var (
		userRepo     user.ReadWriter
		productRepo  product.ReadWriter
		categoryRepo category.ReadWriter
	)
	switch driver {
	case "pgx":
		{
			userRepo = psqlUser.NewRepository(db)
			productRepo = psqlProduct.NewRepository(db)
			categoryRepo = psqlCategory.NewRepository(db)
		}
	case "mssql":
		{
			userRepo = mssqlUser.NewRepository(db)
			productRepo = nil
			categoryRepo = nil
		}
	default:
		{
			log.Fatal("unknown driver in config file", fmt.Sprintf("%v", db.Driver()))
			return &v1.Handler{}
		}
	}
	userService := user.NewService(userRepo)
	productService := product.NewService(productRepo)
	categoryService := category.NewService(categoryRepo)
	svc := v1.NewAllServices(userService, productService, categoryService)
	return v1.NewHandler(svc)
}
