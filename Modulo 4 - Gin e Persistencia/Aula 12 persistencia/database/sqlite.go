// Pacote database encapsula a configuração e conexão com o SQLite.
package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

// DB é a referência global utilizada pelas demais camadas.
var DB *sql.DB

// InitDB abre a conexão com o SQLite e cria a tabela de usuários caso não exista.
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS usuarios(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nome TEXT NOT NULL,
		email TEXT NOT NULL
	);
	CREATE UNIQUE INDEX IF NOT EXISTS idx_usuarios_email on usuarios(email);
	`)
	if err != nil {
		log.Fatal(err)
	}
}
