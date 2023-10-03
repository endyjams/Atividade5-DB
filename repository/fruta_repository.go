package repository

import (
	"Atividade5-DB/database"
	"Atividade5-DB/model"
)

type FrutaRepository interface {
	GetFruta(nome string) (*model.Fruta, error)
	CreateFruta(estoque *model.Fruta) error
	UpdateFruta(estoque *model.Fruta) error
	DeleteFruta(estoque *model.Fruta) error
}

type FrutaDatabase struct {
	Db *database.Database
}

func (frutaDatabase *FrutaDatabase) GetFruta(nome string) (*model.Fruta, error) {
	var fruta model.Fruta

	query := `SELECT fruta.nome, fruta.preco FROM fruta where fruta.nome = $1`

	err := frutaDatabase.Db.Conn.QueryRow(query, nome).Scan(&fruta.Nome, &fruta.Preco)

	if err != nil {
		return nil, err
	}

	return &fruta, nil
}

func (frutaRepository *FrutaDatabase) CreateFruta(fruta *model.Fruta) error {
	_, err := frutaRepository.Db.Conn.Exec("INSERT INTO fruta (nome, preco) VALUES ($1, $2)", fruta.Nome, fruta.Preco)
	return err
}

func (frutaRepository *FrutaDatabase) UpdateFruta(fruta *model.Fruta) error {
	_, err := frutaRepository.Db.Conn.Exec("UPDATE fruta SET preco = $2 where nome = $1", fruta.Nome, fruta.Preco)
	return err
}

func (frutaRepository *FrutaDatabase) DeleteFruta(fruta *model.Fruta) error {
	_, err := frutaRepository.Db.Conn.Exec("DELETE FROM fruta WHERE preco = $2 AND nome = $1", fruta.Nome, fruta.Preco)
	return err
}
