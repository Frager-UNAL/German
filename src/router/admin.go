package router

import (
	"encoding/json"
	"net/http"

	"../dbconn"
	_ "../models"
	_ "github.com/gorilla/mux"
)

func admin_initialize() {
	router.HandleFunc("/admin/login", admin_login).Methods("POST")
}

func admin_login(w http.ResponseWriter, r *http.Request) {

	correct, correctMap := verifyRequest(r, []string{"correo", "password"})
	correctMap["ok"] = false

	if !correct {
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	found, admin, error := dbconn.Admin_getByCorreo(r.Form.Get("correo"))

	if error != nil {
		correctMap["error"] = error.Error()
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

	correctMap["ok"] = true
	correctMap["msg"] = "Success Log in"
	json.NewEncoder(w).Encode(correctMap)
}
