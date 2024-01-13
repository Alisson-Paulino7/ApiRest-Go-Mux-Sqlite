package entity

import "errors"

type Usuarios struct {
	Users []*Usuario `json:"users"`
}

type Usuario struct {
	ID 		 string `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Senha    string `json:"senha"`
	Endereco Local  `json:"endereco"`
	Telefone string `json:"telefone"`
}

type Local struct {
	Cidade string `json:"cidade"`
	Estado string `json:"estado"`
}

func AddUser(id, nome, email, senha, cidade, estado, Telefone string) (*Usuario, error) {
	user := &Usuario{
		ID: id,
		Nome: nome,
		Email: email,
		Senha: senha,
		Endereco: Local{
			Cidade: cidade,
			Estado: estado,
		},
		Telefone: Telefone,
	}		
		err := user.ValidateUser()
		if err != nil {
			return nil, err
		}
		return user, nil
}

func (u *Usuario) ValidateUser() error {
	
	if u.Nome == "" {
		return errors.New("o campo nome não pode ser vazio")
	}
	if u.Email == "" {
		return errors.New("o campo email não pode ser vazio")
	}
	if u.Senha == "" {
		return errors.New("o campo senha não pode ser vazio")
	}
	if u.Endereco.Cidade == "" {
		return errors.New("o campo cidade não pode ser vazio")
	}
	if u.Endereco.Estado == "" {
		return errors.New("o campo estado não pode ser vazio")
	}
	if u.Telefone == "" {
		return errors.New("o campo telefone não pode ser vazio")
	}
	return nil
}
