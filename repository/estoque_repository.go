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

func (frutaDatabase *FrutaDatabase) GetEstoque(nomeFruta string, nomeFornecedor string) (*model.Estoque, error) {
	var estoque model.Estoque

	err := frutaDatabase.Db.Conn.QueryRow("
	WITH
	id_fru AS (
		SELECT id_fruta 
		FROM fruta
		WHERE fruta.nome = $1
	),
	id_forn AS (
		SELECT id_fornecedor 
		FROM fornecedor
		WHERE fornecedor.nome = $2
	)
	
	SELECT fruta.nome, fornecedor.nome, estoque.quantidade
	FROM fornecedor
	JOIN id_forn ON fornecedor.id_fornecedor = id_forn.id_fornecedor
	JOIN estoque ON fornecedor.id_fornecedor = estoque.id_fornecedor
	JOIN id_fru ON estoque.id_fruta = id_fru.id_fruta
	JOIN fruta ON estoque.id_fruta = fruta.id_fruta;", nomeFruta, nomeFornecedor).Scan(&estoque.nomeFruta, &estoque.nomeFornecedor, &estoque.quantidade)

	if err != nil {
		return nil, err
	}

	return &estoque, nil
}

// func (userRepository *UserDatabase) CreateUser(user *model.User) error {
// 	_, err := userRepository.Db.Conn.Exec("INSERT INTO users (cpf, name, dateOfBirth) VALUES ($1, $2, $3)", user.Cpf, user.Name, user.DateOfBirth)
// 	return err
// }
