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
	var frutaExists *model.Fruta

	frutaExists, err := frutaService.FrutaRepository.GetFruta(fruta.Nome)

	if frutaExists != nil && err == nil {
		return frutaService.FrutaRepository.UpdateFruta(fruta)
	}

	return frutaService.FrutaRepository.CreateFruta(fruta)
}
