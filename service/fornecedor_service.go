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
	return fornecedorService.FornecedorRepository.CreateFornecedor(fornecedor)
}

func (fornecedorService *FornecedorService) UpdateFornecedor(fornecedor *model.Fornecedor) error {
	return fornecedorService.FornecedorRepository.UpdateFornecedor(fornecedor)
}
