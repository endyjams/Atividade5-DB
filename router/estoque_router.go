package router

import (
	"Atividade5-DB/controller"

	"github.com/gin-gonic/gin"
)

func setEstoqueRoute(estoqueController *controller.EstoqueController, r *gin.Engine) {
	r.GET("obter/estoque/:nomeFornecedor/:nomeFruta", estoqueController.GetEstoque)
	r.POST("/criar/estoque", estoqueController.CreateEstoque)
	r.PUT("/atualizar/estoque", estoqueController.UpdateEstoque)
	r.DELETE("/deletar/estoque", estoqueController.DeleteEstoque)
	r.POST("/preencher/estoque", estoqueController.FillEstoque)
}
