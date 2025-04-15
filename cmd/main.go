package main

import (
	v1 "artyomkorchagin/web-shop/internal/handlers/v1"
	"artyomkorchagin/web-shop/internal/services/category"
	"artyomkorchagin/web-shop/internal/services/order"
	"artyomkorchagin/web-shop/internal/services/product"
	"artyomkorchagin/web-shop/internal/services/user"
	mssqlUser "artyomkorchagin/web-shop/internal/storage/mssql/user"
	psqlCategory "artyomkorchagin/web-shop/internal/storage/postgresql/category"
	psqlOrder "artyomkorchagin/web-shop/internal/storage/postgresql/order"
	psqlProduct "artyomkorchagin/web-shop/internal/storage/postgresql/product"
	psqlUser "artyomkorchagin/web-shop/internal/storage/postgresql/user"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/microsoft/go-mssqldb"
)

func main() {

	db, err := sql.Open("pgx", "host=yamabiko.proxy.rlwy.net port=33619 dbname=railway user=postgres password=AUJMTvvTqAPcrZyVLkpoBaUhXGkWAPBI sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to the database!")

	defer db.Close()

	handler := initHandler(db, "pgx")

	r := handler.InitRoutes()

	r.Run(":80")
}

func initHandler(db *sql.DB, driver string) *v1.Handler {
	var (
		userRepo     user.ReadWriter
		productRepo  product.ReadWriter
		categoryRepo category.ReadWriter
		orderRepo    order.ReadWriter
	)
	switch driver {
	case "pgx":
		{
			userRepo = psqlUser.NewRepository(db)
			productRepo = psqlProduct.NewRepository(db)
			categoryRepo = psqlCategory.NewRepository(db)
			orderRepo = psqlOrder.NewRepository(db)
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
	orderService := order.NewService(orderRepo)
	svc := v1.NewAllServices(userService, productService, categoryService, orderService)
	return v1.NewHandler(svc)
}
