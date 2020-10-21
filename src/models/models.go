package models

type Administrador struct {
	Id                         int
	Nombre, Correo, Contrasena string
}

type Reporte struct {
	Id, Id_pregunta, Id_administrador int
	Comentario, Id_usuario_reporte    string
	Solucionado                       bool
}
