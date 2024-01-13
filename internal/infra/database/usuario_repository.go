package database

import (
	"github.com/alisson-paulino7/ApiRest-Go-Mux-Sqlite/internal/entity"

	"database/sql"
	"fmt"
	// "log"
	// "errors"
)

type UsuarioRepository struct {
	Db *sql.DB
}

// Ao executar ela retorna uma instância que é passada diretamente para a variável global
func NewUsuarioRepository(db *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{Db: db}
}

// Criando uma variável global com a instancia da estrutura que vai ser usada em todos os métodos
// Assim consigo importar e usar os métodos em outros arquivos
var UserRepository *UsuarioRepository

// Funcão que recebe a conexão com o banco lá do Main e atribui a variável global
// Ela inicia o repositório com o banco de dados
func InitRepository(db *sql.DB) {
	UserRepository = NewUsuarioRepository(db)
}

func (r *UsuarioRepository) Save(u *entity.Usuario) error {

	stmt, err := r.Db.Prepare("INSERT INTO usuarios (id, nome, email, senha, cidade, estado, telefone) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		// log.Printf("Erro ao preparar a declaração SQL: %v", err)
		return fmt.Errorf("falha ao preparar a declaração SQL: %v", err)
	}
	defer stmt.Close()
	add, err := stmt.Exec(u.ID, u.Nome, u.Email, u.Senha, u.Endereco.Cidade, u.Endereco.Estado, u.Telefone)
	if err != nil {
		// log.Printf("Erro ao executar a declaração SQL: %v", err)
		return fmt.Errorf("falha ao executar a declaração SQL: %v", err)
	}
	rowsAffected, err := add.RowsAffected()
	// Verificando se as linhas afetadas tiveram erro na execução do delete
	if err != nil {
		return fmt.Errorf("falha ao salvar os dados devido o erro: %v", err)
	}
	// Verificando se não tiveram linhas afetadas
	if rowsAffected == 0 {
		return fmt.Errorf("nenhum usuário cadastrado: %v", err)
	}
	return nil
}

func (r *UsuarioRepository) FindAll() ([]*entity.Usuario, error) {
	rows, err := r.Db.Query("SELECT * FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Definindo uma fatia que aponta pra estrutura de Usuario e armazenará todos os dados recebidos de todos os usuários
	var usuarios []*entity.Usuario
	for rows.Next() {
		// Definindo uma variável de uso local pra armazenar os usuários individuais
		var u entity.Usuario
		err := rows.Scan(&u.ID, &u.Nome, &u.Email, &u.Senha, &u.Endereco.Cidade, &u.Endereco.Estado, &u.Telefone)
		if err != nil {
			return nil, err
		}
		// Adicionando cada usuário individual na fatia de usuários
		usuarios = append(usuarios, &u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (r *UsuarioRepository) FindOne(id string) (*entity.Usuario, error) {

	stmt, err := r.Db.Prepare("SELECT * FROM usuarios WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("falha ao preparar a declaração SQL: %v", err)
	}
	defer stmt.Close()

	var user entity.Usuario
	// Pegando os valores no Scan e já inserindo na struct
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Nome, &user.Email, &user.Senha, &user.Endereco.Cidade, &user.Endereco.Estado, &user.Telefone)
	if err != nil {
		if err == sql.ErrNoRows {
			// Retorno se não achar ninguém
			return nil, fmt.Errorf("nenhum usuário encontrado com o id %s", id)
		}
		return nil, fmt.Errorf("falha ao executar a declaração SQL: %v", err)
	}

	return &user, nil
}

func (r *UsuarioRepository) Delete(id string) error {

	stmt, err := r.Db.Prepare("DELETE FROM usuarios WHERE ID = ?")
	if err != nil {
		return fmt.Errorf("falha ao preparar a declaração SQL: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("falha ao executar a declaração SQL: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	// Verificando se as linhas afetadas tiveram erro na execução do delete
	if err != nil {
		return fmt.Errorf("falha ao obter o número de linhas afetadas pelo comando: %v", err)
	}
	// Verificando se não tiveram linhas afetadas
	if rowsAffected == 0 {
		return fmt.Errorf("nenhum usuário encontrado com o id %s", id)
	}
	return nil
}

// Mesma coisa do Save
func (r *UsuarioRepository) Update(user *entity.Usuario) error {
	
	stmt, err := r.Db.Prepare("UPDATE usuarios SET nome = ?, email = ?, senha = ?, cidade = ?, estado = ?, telefone = ? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("falha ao preparar a declaração SQL: %v", err)
	}
	defer stmt.Close()
	
	result, err := stmt.Exec(user.Nome, user.Email, user.Senha, user.Endereco.Cidade, user.Endereco.Estado, user.Telefone, user.ID)
	if err != nil {
		return fmt.Errorf("falha ao executar a declaração SQL: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("falha ao obter o número de linhas afetadas pelo comando: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("nenhum usuário encontrado com o id %s", user.ID)
	}

	return nil
}

// Verificando quantos usuários já cadastrei no banco
func (r *UsuarioRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("SELECT count(*) FROM usuarios").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
