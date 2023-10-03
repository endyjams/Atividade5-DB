package router

import (
	"Atividade5-DB/controller"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title DB - Atividade 5
// @version 1.0
// @description última parte do trabalho para disciplina banco de dados
// @host localhost:8000

func NewRouter(frutaController *controller.FrutaController, fornecedorController *controller.FornecedorController, estoqueController *controller.EstoqueController) *gin.Engine {
	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	setFrutaRoute(frutaController, r)
	setFornecedorRoute(fornecedorController, r)
	setEstoqueRoute(estoqueController, r)

	return r
}
