CREATE DATABASE IF NOT EXISTS FRAGER_ADMIN;

USE FRAGER_ADMIN;

/* ------------------Administrador----------------------------------*/
DROP TABLE IF EXISTS Administrador;
CREATE TABLE Administrador (
	id integer primary key unique auto_increment,
    nombre varchar(60) not null,
    correo_electronico varchar(60) not null unique,
    contrasena varchar(120) not null
);

/* Default admin*/
INSERT INTO Administrador(nombre, correo_electronico, contrasena) VALUES ("admin", "admin@admin.com", "0000");

/* ------------------Reporte----------------------------------*/

DROP TABLE IF EXISTS Reporte;
CREATE TABLE Reporte (
	id integer primary key unique auto_increment,
    id_pregunta integer,
    id_usuario_reporte integer,
    id_administrador_solucionado integer,
    comentario varchar(300),
    solucionado boolean,
	FOREIGN KEY (id_administrador_solucionado) REFERENCES Administrador(id)
);
