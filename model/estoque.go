package model

// Estoque representa o estoque referente a uma fruta do sistema
// @swaggertype "object"
type Estoque struct {
	nomeFruta      string `json:"nomeFruta"`
	nomeFornecedor string `json:"nomeFornecedor"`
	quantidade     int32  `json:"quantidade"`
}
