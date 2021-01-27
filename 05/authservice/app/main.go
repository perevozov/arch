package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/perevozov/arch/05/authservice/model"
	"github.com/perevozov/arch/05/authservice/service"
)

func main() {
	log.Printf("Server started")

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?loc=Local",
		service.ServiceConfig.DBUser,
		service.ServiceConfig.DBPasswd,
		service.ServiceConfig.DBHost,
		service.ServiceConfig.DBName,
	)
	db, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	log.Println("Connected to database")

	env := service.Env{DB: &model.DB{db}}

	router := service.NewRouter(&env)

	log.Fatal(http.ListenAndServe(":"+service.ServiceConfig.ListenPort, router))
}
