package dbconn

import (
	"fmt"
	"strings"

	"../models"
)

func Admin_getByCorreo(correo string) (bool, models.Administrador, error) {
	var exit models.Administrador = models.Administrador{Id: -1}

	results, err := Db.Query("SELECT * from Administrador WHERE correo_electronico = '" + strings.ToLower(correo) + "'")

	if err != nil {
		return false, exit, err
	}

	if results.Next() {
		results.Scan(&exit.Id, &exit.Nombre, &exit.Correo, &exit.Contrasena)
		return true, exit, nil
	}

	return false, exit, nil
}

func Admin_create(nombre string, correo string, contrasena string) (models.Administrador, error) {

	query := fmt.Sprintf("INSERT INTO Administrador(nombre, correo_electronico, contrasena) VALUES ('%s', '%s', '%s')", nombre, correo, contrasena)

	results, err := Db.Query(query)

	var administrador models.Administrador = models.Administrador{Id: -1}

	if err != nil {
		return administrador, err
	}

	query = fmt.Sprintf("SELECT * FROM Administrador WHERE correo_electronico='%s'", correo)

	results, err = Db.Query(query)

	if err != nil {
		return administrador, err
	}

	if results.Next() {
		results.Scan(&administrador.Id, &administrador.Nombre, &administrador.Correo, &administrador.Contrasena)
	}

	return administrador, nil
}

func Admin_update(id int, nombre string, correo string, contrasena string) (models.Administrador, error) {

	query := "UPDATE Administrador SET "

	useComa := false

	if nombre != "" {
		query += fmt.Sprintf("nombre='%s'", nombre)
		useComa = true
	}

	if correo != "" {
		if useComa {
			query += ", "
		}

		useComa = true

		query += fmt.Sprintf("correo_electronico='%s'", correo)

	}

	if contrasena != "" {

		if useComa {
			query += ", "
		}

		query += fmt.Sprintf("contrasena='%s'", contrasena)
	}

	query += fmt.Sprintf(" WHERE id=%d", id)

	_, err := Db.Query(query)

	var administrador models.Administrador = models.Administrador{Id: -1}

	if err != nil {
		return administrador, err
	}

	query = fmt.Sprintf("SELECT * FROM Administrador WHERE id=%d", id)

	results, err := Db.Query(query)

	if err != nil {
		return administrador, err
	}

	if results.Next() {
		results.Scan(&administrador.Id, &administrador.Nombre, &administrador.Correo, &administrador.Contrasena)
	}

	return administrador, nil
}

func Admin_delete(id int) (models.Administrador, error) {

	query := fmt.Sprintf("SELECT * FROM Administrador WHERE id=%d", id)

	results, err := Db.Query(query)

	var administrador models.Administrador = models.Administrador{Id: -1}

	if err != nil {
		return administrador, err
	}

	if results.Next() {
		results.Scan(&administrador.Id, &administrador.Nombre, &administrador.Correo, &administrador.Contrasena)
	}

	query = fmt.Sprintf("DELETE FROM Administrador WHERE id=%d", id)

	_, err = Db.Query(query)

	if err != nil {
		return administrador, err
	}

	return administrador, nil
}

func Admin_all() ([]models.Administrador, error) {

	query := fmt.Sprintf("SELECT * FROM Administrador")

	results, err := Db.Query(query)

	exit := make([]models.Administrador, 0)

	if err != nil {
		return exit, err
	}

	var tmp models.Administrador
	for results.Next() {
		results.Scan(&tmp.Id, &tmp.Nombre, &tmp.Correo, &tmp.Contrasena)
		exit = append(exit, tmp)
	}

	return exit, nil
}

func Admin_byId(id int) (models.Administrador, error) {

	query := fmt.Sprintf("SELECT * FROM Administrador WHERE id=%d", id)

	results, err := Db.Query(query)

	var administrador models.Administrador = models.Administrador{Id: -1}

	if err != nil {
		return administrador, err
	}

	if results.Next() {
		results.Scan(&administrador.Id, &administrador.Nombre, &administrador.Correo, &administrador.Contrasena)
	}

	return administrador, nil
}
