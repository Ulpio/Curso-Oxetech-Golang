// Pacote routes centraliza a configuração das rotas HTTP do módulo.
package routes

import (
	"aula11-oxetech/handlers"
	"aula11-oxetech/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// SetupRotas cria o roteador do Gin, registra as rotas e inicia o servidor.
func SetupRotas() {
	// Popula a base em memória com um usuário inicial.
	models.Usuarios = append(models.Usuarios, models.Usuario{ID: 1, Nome: "Ulpio", Idade: 23})
	models.IdCounter = 2

	roteador := gin.Default()

	// Rotas REST básicas.
	roteador.GET("/usuarios", handlers.GetUsuarios)
	roteador.POST("/usuarios", handlers.CreateUsuarios)
	roteador.PUT("/usuarios/:id", handlers.UpdateUsuario)
	roteador.DELETE("/usuarios/:id", handlers.DeleteUsuario)

	fmt.Println("Servidor rodando em http://localhost:8080")
	if err := roteador.Run(":8080"); err != nil {
		panic(err)
	}
}
