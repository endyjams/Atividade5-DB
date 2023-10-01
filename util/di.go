package util

import (
	"Atividade5-DB/controller"
	"Atividade5-DB/database"
	"Atividade5-DB/repository"
	"Atividade5-DB/router"
	"Atividade5-DB/service"
	"fmt"
	"log"
)

func InjectDependencies(db *database.Database, port string) {
	estoqueRepository := &repository.EstoqueDatabase{Db: db}
	estoqueService := &service.EstoqueService{EstoqueRepository: estoqueRepository}
	estoqueController := &controller.EstoqueController{EstoqueService: estoqueService}

	fornecedorRepository := &repository.FornecedorDatabase{Db: db}
	fornecedorService := &service.FornecedorService{FornecedorRepository: fornecedorRepository}
	fornecedorController := &controller.FornecedorController{fornecedorService: fornecedorService}

	frutaRepository := &repository.FrutaDatabase{Db: db}
	frutaService := &service.FrutaService{FrutaRepository: frutaRepository}
	frutaController := &controller.FrutaController{FrutaService: frutaService}

	r := router.NewRouter(frutaController, fornecedorController, estoqueController)

	address := fmt.Sprintf(":%s", port)

	log.Fatal(r.Run(address))
}
