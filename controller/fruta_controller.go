package controller

import (
	"Atividade5-DB/model"
	"Atividade5-DB/service"

	"github.com/gin-gonic/gin"
)

type FrutaController struct {
	FrutaService *service.FrutaService
}

// @Summary Busca por uma fruta pelo nome
// @Description Retorna as informaçōes de uma fruta a partir de seu nome
// @ID get-fruta
// @Produce json
// @Param nome path string true "nome"
// @Success 200 {object} model.Fruta "Retorna as informaçōes da fruta"
// @Failure 400 {object} string "O nome da fruta não deve ser vazio, e pode conter no máximo 50 caracteres"
// @Failure 404 {object} string "Fruta não encontrada"
// @Router /obter/fruta/{nome} [get]
func (frutaController *FrutaController) GetFruta(ctx *gin.Context) {
	nome := ctx.Param("nome")

	if nome == "" || len(nome) > 50 {
		ctx.JSON(400, gin.H{"erro": "O nome da fruta não deve ser vazio, e pode conter no máximo 50 caracteres"})
		return
	}

	fruta, err := frutaController.FrutaService.GetFruta(nome)

	if err != nil {
		ctx.JSON(404, gin.H{"erro": "Fruta não encontrada."})
		return
	}

	ctx.JSON(200, fruta)
}

// @Summary Registra uma nova fruta no sistema a partir de suas informaçōes
// @Description Registra uma nova fruta a partir do seu nome e preco
// @ID post-fruta
// @Produce json
// @Param nome path string true "Nome"
// @Param preco path int true "Preco"
// @Success 201 {object} string "Fruta registrada com sucesso"
// @Failure 400 {object} string "Informaçōes inválidas."
// @Failure 409 {object} string "Uma fruta com esse nome já existe."
// @Failure 500 {object} string "Falha ao registrar Fruta. Por favor, tente novamente"
// @Router /criar/fruta [post]
func (frutaController *FrutaController) CreateFruta(ctx *gin.Context) {
	var fruta model.Fruta

	errorBinding := ctx.ShouldBindJSON(&fruta)

	if errorBinding != nil {
		ctx.JSON(400, gin.H{"erro": "Informaçōes inválidas."})
		return
	}

	existingFruta, _ := frutaController.FrutaService.GetFruta(fruta.Nome)

	if existingFruta != nil {
		ctx.JSON(409, gin.H{"erro": "Uma fruta com esse nome já existe."})
		return
	}

	errorCreating := frutaController.FrutaService.CreateFruta(&fruta)

	if errorCreating != nil {
		ctx.JSON(500, gin.H{"erro": "Falha ao registrar Fruta. Por favor, tente novamente"})
		return
	}

	ctx.JSON(201, gin.H{"mensagem": "Fruta registrada com sucesso."})
}

// @Summary Atualiza as informações de uma fruta pelo nome
// @Description Atualiza as informações de uma fruta a partir do seu nome
// @ID update-fruta
// @Produce json
// @Param nome path string true "nome"
// @Param fruta body model.Fruta true "Fruta"
// @Success 200 {object} string "Fruta atualizada com sucesso"
// @Failure 400 {object} string "Informaçōes inválidas."
// @Failure 404 {object} string "Fruta não encontrada"
// @Failure 500 {object} string "Falha ao atualizar a Fruta. Por favor, tente novamente"
// @Router /atualizar/fruta [put]
func (frutaController *FrutaController) UpdateFruta(ctx *gin.Context) {
	var fruta model.Fruta

	if err := ctx.ShouldBindJSON(&fruta); err != nil {
		ctx.JSON(400, gin.H{"erro": "Informaçōes inválidas."})
		return
	}

	existingFruta, err := frutaController.FrutaService.GetFruta(fruta.Nome)

	if err != nil || existingFruta == nil {
		ctx.JSON(404, gin.H{"erro": "Fruta não encontrada."})
		return
	}

	err = frutaController.FrutaService.UpdateFruta(&fruta)

	if err != nil {
		ctx.JSON(500, gin.H{"erro": "Falha ao atualizar a Fruta. Por favor, tente novamente"})
		return
	}

	ctx.JSON(200, gin.H{"mensagem": "Fruta atualizada com sucesso"})
}
