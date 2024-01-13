package api

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"

	"github.com/alisson-paulino7/ApiRest-Go-Mux-Sqlite/internal/entity"
	"github.com/alisson-paulino7/ApiRest-Go-Mux-Sqlite/internal/infra/database"
	"github.com/gorilla/mux"

	"time"

	"github.com/google/uuid"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// Só uma bobagem pra teste
	w.Write([]byte("Olá, mundo!"))
}

func CadastrarUser(w http.ResponseWriter, r *http.Request) {

	// Instancia da estrutura
	var user entity.Usuario
	// Decodificando e preenchendo a estrutura
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Definindo um Id único aleatório
	user.ID = uuid.New().String()
	// Iniciando a a função que lança os dados na estrutura, valida se estão todos corretos e retorna ela ou um erro
	createdUser, err := entity.AddUser(user.ID, user.Nome, user.Email, user.Senha, user.Endereco.Cidade, user.Endereco.Estado, user.Telefone)
	if err != nil {
		http.Error(w, "Erro ao adicionar usuário: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Passando a estrutura de retorno validada para o método que vai salvar no BD
	err = database.UserRepository.Save(createdUser)
	if err != nil {
		http.Error(w, "Erro ao salvar usuário: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Retornando a estrutura caso bem sucedida a validação dos dados
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)

}

func BuscarAllUser(w http.ResponseWriter, r *http.Request) {
	allUsers, err := database.UserRepository.FindAll()
	if err != nil {
		http.Error(w, "Erro ao obter usuários do banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Codificando e enviando a lista de usuários como resposta
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allUsers); err != nil {
		http.Error(w, "Erro ao codificar JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func BuscarOneUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := database.UserRepository.FindOne(params["id"])
	if err != nil {
		http.Error(w, "Erro ao obter usuário do banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Erro ao codificar JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeletarUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := database.UserRepository.Delete(params["id"])
	if err != nil {
		http.Error(w, "Erro ao deletar usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Usuário deletado com sucesso!"}`))

}


func AtualizarUser(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	var user entity.Usuario
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Eu ia pegar apenas o parâmetro ID, mas como o usuário pode querer atualizar apenas um campo, eu pego todos os campos e atualizo
	// Eu fiquei na preguiça de criar uma função pra validar os Updates, ai meti logo no AddUser, obrigando a você passar todos os campos
	// user.ID = params["id"]
	updatedUser, err := entity.AddUser(user.ID, user.Nome, user.Email, user.Senha, user.Endereco.Cidade, user.Endereco.Estado, user.Telefone)
	if err != nil {
		http.Error(w, "Erro ao atualizar usuário: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Enviando para o banco após confirmar que os dados estão OK e validados
	err = database.UserRepository.Update(updatedUser)
	if err != nil {
		http.Error(w, "Erro ao atualizar usuário: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Usuário atualizado com sucesso!"}`))

}


//Busca a contagem de usuários executando a função da interface no Bd e retorna o total

func TotalUser(w http.ResponseWriter, r *http.Request) {
	total, err := database.UserRepository.GetTotal()
	if err != nil {
		http.Error(w, "Erro ao obter total de usuários: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Converte o total para string
	totalString := fmt.Sprintf("%d", total)
	// Me retorna informando até a data da busca
	response := fmt.Sprintf("total de registros até %v: %s", time.Now().Format("02-01-2006 15:04:05"), totalString)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Tive que fazer assim porquê tava frescanso pra fazer diferente
	w.Write([]byte(response))

}
