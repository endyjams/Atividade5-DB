package service

import (
	"Atividade5-DB/model"
	"Atividade5-DB/repository"
)

type FornecedorService struct {
	FornecedorRepository repository.FornecedorRepository
}

func (fornecedorService *FornecedorService) GetFornecedor(nome string) (*model.Fornecedor, error) {
	return fornecedorService.FornecedorRepository.GetFornecedor(nome)
}

func (fornecedorService *FornecedorService) CreateFornecedor(fornecedor *model.Fornecedor) error {
	var fornecedorExists *model.Fornecedor

	fornecedorExists, err := fornecedorService.FornecedorRepository.GetFornecedor(fornecedor.Nome)

	if fornecedorExists != nil && err == nil {
		fornecedorService.FornecedorRepository.UpdateFornecedor(fornecedor)
	}

	return fornecedorService.FornecedorRepository.CreateFornecedor(fornecedor)
}
