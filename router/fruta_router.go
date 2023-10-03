package router

import (
	"Atividade5-DB/controller"

	"github.com/gin-gonic/gin"
)

func setFrutaRoute(frutaController *controller.FrutaController, r *gin.Engine) {
	r.GET("/fruta/:nome", frutaController.GetFruta)
	r.POST("/fruta", frutaController.CreateFruta)
}
