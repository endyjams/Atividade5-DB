package model

// Fruta representa uma fruta do sistema
// @swaggertype "object"
type Fruta struct {
	Nome  string  `json:"nome"`
	Preco float32 `json:"preco"`
}
