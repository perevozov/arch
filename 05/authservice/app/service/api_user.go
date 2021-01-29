package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/perevozov/arch/05/authservice/model"

	"github.com/gorilla/mux"
)

const SessionCookieName = "SessionId"

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

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

func (env *Env) Register(w http.ResponseWriter, r *http.Request) {
	var u model.User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		log.Printf("Json decode error: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var p struct{ Password string }
	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Printf("Json decode error: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = env.DB.WithTransaction(func() error {
		userID, err := env.DB.AddUser(&u)
		if err != nil {
			log.Printf("Error: %s\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil
		}

		env.DB.SetUserPassword(&u, p.Password)
		writeJsonResponse(w, model.User{Id: userID})
		return nil
	})

	if err != nil {
		log.Printf("Error: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (env *Env) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if loginRequest.Login == "" || loginRequest.Password == "" {
		http.Error(w, "Fields login and password are required", http.StatusBadRequest)
		return
	}
	user, err := env.DB.CheckUserPassword(loginRequest.Login, loginRequest.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	session := CreateSession(user.Id)
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    session.SessionID,
		Expires:  session.Expires,
		Path:     "/",
		HttpOnly: true,
	})
}

func (env *Env) Authorize(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie(SessionCookieName)
	if err != nil {
		log.Printf("Error: %s\n", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	sessionId := sessionCookie.Value
	session := GetSession(sessionId)
	if session == nil {
		log.Printf("Did not find session: %s\n", sessionId)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	user, err := env.DB.LoadUserWithId(session.UserID)
	if err != nil {
		if err == model.ErrUserNotFound {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
	w.Header().Add("X-UserId", strconv.FormatInt(user.Id, 10))
	w.Header().Add("X-UserName", *user.Username)
	w.Header().Add("X-UserFirstName", *user.FirstName)
	w.Header().Add("X-UserLastName", *user.LastName)
}

func (env *Env) Logout(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie(SessionCookieName)
	if err != nil {
		return
	}
	sessionId := sessionCookie.Value
	DeleteSession(sessionId)
	http.SetCookie(w, &http.Cookie{
		Name:  SessionCookieName,
		Value: "",
	})
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
