/*
 * User Service
 *
 * This is simple client API
 *
 * API version: 1.0.0
 * Contact: schetinnikov@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/perevozov/arch-userservice/model"
	"github.com/perevozov/arch-userservice/service"
)

func main() {
	log.Printf("Server started")

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
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
