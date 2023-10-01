package controller

import (
	"db-api-example/service"

	model "command-line-arguments/Users/endy.jams/Desktop/db-atividades/Atividade5-DB/model/estoque.go"

	"github.com/gin-gonic/gin"
)

type EstoqueController struct {
	EstoqueService *service.EstoqueService
}

// @Summary Busca a quantidade de uma fruta no estoque a partir do nome da fruta
// @Description Retorna a quantidade de uma fruta no estoque a partir do nome da fruta
// @ID get-estoque
// @Produce json
// @Param nomeFruta path string true "nomeFruta"
// @Success 200 {object} model.nomeFruta "Retorna as informaçōes do estoque da Fruta"
// @Failure 400 {object} string "O nome da fruta não deve ser vazio, e pode conter no máximo 50 caracteres"
// @Failure 404 {object} string "Fruta não encontrada em Estoque"
// @Router /estoque/{nomeFruta} [get]
func (estoqueController *EstoqueController) GetEstoque(ctx *gin.Context) {
	nomeFruta := ctx.Param("nomeFruta")

	if nomeFruta == "" || len(nomeFruta) > 50 {
		ctx.JSON(400, gin.H{"erro": "O nome da fruta não deve ser vazio, e pode conter no máximo 50 caracteres"})
		return
	}

	estoque, err := estoqueController.EstoqueService.GetEstoque(nomeFruta)

	if err != nil {
		ctx.JSON(404, gin.H{"erro": "Fruta não encontrada em Estoque."})
		return
	}

	ctx.JSON(200, estoque)
}

// @Summary Registra uma nova fruta em estoque a partir do nome do fornecedor, nome da fruta e quantidade fornecida
// @Description Registra um novo estoque de fruta a partir do nome do fornecedor, nome da fruta e quantidade fornecida
// @ID post-estoque
// @Produce json
// @Param nomeFruta path string true "NomeFruta"
// @Param nomeFornecedor path string true "NomeFornecedor"
// @Param quantidade path int true "Quantidade"
// @Success 201 {object} string "Estoque de fruta registrado com sucesso"
// @Failure 400 {object} string "Informaçōes inválidas."
// @Failure 409 {object} string "Uma fruta com esse nome já existe."
// @Failure 500 {object} string "Falha ao registrar Fruta. Por favor, tente novamente"
// @Router /fruta [post]
func (estoqueController *EstoqueController) CreateEstoque(ctx *gin.Context) {
	var estoque model.Estoque

	errorBinding := ctx.ShouldBindJSON(&estoque)

	if errorBinding != nil {
		ctx.JSON(400, gin.H{"erro": "Informaçōes inválidas."})
		return
	}

	errorCreating := estoqueController.EstoqueService.CreateEstoque(&estoque)

	if errorCreating != nil {
		ctx.JSON(500, gin.H{"erro": "Falha ao registrar Estoque. Por favor, tente novamente"})
		return
	}

	ctx.JSON(201, gin.H{"mensagem": "Estoque registrado com sucesso."})
}
