package main

import (

	"database/sql"
	"fmt"
	
	"github.com/alisson-paulino7/ApiRest-Go-Mux-Sqlite/internal/infra/database"
	"github.com/alisson-paulino7/ApiRest-Go-Mux-Sqlite/internal/routes"

	_ "github.com/mattn/go-sqlite3"

)

func main() {

	// Criando a tabela de nome db.sqlite3
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		fmt.Printf("Erro ao abrir o banco de dados: %v\n", err)
		return
	}
	defer db.Close()

	// _, err = db.Exec(`
    //     CREATE TABLE usuarios (
    //         id TEXT PRIMARY KEY,
    //         nome TEXT,
    //         email TEXT,
    //         senha TEXT,
    //         cidade TEXT,
    //         estado TEXT,
    //         telefone TEXT
    //     )
    // `)
	// if err != nil {
	// 	fmt.Printf("Erro ao criar a tabela de usuários: %v\n", err)
	// 	return
	// }
	// Passando para o repositório Iniciar recebendo como parâmetro o banco de dados criado na Main
	database.InitRepository(db)

	routes.LoadRoutes()

	// var newusers entity.Usuarios

	// newusers.Users = append(newusers.Users, &entity.Usuario{ID: "1",
	// 	Nome:     "Alisson",
	// 	Email:    "alisso♂n@gmail.com",
	// 	Senha:    "12345",
	// 	Endereco: entity.Local{Cidade: "Juzeiro", Estado: "Ceará"},
	// 	Telefone: "56546444"})

}
