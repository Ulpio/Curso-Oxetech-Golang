// Exemplo inicial de servidor HTTP usando o framework Gin.
package main

import "aula11-oxetech/routes"

// main é o ponto de entrada e apenas delega a configuração das rotas.
func main() {
	routes.SetupRotas()
}
