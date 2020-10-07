package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../dbconn"
	_ "../models"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

func admin_initialize() {
	router.HandleFunc("/admin/login", admin_login).Methods("POST")
	router.HandleFunc("/admin/create", admin_create).Methods("POST")
	router.HandleFunc("/admin/update", admin_update).Methods("POST")
	router.HandleFunc("/admin/delete", admin_delete).Methods("POST")
	router.HandleFunc("/admin/all", admin_all).Methods("GET")
	router.HandleFunc("/admin/byId/{id}", admin_byId).Methods("GET")
}

func admin_login(w http.ResponseWriter, r *http.Request) {

	correct, correctMap := verifyRequest(r, []string{"correo", "password"})
	correctMap["ok"] = false

	if !correct {
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	found, admin, err := dbconn.Admin_getByCorreo(r.Form.Get("correo"))

	if err != nil {
		correctMap["error"] = err.Error()
		correctMap["msg"] = "Error on DB"
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	if !found {
		correctMap["msg"] = "No found"
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	if admin.Contrasena != r.Form.Get("password") {
		correctMap["msg"] = "Bad password"
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	admin.Contrasena = ""

	correctMap["ok"] = true
	correctMap["msg"] = "Success Log in"
	correctMap["admin"] = admin
	json.NewEncoder(w).Encode(correctMap)
}

func admin_create(w http.ResponseWriter, r *http.Request) {

	correct, correctMap := verifyRequest(r, []string{"nombre", "correo", "password"})
	correctMap["ok"] = false

	if !correct {
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	admin, err := dbconn.Admin_create(r.Form.Get("nombre"), r.Form.Get("correo"), r.Form.Get("password"))

	if err != nil {
		correctMap["error"] = err.Error()
		correctMap["msg"] = "Error on DB"
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	admin.Contrasena = ""

	correctMap["ok"] = true
	correctMap["msg"] = "Success creation"
	correctMap["admin"] = admin
	json.NewEncoder(w).Encode(correctMap)
}

func admin_update(w http.ResponseWriter, r *http.Request) {

	correct, correctMap := verifyRequest(r, []string{"id"})
	correctMap["ok"] = false

	if !correct {
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	id, errId := strconv.Atoi(r.Form.Get("id"))

	if errId != nil {
		correctMap["id"] = fmt.Sprintf("'%s' is not a number: %s", r.Form.Get("id"), errId.Error())
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	admin, err := dbconn.Admin_update(id, r.Form.Get("nombre"), r.Form.Get("correo"), r.Form.Get("password"))

	if err != nil {
		correctMap["error"] = err.Error()
		correctMap["msg"] = "Error on DB"
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	admin.Contrasena = ""

	correctMap["ok"] = true
	correctMap["msg"] = "Success update"
	correctMap["admin"] = admin
	json.NewEncoder(w).Encode(correctMap)
}

func admin_delete(w http.ResponseWriter, r *http.Request) {

	correct, correctMap := verifyRequest(r, []string{"id"})
	correctMap["ok"] = false

	if !correct {
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	id, errId := strconv.Atoi(r.Form.Get("id"))

	if errId != nil {
		correctMap["id"] = fmt.Sprintf("'%s' is not a number: %s", r.Form.Get("id"), errId.Error())
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	admin, err := dbconn.Admin_delete(id)

	if err != nil {
		correctMap["error"] = err.Error()
		correctMap["msg"] = "Error on DB"
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	admin.Contrasena = ""

	correctMap["ok"] = true
	correctMap["msg"] = "Success delete"
	correctMap["admin"] = admin
	json.NewEncoder(w).Encode(correctMap)
}

func admin_all(w http.ResponseWriter, r *http.Request) {

	correct, correctMap := verifyRequest(r, []string{})
	correctMap["ok"] = false

	if !correct {
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	admins, err := dbconn.Admin_all()

	if err != nil {
		correctMap["error"] = err.Error()
		correctMap["msg"] = "Error on DB"
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	for _, admin := range admins {
		admin.Contrasena = ""
	}

	correctMap["ok"] = true
	correctMap["msg"] = "Success get all admins"
	correctMap["admins"] = admins
	json.NewEncoder(w).Encode(correctMap)
}

func admin_byId(w http.ResponseWriter, r *http.Request) {

	correct, correctMap := verifyRequest(r, []string{})
	correctMap["ok"] = false

	if !correct {
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	idParam := mux.Vars(r)["id"]

	id, errId := strconv.Atoi(idParam)

	if errId != nil {
		correctMap["id"] = fmt.Sprintf("'%s' is not a number: %s", idParam, errId.Error())
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	admin, err := dbconn.Admin_byId(id)

	if err != nil {
		correctMap["error"] = err.Error()
		correctMap["msg"] = "Error on DB"
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	admin.Contrasena = ""

	correctMap["ok"] = true
	correctMap["msg"] = "Success get admin"
	correctMap["admin"] = admin
	json.NewEncoder(w).Encode(correctMap)
}
