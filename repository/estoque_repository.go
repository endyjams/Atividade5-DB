package repository

import (
	"Atividade5-DB/database"
	"Atividade5-DB/model"
	"errors"
)

type EstoqueRepository interface {
	GetEstoque(nomeFruta string, nomeFornecedor string) (*model.Estoque, error)
	CreateEstoque(estoque *model.Estoque) error
	UpdateEstoque(estoque *model.Estoque) error
	DeleteEstoque(estoque *model.Estoque) error
	FillEstoque() error
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

func (estoqueRepository *EstoqueDatabase) CreateEstoque(estoque *model.Estoque) error {
	query :=
		` INSERT INTO estoque` +
			` (quantidade, id_fruta, id_fornecedor)` +
			` VALUES ($3, (SELECT id_fruta FROM fruta WHERE fruta.nome = $1), (SELECT id_fornecedor FROM fornecedor WHERE fornecedor.nome = $2))`

	_, err := estoqueRepository.Db.Conn.Exec(query, estoque.NomeFruta, estoque.NomeFornecedor, estoque.Quantidade)

	return err
}

func (estoqueRepository *EstoqueDatabase) UpdateEstoque(estoque *model.Estoque) error {
	query :=
		` UPDATE estoque` +
			` SET quantidade = $3` +
			` WHERE id_fornecedor = (SELECT id_fornecedor FROM fornecedor WHERE fornecedor.nome = $2)` +
			` AND id_fruta = (SELECT id_fruta FROM fruta WHERE fruta.nome = $1)`

	result, err := estoqueRepository.Db.Conn.Exec(query, estoque.NomeFruta, estoque.NomeFornecedor, estoque.Quantidade)

	rows, err := result.RowsAffected()

	if rows == 0 {
		return errors.New("erro: nenhuma linha foi afetada")
	}

	return err
}

func (estoqueRepository *EstoqueDatabase) DeleteEstoque(estoque *model.Estoque) error {
	query :=
		` DELETE FROM estoque` +
			` WHERE quantidade = $3` +
			` AND id_fornecedor = (SELECT id_fornecedor FROM fornecedor WHERE fornecedor.nome = $2)` +
			` AND id_fruta = (SELECT id_fruta FROM fruta WHERE fruta.nome = $1)`

	result, err := estoqueRepository.Db.Conn.Exec(query, estoque.NomeFruta, estoque.NomeFornecedor, estoque.Quantidade)

	rows, err := result.RowsAffected()

	if rows == 0 {
		return errors.New("erro: nenhuma linha foi afetada")
	}

	return err
}

func (estoqueRepository *EstoqueDatabase) FillEstoque() error {
	tx, err := estoqueRepository.Db.Conn.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	transaction :=
		`WITH FrutaEsgotada AS (
		SELECT id_fruta 
		FROM Fruta
		WHERE id_fruta NOT IN (SELECT id_fruta FROM Estoque)
		LIMIT 1
	),
	FornecedorBaixoEstoque AS (
		SELECT id_fornecedor
		FROM Estoque
		GROUP BY id_fornecedor
		ORDER BY SUM(quantidade) ASC
		LIMIT 1
	)
	
	INSERT INTO Estoque 
	(quantidade, id_fruta, id_fornecedor)
	VALUES (
			10, 
			(SELECT id_fruta FROM FrutaEsgotada), 
			(SELECT id_fornecedor FROM FornecedorBaixoEstoque)
		   )`

	_, err = tx.Exec(transaction)

	if err != nil {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return err
}
