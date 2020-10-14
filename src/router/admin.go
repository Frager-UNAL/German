package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../dbconn"
	"../models"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

func admin_initialize() {
	router.HandleFunc("/admin/login", admin_login).Methods("POST")
	router.HandleFunc("/admin", admin_create).Methods("POST")
	router.HandleFunc("/admin", admin_update).Methods("PUT")
	router.HandleFunc("/admin/{id}", admin_delete).Methods("DELETE")
	router.HandleFunc("/admin", admin_all).Methods("GET")
	router.HandleFunc("/admin/{id}", admin_byId).Methods("GET")
}

func admin_getErrorValue() models.Administrador {
	return models.Administrador{Id: -1}
}

func admin_login(w http.ResponseWriter, r *http.Request) {

	correct, _ := verifyRequest(r, []string{"correo", "password"})

	if !correct {
		json.NewEncoder(w).Encode(admin_getErrorValue())
		return
	}

	found, admin, err := dbconn.Admin_getByCorreo(r.Form.Get("correo"))

	if err != nil {
		json.NewEncoder(w).Encode(admin_getErrorValue())
		return
	}

	if !found {
		json.NewEncoder(w).Encode(admin_getErrorValue())
		return
	}

	if admin.Contrasena != r.Form.Get("password") {
		json.NewEncoder(w).Encode(admin_getErrorValue())
		return
	}

	admin.Contrasena = ""

	json.NewEncoder(w).Encode(admin)
}

func admin_create(w http.ResponseWriter, r *http.Request) {

	correct, _ := verifyRequest(r, []string{"nombre", "correo", "password"})

	if !correct {
		json.NewEncoder(w).Encode(admin_getErrorValue())
		return
	}

	admin, err := dbconn.Admin_create(r.Form.Get("nombre"), r.Form.Get("correo"), r.Form.Get("password"))

	if err != nil {
		json.NewEncoder(w).Encode(admin_getErrorValue())
		return
	}

	admin.Contrasena = ""
	json.NewEncoder(w).Encode(admin)
}

func admin_update(w http.ResponseWriter, r *http.Request) {

	correct, _ := verifyRequest(r, []string{"id"})

	if !correct {
		json.NewEncoder(w).Encode(admin_getErrorValue())
		return
	}

	id, errId := strconv.Atoi(r.Form.Get("id"))

	if errId != nil {
		json.NewEncoder(w).Encode(admin_getErrorValue())
		return
	}

	admin, err := dbconn.Admin_update(id, r.Form.Get("nombre"), r.Form.Get("correo"), r.Form.Get("password"))

	if err != nil {
		json.NewEncoder(w).Encode(admin_getErrorValue())
		return
	}

	admin.Contrasena = ""
	json.NewEncoder(w).Encode(admin)
}

func admin_delete(w http.ResponseWriter, r *http.Request) {

	correct, _ := verifyRequest(r, []string{})

	if !correct {
		json.NewEncoder(w).Encode(admin_getErrorValue())
		return
	}

	idParam := mux.Vars(r)["id"]
	id, errId := strconv.Atoi(idParam)

	if errId != nil {
		json.NewEncoder(w).Encode(admin_getErrorValue())
		return
	}

	admin, err := dbconn.Admin_delete(id)

	if err != nil {
		json.NewEncoder(w).Encode(admin_getErrorValue())
		return
	}

	admin.Contrasena = ""
	json.NewEncoder(w).Encode(admin)
}

func admin_all(w http.ResponseWriter, r *http.Request) {

	correct, _ := verifyRequest(r, []string{})

	if !correct {
		json.NewEncoder(w).Encode([]models.Administrador{admin_getErrorValue()})
		return
	}

	admins, err := dbconn.Admin_all()

	if err != nil {
		json.NewEncoder(w).Encode([]models.Administrador{admin_getErrorValue()})
		return
	}

	for _, admin := range admins {
		admin.Contrasena = ""
	}

	json.NewEncoder(w).Encode(admins)
}

func admin_byId(w http.ResponseWriter, r *http.Request) {

	correct, _ := verifyRequest(r, []string{})

	if !correct {
		json.NewEncoder(w).Encode(admin_getErrorValue())
		return
	}

	idParam := mux.Vars(r)["id"]

	id, errId := strconv.Atoi(idParam)

	if errId != nil {
		json.NewEncoder(w).Encode(admin_getErrorValue())
		return
	}

	admin, err := dbconn.Admin_byId(id)

	if err != nil {
		json.NewEncoder(w).Encode(admin_getErrorValue())
		return
	}

	admin.Contrasena = ""

	json.NewEncoder(w).Encode(admin)
}
