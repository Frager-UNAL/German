package dbconn

import (
	"fmt"

	"../models"
)

func Report_getAll() ([]models.Reporte, error) {
	exit := make([]models.Reporte, 0)

	results, err := Db.Query("SELECT * FROM Reporte")

	if err != nil {
		return exit, err
	}

	var tmp models.Reporte
	for results.Next() {
		results.Scan(&tmp.Id, &tmp.Id_pregunta, &tmp.Id_usuario_reporte, &tmp.Id_administrador, &tmp.Comentario, &tmp.Solucionado)
		exit = append(exit, tmp)
	}

	err = results.Err()

	return exit, err
}

func Report_getUnsolved() ([]models.Reporte, error) {
	exit := make([]models.Reporte, 0)

	results, err := Db.Query("SELECT * FROM Reporte WHERE solucionado=FALSE")

	if err != nil {
		return exit, err
	}

	var tmp models.Reporte
	for results.Next() {
		results.Scan(&tmp.Id, &tmp.Id_pregunta, &tmp.Id_usuario_reporte, &tmp.Id_administrador, &tmp.Comentario, &tmp.Solucionado)
		exit = append(exit, tmp)
	}

	return exit, nil
}

func Report_create(id_pregunta int, id_usuario_reporte string, comentario string) (models.Reporte, error) {

	report := models.Reporte{
		Id:                 0,
		Id_pregunta:        id_pregunta,
		Id_usuario_reporte: id_usuario_reporte,
		Comentario:         comentario,
		Solucionado:        false,
	}

	query := fmt.Sprintf(
		"INSERT INTO Reporte(id_pregunta, id_usuario_reporte, comentario, solucionado) VALUES (%d, '%s', '%s', %t)",
		report.Id_pregunta,
		report.Id_usuario_reporte,
		report.Comentario,
		report.Solucionado,
	)

	_, err := Db.Query(query)

	if err != nil {
		return report, err
	}

	return report, nil
}

func Report_getById(id int) (models.Reporte, error) {
	report := models.Reporte{Id: -1}

	query := fmt.Sprintf("SELECT * FROM Reporte WHERE id=%d", id)

	results, err := Db.Query(query)

	if err != nil {
		return report, err
	}

	if results.Next() {
		results.Scan(&report.Id, &report.Id_pregunta, &report.Id_usuario_reporte, &report.Id_administrador, &report.Comentario, &report.Solucionado)
	}

	return report, nil
}

func Report_update(id_reporte int, id_pregunta string, id_usuario_reporte string, id_admin string, comentario string, solucionado string) (models.Reporte, error) {

	query := "UPDATE Reporte SET "

	useComa := false

	if id_pregunta != "" {
		query += fmt.Sprintf("id_pregunta=%s", id_pregunta)
		useComa = true
	}

	if id_usuario_reporte != "" {
		if useComa {
			query += ", "
		}
		useComa = true
		query += fmt.Sprintf("id_usuario_reporte='%s'", id_usuario_reporte)
	}

	if id_admin != "" {
		if useComa {
			query += ", "
		}
		useComa = true
		query += fmt.Sprintf("id_administrador_solucionado=%s", id_admin)
	}

	if comentario != "" {
		if useComa {
			query += ", "
		}
		useComa = true
		query += fmt.Sprintf("comentario='%s'", comentario)
	}

	if solucionado != "" {
		if useComa {
			query += ", "
		}
		useComa = true
		query += fmt.Sprintf("solucionado=%t", solucionado == "true")
	}

	query += fmt.Sprintf(" WHERE id=%d", id_reporte)

	_, err := Db.Query(query)

	report := models.Reporte{Id: -1}

	if err != nil {
		return report, err
	}

	query = fmt.Sprintf("SELECT * FROM Reporte WHERE id=%d", id_reporte)

	results, err := Db.Query(query)

	if err != nil {
		return report, err
	}

	if results.Next() {
		err := results.Scan(&report.Id, &report.Id_pregunta, &report.Id_usuario_reporte, &report.Id_administrador, &report.Comentario, &report.Solucionado)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = results.Err()
	if err != nil {
		fmt.Println(err)
	}

	return report, nil
}

func Report_delete(id int) (models.Reporte, error) {
	report := models.Reporte{Id: -1}

	query := fmt.Sprintf("SELECT * FROM Reporte WHERE id=%d", id)

	results, err := Db.Query(query)

	if err != nil {
		return report, err
	}

	if results.Next() {
		results.Scan(&report.Id, &report.Id_pregunta, &report.Id_usuario_reporte, &report.Id_administrador, &report.Comentario, &report.Solucionado)
	}

	query = fmt.Sprintf("DELETE FROM Reporte WHERE id=%d", id)

	_, err = Db.Query(query)

	if err != nil {
		return report, err
	}

	return report, nil
}
