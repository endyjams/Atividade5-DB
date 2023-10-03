package service

import (
	"Atividade5-DB/model"
	repository "Atividade5-DB/repository"
)

type EstoqueService struct {
	EstoqueRepository repository.EstoqueRepository
}

func (estoqueService *EstoqueService) GetEstoque(nomeFruta string, nomeFornecedor string) (*model.Estoque, error) {
	return estoqueService.EstoqueRepository.GetEstoque(nomeFruta, nomeFornecedor)
}

func (estoqueService *EstoqueService) CreateEstoque(estoque *model.Estoque) error {
	return estoqueService.EstoqueRepository.CreateEstoque(estoque)
}

func (estoqueService *EstoqueService) UpdateEstoque(estoque *model.Estoque) error {
	return estoqueService.EstoqueRepository.UpdateEstoque(estoque)
}
