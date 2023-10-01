package model

// Estoque representa o estoque referente a uma fruta do sistema
// @swaggertype "object"
type Estoque struct {
	NomeFruta      string `json:"nomeFruta"`
	NomeFornecedor string `json:"nomeFornecedor"`
	Quantidade     int32  `json:"quantidade"`
}
