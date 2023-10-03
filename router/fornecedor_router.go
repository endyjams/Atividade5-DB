package router

import (
	"Atividade5-DB/controller"

	"github.com/gin-gonic/gin"
)

func setFornecedorRoute(fornecedorController *controller.FornecedorController, r *gin.Engine) {
	r.GET("/fornecedor/:nome", fornecedorController.GetFornecedor)
	r.POST("/fornecedor", fornecedorController.CreateFornecedor)
	r.PUT("/fornecedor", fornecedorController.UpdateFornecedor)
}
