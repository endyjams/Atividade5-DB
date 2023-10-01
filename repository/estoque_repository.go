package repository

import (
	"Atividade5-DB/database"
	"Atividade5-DB/model"
)

type EstoqueRepository interface {
	GetEstoque(nomeFruta string) (*model.Estoque, error)
	CreateEstoque(estoque *model.Estoque) error
}

type FrutaDatabase struct {
	Db *database.Database
}

func (frutaDatabase *FrutaDatabase) GetEstoque(nomeFruta string) (*model.Estoque, error) {
	var estoque model.Estoque

	err := frutaDatabase.Db.Conn.QueryRow("SELECT * FROM users WHERE cpf=$1", nomeFruta).Scan(&estoque.nomeFruta, &estoque.nomeFornecedor, &estoque.quantidade)

	if err != nil {
		return nil, err
	}

	return &estoque, nil
}

// func (userRepository *UserDatabase) CreateUser(user *model.User) error {
// 	_, err := userRepository.Db.Conn.Exec("INSERT INTO users (cpf, name, dateOfBirth) VALUES ($1, $2, $3)", user.Cpf, user.Name, user.DateOfBirth)
// 	return err
// }
