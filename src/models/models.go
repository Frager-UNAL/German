package models

type Administrador struct {
	Id                         int
	Nombre, Correo, Contrasena string
}

type Reporte struct {
	Id, Id_pregunta, Id_usuario_reporte, Id_administrador int
	Comentario                                            string
	Solucionado                                           bool
}
