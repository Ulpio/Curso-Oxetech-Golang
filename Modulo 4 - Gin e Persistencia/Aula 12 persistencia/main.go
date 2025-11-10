// Aplicação Gin com persistência em SQLite utilizando database/sql.
package main

import (
	"aula11-oxetech/database"
	"aula11-oxetech/routes"
)

// main inicializa a conexão com o banco e sobe o servidor HTTP.
func main() {
	database.InitDB()
	routes.SetupRotas()
}
