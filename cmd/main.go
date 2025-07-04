package main

import (
	"log"

	"user/micro/config"
	"user/micro/route"

	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró .env, usando variables del sistema")
	}

	// Inicializar MongoDB
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error conectando a MongoDB:", err)
	}
	defer db.Client().Disconnect(nil)

	// Configurar rutas
	router := route.SetupRouter(db)

	log.Println("🚀 Servidor escuchando en :8080")
	router.Run(":8080")
}
