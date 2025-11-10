// Pacote routes centraliza a configuração das rotas HTTP do módulo com persistência.
package routes

import (
	"aula11-oxetech/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

// SetupRotas cria o roteador do Gin, associa as rotas e inicia o servidor.
func SetupRotas() {
	roteador := gin.Default()

	// Rotas REST conectadas ao banco SQLite.
	roteador.GET("/usuarios", handlers.GetUsuarios)
	roteador.POST("/usuarios", handlers.CreateUsuarios)
	roteador.PUT("/usuarios/:id", handlers.UpdateUsuario)
	roteador.DELETE("/usuarios/:id", handlers.DeleteUsuario)

	fmt.Println("Servidor rodando em http://localhost:8080")
	if err := roteador.Run(":8080"); err != nil {
		panic(err)
	}
}
