package database

import (
	"log"

	model "github.com/ProyectoIntegradorSoftware/MicroservicioUsuario/dominio"
	"gorm.io/gorm"
)

// EjecutarMigraciones realiza todas las migraciones necesarias en la base de datos.
func EjecutarMigraciones(db *gorm.DB) {

	db.AutoMigrate(&model.UsuarioGORM{})

	log.Println("Migraciones completadas")
}
