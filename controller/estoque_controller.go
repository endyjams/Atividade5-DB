package controller

import (
	"Atividade5-DB/service"

	"Atividade5-DB/model"

	"github.com/gin-gonic/gin"
)

type EstoqueController struct {
	EstoqueService *service.EstoqueService
}

// @Summary Busca a quantidade de uma fruta fornecida por um fornecedor no estoque a partir do nome da fruta e nome do fornecedor
// @Description Busca a quantidade de uma fruta fornecida por um fornecedor no estoque a partir do nome da fruta e nome do fornecedor
// @ID get-estoque
// @Produce json
// @Param nomeFruta path string true "nomeFruta"
// @Param nomeFornecedor path string true "nomeFornecedor"
// @Success 200 {object} model.Estoque "Retorna as informaçōes do estoque da Fruta"
// @Failure 400 {object} string "O nome da fruta não deve ser vazio, e pode conter no máximo 50 caracteres"
// @Failure 404 {object} string "Fruta não encontrada em Estoque"
// @Router /obter/estoque/{nomeFruta}/{nomeFornecedor} [get]
func (estoqueController *EstoqueController) GetEstoque(ctx *gin.Context) {
	nomeFruta := ctx.Param("nomeFruta")
	nomeFornecedor := ctx.Param("nomeFornecedor")

	if nomeFruta == "" || len(nomeFruta) > 50 || nomeFornecedor == "" || len(nomeFornecedor) > 50 {
		ctx.JSON(400, gin.H{"erro": "O nome da fruta e do fornecedor não devem ser vazios, e podem conter no máximo 50 caracteres"})
		return
	}

	estoque, err := estoqueController.EstoqueService.GetEstoque(nomeFruta, nomeFornecedor)

	if err != nil {
		ctx.JSON(404, gin.H{"erro": "Fruta não encontrada no estoque do fornecedor."})
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
// @Router /criar/estoque [post]
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

// @Summary Atualiza a quantidade de uma fruta no estoque
// @Description Atualiza a quantidade de uma fruta no estoque a partir das novas informaçōes sobre o estoque
// @ID update-estoque
// @Produce json
// @Param estoque body model.Estoque true "Estoque"
// @Success 200 {object} string "Estoque atualizado com sucesso"
// @Failure 400 {object} string "Informaçōes inválidas."
// @Failure 404 {object} string "Estoque não encontrado"
// @Failure 500 {object} string "Falha ao atualizar o estoque. Por favor, tente novamente"
// @Router /atualizar/estoque [put]
func (estoqueController *EstoqueController) UpdateEstoque(ctx *gin.Context) {
	var estoque model.Estoque

	errorBinding := ctx.ShouldBindJSON(&estoque)

	if errorBinding != nil {
		ctx.JSON(400, gin.H{"erro": "Informaçōes inválidas."})
		return
	}

	existingEstoque, err := estoqueController.EstoqueService.GetEstoque(estoque.NomeFruta, estoque.NomeFornecedor)

	if err != nil || existingEstoque == nil {
		ctx.JSON(404, gin.H{"erro": "Estoque não encontrado."})
		return
	}

	err = estoqueController.EstoqueService.UpdateEstoque(&estoque)

	if err != nil {
		ctx.JSON(500, gin.H{"erro": "Falha ao atualizar o estoque. Por favor, tente novamente"})
		return
	}

	ctx.JSON(200, gin.H{"mensagem": "Estoque atualizado com sucesso"})
}

// @Summary Deleta um estoque a partir de suas informaçōes
// @Description Deleta o estoque de uma fruta associada a um fornecedor
// @ID delete-estoque
// @Produce json
// @Param estoque body model.Estoque true "Estoque"
// @Success 200 {object} string "Estoque deletado com sucesso"
// @Failure 400 {object} string "Informaçōes inválidas."
// @Failure 404 {object} string "Estoque não encontrado"
// @Failure 500 {object} string "Falha ao deletar o estoque. Por favor, tente novamente"
// @Router /deletar/estoque [delete]
func (estoqueController *EstoqueController) DeleteEstoque(ctx *gin.Context) {
	var estoque model.Estoque

	errorBinding := ctx.ShouldBindJSON(&estoque)

	if errorBinding != nil {
		ctx.JSON(400, gin.H{"erro": "Informaçōes inválidas."})
		return
	}

	existingEstoque, err := estoqueController.EstoqueService.GetEstoque(estoque.NomeFruta, estoque.NomeFornecedor)

	if err != nil || existingEstoque == nil {
		ctx.JSON(404, gin.H{"erro": "Estoque não encontrado."})
		return
	}

	err = estoqueController.EstoqueService.DeleteEstoque(&estoque)

	if err != nil {
		ctx.JSON(500, gin.H{"erro": "Falha ao deletar o estoque. Por favor, tente novamente"})
		return
	}

	ctx.JSON(200, gin.H{"mensagem": "Estoque deletado com sucesso"})
}

// @Summary Preenche o estoque de alguma fruta que esgotou adicionando 10
// @Description Preencher o estoque de fruta esgotada com 10
// @ID fill-estoque
// @Produce json
// @Success 201 {object} string "Estoque preenchido com sucesso"
// @Failure 500 {object} string "Falha ao preencher o estoque. Não há frutas esgotadas"
// @Router /preencher/estoque [post]
func (estoqueController *EstoqueController) FillEstoque(ctx *gin.Context) {
	var err = estoqueController.EstoqueService.FillEstoque()

	if err != nil {
		ctx.JSON(500, gin.H{"erro": "Falha ao preencher estoque. Provavelmente não há frutas esgotadas"})
		return
	}

	ctx.JSON(200, gin.H{"mensagem": "Estoque preenchido com sucesso"})
}
