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
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/perevozov/arch-userservice/model"

	"github.com/gorilla/mux"
)

func writeJsonResponse(w http.ResponseWriter, s interface{}) error {
	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	jsonPayload, err := json.Marshal(s)
	if err != nil {
		return err
	}
	w.Write(jsonPayload)
	return nil
}

func (env *Env) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	emailCheckClient := NewEmailCheckerClient(&url.URL{Scheme: "http", Host: ServiceConfig.EmailCheckerHost})
	emailValid, err := emailCheckClient.CheckEmail(*u.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !emailValid {
		http.Error(w, "Provided Email is invalid", http.StatusBadRequest)
		return
	}

	err = env.DB.WithTransaction(func() error {
		userID, err := env.DB.AddUser(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil
		}
		writeJsonResponse(w, model.User{Id: userID})
		return nil
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (env *Env) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if env.DB == nil {
		log.Fatal("DB pointer is nil")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	userID, _ := strconv.ParseInt(vars["userId"], 10, 64)
	err := env.DB.WithTransaction(func() error {
		return env.DB.DeleteUser(userID)
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (env *Env) FindUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.ParseInt(vars["userId"], 10, 64)

	if env.DB == nil {
		log.Fatal("DB pointer is nil")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	user, err := env.DB.LoadUserWithId(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJsonResponse(w, user)
}

func (env *Env) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.ParseInt(vars["userId"], 10, 64)

	if env.DB == nil {
		log.Fatal("DB pointer is nil")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Id = userID

	err = env.DB.WithTransaction(func() error {
		return env.DB.UpdateUser(&user)
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
