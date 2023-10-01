package model

// Fruta representa uma fruta do sistema
// @swaggertype "object"
type Fruta struct {
	nome  string  `json:"nome"`
	preco float32 `json:"preco"`
}
