package repository

import (
	"Atividade5-DB/database"
	model "Atividade5-DB/model"
)

type FornecedorRepository interface {
	GetFornecedor(nome string) (*model.Fornecedor, error)
	CreateFornecedor(estoque *model.Fornecedor) error
}

type FornecedorDatabase struct {
	Db *database.Database
}

func (fornecedorDatabase *FornecedorDatabase) GetFornecedor(nome string) (*model.Fornecedor, error) {
	var fornecedor model.Fornecedor

	query := `SELECT fornecedor.nome, fornecedor.telefone FROM fornecedor where fornecedor.nome = $1`

	err := fornecedorDatabase.Db.Conn.QueryRow(query, nome).Scan(&fornecedor.Nome)

	if err != nil {
		return nil, err
	}

	return &fornecedor, nil
}

func (fornecedorRepository *FornecedorDatabase) CreateFornecedor(fornecedor *model.Fornecedor) error {
	_, err := fornecedorRepository.Db.Conn.Exec("INSERT INTO users (cpf, name, dateOfBirth) VALUES ($1, $2, $3)", fornecedor.Nome, fornecedor.Telefone)
	return err
}
