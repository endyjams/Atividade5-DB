package service

import (
	"db-api-example/model"
	"db-api-example/repository"
)

type EstoqueService struct {
	EstoqueRepository repository.FrutaRepository
}

func (estoqueService *EstoqueService) GetEstoque(nomeFruta string) (*model.Estoque, error) {
	return estoqueService.EstoqueRepository.GetEstoque(nomeFruta)
}

func (estoqueService *EstoqueService) CreateEstoque(estoque *model.Estoque) error {
	return estoqueService.EstoqueRepository.CreateEstoque(estoque)
}
