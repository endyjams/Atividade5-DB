package service

import (
	"db-api-example/model"
	"db-api-example/repository"
)

type FornecedorService struct {
	FornecedorRepository repository.FornecedorRepository
}

func (fornecedorService *FornecedorService) GetFornecedor(nome string) (*model.Fornecedor, error) {
	return fornecedorService.FornecedorRepository.GetFornecedor(nome)
}

func (fornecedorService *FornecedorService) CreateFornecedor(fornecedor *model.Fornecedor) error {
	return fornecedorService.FornecedorRepository.CreateFornecedor(fornecedor)
}
