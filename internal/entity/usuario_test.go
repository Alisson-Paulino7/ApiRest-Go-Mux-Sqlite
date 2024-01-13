package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorIfIDIsBlank(t *testing.T) {
	usuario := Usuario{}
	// if usuario.Validate() == nil {
	// 	t.Error("ID is required")
	// }
	assert.Error(t, usuario.ValidateUser(), "ID is required")
}

func Test_If_It_Gets_An_Error_If_Nome_Is_Blank(t *testing.T) {
	usuario := Usuario{ID: "123"}
	assert.Error(t, usuario.ValidateUser(), "Nome is required")
}

func Test_If_It_Gets_An_Error_If_Senha_Is_Blank(t *testing.T) {
	usuario := Usuario{ID: "123", Nome: "Alisson", Email: "alisson@gmail.com"}
	assert.Error(t, usuario.ValidateUser(), "Senha is required")
}

func Test_If_It_Gets_An_Error_If_Endereco_Is_Blank(t *testing.T) {
	usuario := Usuario{ID: "123", Nome: "Alisson", Email: "alisson@gmail.com", Senha: "123456"}
	assert.Error(t, usuario.ValidateUser(), "Endereco is required")
}

func Test_If_It_Gets_An_Error_If_Telefone_Is_Blank(t *testing.T) {
	usuario := Usuario{ID: "123", Nome: "Alisson", Email: "alisson@gmail.com", Senha: "123456", Endereco: Local{Cidade: "São Paulo", Estado: "SP"}}
	assert.Error(t, usuario.ValidateUser(), "Endereco is required")
}