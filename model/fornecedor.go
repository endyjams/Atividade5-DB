package model

// Fornecedor representa o fornecedor referente ao estoque de frutas
// @swaggertype "object"
type Fornecedor struct {
	nome     string `json:"nome"`
	telefone string `json:"telefone"`
}
