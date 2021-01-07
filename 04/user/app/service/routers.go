/*
 * User Service
 *
 * This is simple client API
 *
 * API version: 1.0.0
 * Contact: schetinnikov@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package service

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gorilla/mux"
	"github.com/perevozov/arch-userservice/model"
)

type Env struct {
	DB model.Datastore
}

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.Handler
}

type Routes []Route

func NewRouter(env *Env) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	var routes = Routes{
		Route{
			"Index",
			"GET",
			"/api/v1/",
			http.HandlerFunc(Index),
		},

		Route{
			"CreateUser",
			"POST",
			"/api/v1/user",
			http.HandlerFunc(env.CreateUser),
		},

		Route{
			"DeleteUser",
			"DELETE",
			"/api/v1/user/{userId}",
			http.HandlerFunc(env.DeleteUser),
		},

		Route{
			"FindUserById",
			"GET",
			"/api/v1/user/{userId}",
			http.HandlerFunc(env.FindUserById),
		},

		Route{
			"UpdateUser",
			"PUT",
			"/api/v1/user/{userId}",
			http.HandlerFunc(env.UpdateUser),
		},

		Route{
			"Config",
			"GET",
			"/config",
			http.HandlerFunc(GetConfig),
		},

		Route{
			"Status",
			"GET",
			"/status",
			http.HandlerFunc(env.GetStatus),
		},

		Route{
			"Metrics",
			"GET",
			"/metrics",
			promhttp.Handler(),
		},
	}
	for _, route := range routes {
		var handler http.Handler
		handler = route.Handler
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	router.PathPrefix("/").Handler(Logger(http.HandlerFunc(Default), "unknown"))
	return router
}

func Default(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not found", 404)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprintf(w, `{"dbUser":"%s", "dbHost":"%s", "dbName":"%s"}`, ServiceConfig.DBUser, ServiceConfig.DBHost, ServiceConfig.DBName)
}

func (env Env) GetStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(w, `{"status":"OK"}`)
}
