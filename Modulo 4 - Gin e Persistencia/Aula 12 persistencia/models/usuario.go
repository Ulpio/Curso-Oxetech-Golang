// Pacote models contém as estruturas que representam as tabelas do banco.
package models

// Usuario representa a entidade persistida no banco SQLite.
type Usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
