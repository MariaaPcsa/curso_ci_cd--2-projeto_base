// package database

import (
	"fmt"
	"log"
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	// Lê variáveis do ambiente (útil p/ CI, Docker, Prod)
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	// Monta string de conexão
	stringDeConexao := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	// Conecta usando GORM + Postgres
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar com banco de dados: %v", err)
	}

	// Faz migrações automáticas
	err = DB.AutoMigrate(&models.Aluno{})
	if err != nil {
		log.Printf("Erro ao migrar models: %v", err)
	}

	log.Println("Banco de dados conectado com sucesso!")
}
