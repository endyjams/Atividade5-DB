package service

import (
	"db-api-example/model"
	"db-api-example/repository"
)

type FrutaService struct {
	FrutaRepository repository.FrutaRepository
}

func (frutaService *FrutaService) GetFruta(nome string) (*model.Fruta, error) {
	return frutaService.FrutaRepository.GetFruta(nome)
}

func (frutaService *FrutaService) CreateFruta(fruta *model.Fruta) error {
	return frutaService.FrutaRepository.CreateFruta(fruta)
}
