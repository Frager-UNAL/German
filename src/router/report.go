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

func router_initialize() {
	router.HandleFunc("/report/reported", report_reported).Methods("GET")
	router.HandleFunc("/report", report_all).Methods("GET")
	router.HandleFunc("/report/{id}", report_byId).Methods("GET")
	router.HandleFunc("/report", report_create).Methods("POST")
	router.HandleFunc("/report", report_update).Methods("PUT")
	router.HandleFunc("/report/{id}", report_delete).Methods("DELETE")
}

func report_all(w http.ResponseWriter, r *http.Request) {
	correct, correctMap := verifyRequest(r, []string{})
	correctMap["ok"] = false

	if !correct {
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	reports, err := dbconn.Report_getAll()

	if err != nil {
		correctMap["error"] = err.Error()
		correctMap["msg"] = "Error on DB"
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	// correctMap["ok"] = true
	// correctMap["msg"] = "Success Get reports"
	// correctMap["reports"] = reports
	json.NewEncoder(w).Encode(reports)
}

func report_byId(w http.ResponseWriter, r *http.Request) {
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

	report, err := dbconn.Report_getById(id)

	if err != nil {
		correctMap["error"] = err.Error()
		correctMap["msg"] = "Error on DB"
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	// correctMap["ok"] = true
	// correctMap["msg"] = "Success Get report"
	// correctMap["report"] = report
	json.NewEncoder(w).Encode(report)
}

func report_reported(w http.ResponseWriter, r *http.Request) {
	correct, correctMap := verifyRequest(r, []string{})
	correctMap["ok"] = false

	if !correct {
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	reports, error := dbconn.Report_getUnsolved()

	if error != nil {
		correctMap["error"] = error.Error()
		correctMap["msg"] = "Error on DB"
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	// correctMap["ok"] = true
	// correctMap["msg"] = "Success Get reports"
	// correctMap["reports"] = reports
	json.NewEncoder(w).Encode(reports)
}

func report_create(w http.ResponseWriter, r *http.Request) {
	correct, correctMap := verifyRequest(r, []string{"id_pregunta", "id_usuario_reporte", "comentario"})
	correctMap["ok"] = false

	if !correct {
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	id_p, errp := strconv.Atoi(r.Form.Get("id_pregunta"))
	id_u, erru := strconv.Atoi(r.Form.Get("id_usuario_reporte"))

	if errp != nil {
		correctMap["id_pregunta"] = fmt.Sprintf("'%s' is not a number: %s", r.Form.Get("id_pregunta"), errp.Error())
	}

	if erru != nil {
		correctMap["id_usuario_reporte"] = fmt.Sprintf("'%s' is not a number: %s", r.Form.Get("id_usuario_reporte"), errp.Error())
	}

	if errp != nil || erru != nil {
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	report, errDB := dbconn.Report_create(id_p, id_u, r.Form.Get("comentario"))

	if errDB != nil {
		correctMap["error"] = errDB.Error()
		correctMap["msg"] = "Error on DB"
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	// correctMap["ok"] = true
	// correctMap["msg"] = "Success Report Added"
	// correctMap["report"] = report
	json.NewEncoder(w).Encode(report)
}

func report_update(w http.ResponseWriter, r *http.Request) {
	correct, correctMap := verifyRequest(r, []string{"id"})
	correctMap["ok"] = false

	if !correct {
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	id_r, errRep := strconv.Atoi(r.Form.Get("id"))
	console.log(typeof r.Form.Get("id"))

	if errRep != nil {
		correctMap["id"] = fmt.Sprintf("'%s' is not a number: %s", r.Form.Get("id"), errRep.Error())
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	report, err := dbconn.Report_update(id_r, r.Form.Get("id_pregunta"), r.Form.Get("id_usuario_reporte"), r.Form.Get("id_admin"), r.Form.Get("comentario"), r.Form.Get("solucionado"))

	if err != nil {
		correctMap["error"] = err.Error()
		correctMap["msg"] = "Error on DB"
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	// correctMap["ok"] = true
	// correctMap["msg"] = "Success Report Update"
	// correctMap["report"] = report
	json.NewEncoder(w).Encode(report)
}

func report_delete(w http.ResponseWriter, r *http.Request) {
	correct, correctMap := verifyRequest(r, []string{})
	correctMap["ok"] = false

	if !correct {
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	idParam := mux.Vars(r)["id"]

	id, errId := strconv.Atoi(idParam)

	if errId != nil {
		correctMap["id"] = fmt.Sprintf("'%s' is not a number: %s", r.Form.Get("id"), errId.Error())
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	report, err := dbconn.Report_delete(id)

	if err != nil {
		correctMap["error"] = err.Error()
		correctMap["msg"] = "Error on DB"
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	// correctMap["ok"] = true
	// correctMap["msg"] = "Success Report Update"
	// correctMap["report"] = report
	json.NewEncoder(w).Encode(report)
}
