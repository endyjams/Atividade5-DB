package router

import (
	"db-api-example/controller"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Database API example
// @version 1.0
// @description um exemplo de API para a disciplina de banco de dados
// @host localhost:8000
func NewRouter(frutaController *controller.FrutaController, fornecedorController *controller.FornecedorController, estoqueController *controller.EstoqueController) *gin.Engine {
	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/fruta/:nome", frutaController.GetFruta)
	r.POST("/fruta", frutaController.CreateFruta)

	r.GET("/fornecedor/:nome", fornecedorController.GetFornecedor)
	r.POST("/fornecedor", fornecedorController.CreateFornecedor)

	r.GET("/estoque/:nomeFruta", estoqueController.GetEstoque)
	r.POST("/estoque", estoqueController.CreateEstoque)
	return r
}
