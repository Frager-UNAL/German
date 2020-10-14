package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../dbconn"
	"../models"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

func router_initialize() {
	router.HandleFunc("/report/reported", report_reported).Methods("GET")
	router.HandleFunc("/report", report_all).Methods("GET")
	router.HandleFunc("/report/{id}", report_byId).Methods("GET")
	router.HandleFunc("/report", report_create).Methods("POST")
	router.HandleFunc("/report", report_update).Methods("PUT")
	router.HandleFunc("/report/{id}", report_delete).Methods("DELETE")
}

func router_getErrorValue() models.Reporte {
	return models.Reporte{Id: -1}
}

func report_all(w http.ResponseWriter, r *http.Request) {
	correct, _ := verifyRequest(r, []string{})

	if !correct {
		json.NewEncoder(w).Encode([]models.Reporte{router_getErrorValue()})
		return
	}

	reports, err := dbconn.Report_getAll()

	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode([]models.Reporte{router_getErrorValue()})
		return
	}

	json.NewEncoder(w).Encode(reports)
}

func report_byId(w http.ResponseWriter, r *http.Request) {
	correct, _ := verifyRequest(r, []string{})

	if !correct {
		json.NewEncoder(w).Encode(router_getErrorValue())
		return
	}

	idParam := mux.Vars(r)["id"]

	id, errId := strconv.Atoi(idParam)

	if errId != nil {
		json.NewEncoder(w).Encode(router_getErrorValue())
		return
	}

	report, err := dbconn.Report_getById(id)

	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(router_getErrorValue())
		return
	}

	json.NewEncoder(w).Encode(report)
}

func report_reported(w http.ResponseWriter, r *http.Request) {
	correct, _ := verifyRequest(r, []string{})

	if !correct {
		json.NewEncoder(w).Encode([]models.Reporte{router_getErrorValue()})
		return
	}

	reports, error := dbconn.Report_getUnsolved()

	if error != nil {
		json.NewEncoder(w).Encode([]models.Reporte{router_getErrorValue()})
		return
	}

	json.NewEncoder(w).Encode(reports)
}

func report_create(w http.ResponseWriter, r *http.Request) {
	correct, _ := verifyRequest(r, []string{"id_pregunta", "id_usuario_reporte", "comentario"})

	if !correct {
		json.NewEncoder(w).Encode(router_getErrorValue())
		return
	}

	id_p, errp := strconv.Atoi(r.Form.Get("id_pregunta"))
	id_u, erru := strconv.Atoi(r.Form.Get("id_usuario_reporte"))

	if errp != nil || erru != nil {
		json.NewEncoder(w).Encode(router_getErrorValue())
		return
	}

	report, errDB := dbconn.Report_create(id_p, id_u, r.Form.Get("comentario"))

	if errDB != nil {
		json.NewEncoder(w).Encode(router_getErrorValue())
		return
	}

	json.NewEncoder(w).Encode(report)
}

func report_update(w http.ResponseWriter, r *http.Request) {
	correct, _ := verifyRequest(r, []string{"id"})

	if !correct {
		json.NewEncoder(w).Encode(router_getErrorValue())
		return
	}

	id_r, errRep := strconv.Atoi(r.Form.Get("id"))

	if errRep != nil {
		fmt.Println(errRep)
		json.NewEncoder(w).Encode(router_getErrorValue())
		return
	}

	report, err := dbconn.Report_update(id_r, r.Form.Get("id_pregunta"), r.Form.Get("id_usuario_reporte"), r.Form.Get("id_admin"), r.Form.Get("comentario"), r.Form.Get("solucionado"))

	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(router_getErrorValue())
		return
	}

	json.NewEncoder(w).Encode(report)
}

func report_delete(w http.ResponseWriter, r *http.Request) {
	correct, _ := verifyRequest(r, []string{})

	if !correct {
		json.NewEncoder(w).Encode(router_getErrorValue())
		return
	}

	idParam := mux.Vars(r)["id"]

	id, errId := strconv.Atoi(idParam)

	if errId != nil {
		json.NewEncoder(w).Encode(router_getErrorValue())
		return
	}

	report, err := dbconn.Report_delete(id)

	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(router_getErrorValue())
		return
	}

	json.NewEncoder(w).Encode(report)
}
