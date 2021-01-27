package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/perevozov/arch/05/authservice/model"
	"github.com/perevozov/arch/05/authservice/service"
)

func main() {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		"bitrix",    //service.ServiceConfig.DBUser,
		"bitrix",    //service.ServiceConfig.DBPasswd,
		"localhost", //service.ServiceConfig.DBHost,
		"hw05",      //service.ServiceConfig.DBName,
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

	/* u := &model.User{
		Username:  str("user01"),
		FirstName: str("Alex"),
		LastName:  str("Smirnov"),
		Email:     str("user01@mail.me"),
	}
	_, err = env.DB.AddUser(u)
	if err != nil {
		panic(err)
	}

	err = env.DB.SetUserPassword(u, "123456")
	if err != nil {
		panic(err)
	} */

	user, err := env.DB.CheckUserPassword("user01", "1234567")
	if err != nil {
		panic(err)
	}

	fmt.Printf("U: %v\n", user)
}

func str(s string) *string {
	return &s
}
