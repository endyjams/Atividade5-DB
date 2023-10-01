package model

// Fornecedor representa o fornecedor referente ao estoque de frutas
// @swaggertype "object"
type Fornecedor struct {
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
}
