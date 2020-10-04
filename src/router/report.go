package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"../dbconn"
	_ "../models"
	_ "github.com/gorilla/mux"
)

func router_initialize() {
	router.HandleFunc("/report/reported", router_reported).Methods("GET")
	router.HandleFunc("/report/add", router_addReport).Methods("POST")
	router.HandleFunc("/report/solve", router_solveReport).Methods("POST")
}

func router_reported(w http.ResponseWriter, r *http.Request) {
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

	correctMap["ok"] = true
	correctMap["msg"] = "Success Get reports"
	correctMap["reports"] = reports
	json.NewEncoder(w).Encode(correctMap)
}

func router_addReport(w http.ResponseWriter, r *http.Request) {
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

	report, errDB := dbconn.Report_addReport(id_p, id_u, r.Form.Get("comentario"))

	if errDB != nil {
		correctMap["error"] = errDB.Error()
		correctMap["msg"] = "Error on DB"
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	correctMap["ok"] = true
	correctMap["msg"] = "Success Report Added"
	correctMap["report"] = report
	json.NewEncoder(w).Encode(correctMap)
}

func router_solveReport(w http.ResponseWriter, r *http.Request) {
	correct, correctMap := verifyRequest(r, []string{"id_reporte", "solucionado"})
	correctMap["ok"] = false

	if !correct {
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	id_r, errRep := strconv.Atoi(r.Form.Get("id_reporte"))

	if errRep != nil {
		correctMap["id_reporte"] = fmt.Sprintf("'%s' is not a number: %s", r.Form.Get("id_reporte"), errRep.Error())
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	solucionado := strings.Compare(strings.ToLower(r.Form.Get("solucionado")), "true")

	errDB := dbconn.Report_solveReport(id_r, solucionado == 0)

	if errDB != nil {
		correctMap["error"] = errDB.Error()
		correctMap["msg"] = "Error on DB"
		json.NewEncoder(w).Encode(correctMap)
		return
	}

	correctMap["ok"] = true
	correctMap["msg"] = "Success Report Update"
	json.NewEncoder(w).Encode(correctMap)
}
