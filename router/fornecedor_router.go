package router

import (
	"Atividade5-DB/controller"

	"github.com/gin-gonic/gin"
)

func setFornecedorRoute(fornecedorController *controller.FornecedorController, r *gin.Engine) {
	r.GET("/obter/fornecedor/:nome", fornecedorController.GetFornecedor)
	r.POST("/criar/fornecedor", fornecedorController.CreateFornecedor)
	r.PUT("/atualizar/fornecedor", fornecedorController.UpdateFornecedor)
}
