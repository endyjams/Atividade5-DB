package router

import (
	"Atividade5-DB/controller"

	"github.com/gin-gonic/gin"
)

func setFrutaRoute(frutaController *controller.FrutaController, r *gin.Engine) {
	r.GET("/obter/fruta/:nome", frutaController.GetFruta)
	r.POST("/criar/fruta", frutaController.CreateFruta)
	r.PUT("/atualizar/fruta", frutaController.UpdateFruta)
}
