package repository

import (
	"Atividade5-DB/database"
	"Atividade5-DB/model"
)

type EstoqueRepository interface {
	GetEstoque(nomeFruta string, nomeFornecedor string) (*model.Estoque, error)
	CreateEstoque(estoque *model.Estoque) error
}

type EstoqueDatabase struct {
	Db *database.Database
}

func (estoqueDatabase *EstoqueDatabase) GetEstoque(nomeFruta string, nomeFornecedor string) (*model.Estoque, error) {
	var estoque model.Estoque

	query := `WITH id_fru AS (SELECT id_fruta FROM fruta WHERE fruta.nome = $1),` +
		` id_forn AS (SELECT id_fornecedor FROM fornecedor WHERE fornecedor.nome = $2)` +
		` SELECT fruta.nome, fornecedor.nome, estoque.quantidade` +
		` FROM fornecedor` +
		` JOIN id_forn ON fornecedor.id_fornecedor = id_forn.id_fornecedor` +
		` JOIN estoque ON fornecedor.id_fornecedor = estoque.id_fornecedor` +
		` JOIN id_fru ON estoque.id_fruta = id_fru.id_fruta` +
		` JOIN fruta ON estoque.id_fruta = fruta.id_fruta;`

	err := estoqueDatabase.Db.Conn.QueryRow(query, nomeFruta, nomeFornecedor).Scan(&estoque.NomeFruta, &estoque.NomeFornecedor, &estoque.Quantidade)

	if err != nil {
		return nil, err
	}

	return &estoque, nil
}

func (estoqueRepository *EstoqueDatabase) CreateUser(estoque *model.Estoque) error {
	_, err := estoqueRepository.Db.Conn.Exec("INSERT INTO users (cpf, name, dateOfBirth) VALUES ($1, $2, $3)", estoque.NomeFruta, estoque.NomeFornecedor, estoque.Quantidade)
	return err
}
