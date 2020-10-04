package dbconn

import (
	"strings"

	"../models"
)

func Admin_getByCorreo(correo string) (bool, models.Administrador, error) {
	var exit models.Administrador

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
