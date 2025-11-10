// Pacote handlers contém as funções que lidam com as requisições HTTP.
package handlers

import (
	"aula11-oxetech/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUsuarios responde com todos os usuários cadastrados em memória.
func GetUsuarios(c *gin.Context) {
	c.JSON(http.StatusOK, models.Usuarios)
}

// CreateUsuarios lê o JSON da requisição e adiciona um novo usuário à lista.
func CreateUsuarios(c *gin.Context) {
	var novo models.Usuario

	if err := c.ShouldBindJSON(&novo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O corpo da requisicao nao pode ser lido"})
		return
	}

	novo.ID = uint(models.IdCounter)
	models.IdCounter++
	models.Usuarios = append(models.Usuarios, novo)

	c.JSON(http.StatusCreated, novo)

}

// UpdateUsuario atualiza um usuário existente identificado pelo parâmetro :id.
func UpdateUsuario(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}

	var atualizado models.Usuario

	if err := c.ShouldBindJSON(&atualizado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O corpo da requisicao nao pode ser lido"})
		return
	}

	for i, u := range models.Usuarios { // Procuro no Slice
		if u.ID == uint(id) { // Se usuario.ID == id da URI
			atualizado.ID = uint(id)        // Garante que o ID permaneça o mesmo.
			models.Usuarios[i] = atualizado // Atualiza o registro em memória.
			c.JSON(http.StatusOK, atualizado)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "usuario não encontrado"})
}

// DeleteUsuario remove um usuário com o ID informado na rota.
func DeleteUsuario(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}

	for i, u := range models.Usuarios {
		if u.ID == uint(id) {
			models.Usuarios = append(models.Usuarios[:i], models.Usuarios[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "usuario nao encontrado"})
}
