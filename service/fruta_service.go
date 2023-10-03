package service

import (
	"Atividade5-DB/model"
	"Atividade5-DB/repository"
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

func (frutaService *FrutaService) UpdateFruta(fruta *model.Fruta) error {
	return frutaService.FrutaRepository.UpdateFruta(fruta)
}

func (frutaService *FrutaService) DeleteFruta(fruta *model.Fruta) error {
	return frutaService.FrutaRepository.DeleteFruta(fruta)
}
