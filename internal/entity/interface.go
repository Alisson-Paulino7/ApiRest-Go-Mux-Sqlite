package entity


//Padrão que vai falar os métodos que tem que ter
//para persistir dados num banco de dados
type UsuarioRepositoryInterface interface {
	//Método que vai salvar um usuario
	Save(users *Usuario) error
	//Método que vai retornar todos os usuarios
	FindAll() ([]Usuario, error)
	//Método que vai retornar um usuario
	FindOne(id string) (*Usuario, error)
	//Método que vai deletar um usuario
	Delete(id string) error
	//Método que vai atualizar um usuario
	Update(user *Usuario) error
	//Método que vai retornar o total de usuarios
	GetTotal() (int, error)
}
