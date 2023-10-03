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
	var estoqueExists *model.Estoque

	estoqueExists, err := estoqueService.EstoqueRepository.GetEstoque(estoque.NomeFruta, estoque.NomeFornecedor)

	if estoqueExists == nil || err != nil {
		return estoqueService.EstoqueRepository.UpdateEstoque(estoque)
	}

	return estoqueService.EstoqueRepository.CreateEstoque(estoque)
}
