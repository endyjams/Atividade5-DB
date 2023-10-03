package router

import (
	"Atividade5-DB/controller"

	"github.com/gin-gonic/gin"
)

func setEstoqueRoute(estoqueController *controller.EstoqueController, r *gin.Engine) {
	r.GET("/estoque/:nomeFruta/:nomeFornecedor", estoqueController.GetEstoque)
	r.POST("/estoque", estoqueController.CreateEstoque)
	r.PUT("/estoque", estoqueController.UpdateEstoque)
}
