package router

import (
	"Atividade5-DB/controller"

	"github.com/gin-gonic/gin"
)

func setEstoqueRoute(estoqueController *controller.EstoqueController, r *gin.Engine) {
	r.GET("obter/estoque/:nomeFruta/:nomeFornecedor", estoqueController.GetEstoque)
	r.POST("/criar/estoque", estoqueController.CreateEstoque)
	r.PUT("/atualizar/estoque", estoqueController.UpdateEstoque)
}
