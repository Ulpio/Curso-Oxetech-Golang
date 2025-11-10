// Pacote models define as estruturas compartilhadas entre as camadas da aplicação.
package models

// Usuarios representa o armazenamento em memória dos registros de usuário.
var Usuarios []Usuario

// IdCounter fornece IDs incrementais para novos usuários.
var IdCounter = 1

// Usuario descreve o payload aceito/retornado nas rotas.
type Usuario struct {
	ID    uint   `json:"id"`
	Nome  string `json:"nome" binding:"required"`
	Idade int    `json:"idade" binding:"required"`
}
