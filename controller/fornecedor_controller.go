package controller

import (
	"Atividade5-DB/model"
	"Atividade5-DB/service"

	"github.com/gin-gonic/gin"
)

type FornecedorController struct {
	FornecedorService *service.FornecedorService
}

// @Summary Busca um fornecedor pelo nome
// @Description Retorna as informaçōes de um fornecedor a partir de seu nome
// @ID get-fornecedor
// @Produce json
// @Param nome path string true "nome"
// @Success 200 {object} model.Fornecedor "Retorna as informaçōes do fornecedor"
// @Failure 400 {object} string "O nome do fornecedor não deve ser vazio, e pode conter no máximo 50 caracteres"
// @Failure 404 {object} string "Fornecedor não encontrado"
// @Router /obter/fornecedor/{nome} [get]
func (fornecedorController *FornecedorController) GetFornecedor(ctx *gin.Context) {
	nome := ctx.Param("nome")

	if nome == "" || len(nome) > 50 {
		ctx.JSON(400, gin.H{"erro": "O nome do fornecedor não deve ser vazio, e pode conter no máximo 50 caracteres"})
		return
	}

	fornecedor, err := fornecedorController.FornecedorService.GetFornecedor(nome)

	if err != nil {
		ctx.JSON(404, gin.H{"erro": "Fornecedor não encontrado."})
		return
	}

	ctx.JSON(200, fornecedor)
}

// @Summary Registra um novo fornecedor no sistema a partir de suas informaçōes
// @Description Registra um novo fornecedor a partir do seu nome e telefone
// @ID post-fornecedor
// @Produce json
// @Param nome path string true "Nome"
// @Param telefone path string true "Telefone"
// @Success 201 {object} string "Fornecedor registrado com sucesso"
// @Failure 400 {object} string "Informaçōes inválidas."
// @Failure 409 {object} string "Um fornecedor com esse nome já existe."
// @Failure 500 {object} string "Falha ao registrar Fornecedor. Por favor, tente novamente"
// @Router /criar/fornecedor [post]
func (fornecedorController *FornecedorController) CreateFornecedor(ctx *gin.Context) {
	var fornecedor model.Fornecedor

	errorBinding := ctx.ShouldBindJSON(&fornecedor)

	if errorBinding != nil {
		ctx.JSON(400, gin.H{"erro": "Informaçōes inválidas."})
		return
	}

	existingFornecedor, _ := fornecedorController.FornecedorService.GetFornecedor(fornecedor.Nome)

	if existingFornecedor != nil {
		ctx.JSON(409, gin.H{"erro": "Um fornecedor com esse nome já existe."})
		return
	}

	errorCreating := fornecedorController.FornecedorService.CreateFornecedor(&fornecedor)

	if errorCreating != nil {
		ctx.JSON(500, gin.H{"erro": "Falha ao registrar Fornecedor. Por favor, tente novamente"})
		return
	}

	ctx.JSON(201, gin.H{"mensagem": "Fornecedor registrado com sucesso."})
}

// @Summary Atualiza as informaçōes de um fornecedor existente
// @Description Atualiza as informaçōes de um fornecedor a partir do seu nome
// @ID update-fornecedor
// @Produce json
// @Param nome path string true "nome"
// @Param telefone path string true "telefone"
// @Param fornecedor body model.Fornecedor true "Fornecedor"
// @Success 200 {object} string "Fornecedor atualizado com sucesso"
// @Failure 400 {object} string "Informaçōes inválidas."
// @Failure 404 {object} string "Fornecedor não encontrado"
// @Failure 500 {object} string "Falha ao atualizar o fornecedor. Por favor, tente novamente"
// @Router /atualizar/fornecedor [put]

func (fornecedorController *FornecedorController) UpdateFornecedor(ctx *gin.Context) {
	var fornecedor model.Fornecedor

	if err := ctx.ShouldBindJSON(&fornecedor); err != nil {
		ctx.JSON(400, gin.H{"erro": "Informaçōes inválidas."})
		return
	}

	existingFornecedor, err := fornecedorController.FornecedorService.GetFornecedor(fornecedor.Nome)

	if err != nil || existingFornecedor == nil {
		ctx.JSON(404, gin.H{"erro": "Fornecedor não encontrado."})
		return
	}

	err = fornecedorController.FornecedorService.UpdateFornecedor(&fornecedor)

	if err != nil {
		ctx.JSON(500, gin.H{"erro": "Falha ao atualizar o fornecedor. Por favor, tente novamente"})
		return
	}

	ctx.JSON(200, gin.H{"mensagem": "Fornecedor atualizado com sucesso"})
}

// @Summary Deleta as informaçōes de um fornecedor existente
// @Description Deleta as informaçōes de um fornecedor a partir de suas informaçōes
// @ID delete-fornecedor
// @Produce json
// @Param nome path string true "nome"
// @Param telefone path string true "nome"
// @Param fornecedor body model.Fornecedor true "Fornecedor"
// @Success 200 {object} string "Fornecedor deletado com sucesso"
// @Failure 400 {object} string "Informaçōes inválidas."
// @Failure 404 {object} string "Fornecedor não encontrado"
// @Failure 500 {object} string "Falha ao deletar o fornecedor. Por favor, tente novamente"
// @Router /deletar/fornecedor [delete]

func (fornecedorController *FornecedorController) DeleteFornecedor(ctx *gin.Context) {
	var fornecedor model.Fornecedor

	if err := ctx.ShouldBindJSON(&fornecedor); err != nil {
		ctx.JSON(400, gin.H{"erro": "Informaçōes inválidas."})
		return
	}

	existingFornecedor, err := fornecedorController.FornecedorService.GetFornecedor(fornecedor.Nome)

	if err != nil || existingFornecedor == nil {
		ctx.JSON(404, gin.H{"erro": "Fornecedor não encontrado."})
		return
	}

	err = fornecedorController.FornecedorService.DeleteFornecedor(&fornecedor)

	if err != nil {
		ctx.JSON(500, gin.H{"erro": "Falha ao deletar o fornecedor. Por favor, tente novamente"})
		return
	}

	ctx.JSON(200, gin.H{"mensagem": "Fornecedor deletado com sucesso"})
}
