package repository

import (
	"Atividade5-DB/database"
	"Atividade5-DB/model"
)

type FrutaRepository interface {
	GetFruta(nome string) (*model.Fruta, error)
	CreateFruta(estoque *model.Fruta) error
}

type FrutaDatabase struct {
	Db *database.Database
}

//TODO Functions
