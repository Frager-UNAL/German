package dbconn

import (
	"fmt"

	"../models"
)

func Report_getUnsolved() ([]models.Reporte, error) {
	exit := make([]models.Reporte, 0)

	results, err := Db.Query("SELECT * FROM Reporte WHERE solucionado = FALSE")

	if err != nil {
		return exit, err
	}

	var tmp models.Reporte
	for results.Next() {
		tmp.Comentario = "Comentario por defecto"
		results.Scan(&tmp.Id, &tmp.Id_pregunta, &tmp.Id_usuario_reporte, &tmp.Id_administrador, &tmp.Comentario, &tmp.Solucionado)
		exit = append(exit, tmp)
	}

	return exit, nil
}

func Report_addReport(id_pregunta int, id_usuario_reporte int, comentario string) (models.Reporte, error) {

	report := models.Reporte{
		Id_pregunta:        id_pregunta,
		Id_usuario_reporte: id_usuario_reporte,
		Comentario:         comentario,
		Solucionado:        false,
	}

	query := fmt.Sprintf(
		"INSERT INTO Reporte(id_pregunta, id_usuario_reporte, comentario, solucionado) VALUES (%d, %d, '%s', %t)",
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

func Report_solveReport(id_reporte int, solucionado bool) error {
	query := fmt.Sprintf("UPDATE Reporte SET solucionado=%t WHERE id=%d", solucionado, id_reporte)
	_, err := Db.Query(query)
	return err
}
