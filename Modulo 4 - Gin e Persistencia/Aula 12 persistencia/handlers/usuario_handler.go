// Pacote handlers concentra o acesso às camadas de banco e modelos.
package handlers

import (
	"aula11-oxetech/database"
	"aula11-oxetech/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetUsuarios lista usuários aplicando filtros opcionais de nome e email.
func GetUsuarios(c *gin.Context) {
	nome := c.Query("nome")
	email := c.Query("email")

	query := "SELECT id,nome,email FROM usuarios"

	var args []any
	var where []string

	if nome != "" {
		where = append(where, "nome LIKE ?")
		args = append(args, "%"+nome+"%")
	}
	if email != "" {
		where = append(where, "email LIKE ?")
		args = append(args, "%"+email+"%")
	}

	if len(where) > 0 {
		query += " WHERE " + strings.Join(where, " AND ")
	}

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var saida []models.Usuario
	for rows.Next() {
		var u models.Usuario
		if err := rows.Scan(&u.ID, &u.Nome, &u.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		saida = append(saida, u)
	}
	c.JSON(http.StatusOK, saida)
}

// CreateUsuarios cria um novo registro garantindo unicidade do email.
func CreateUsuarios(c *gin.Context) {
	var u models.Usuario
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao ler corpo da requisicao"})
		return
	}

	var existente int
	if err := database.DB.QueryRow(`SELECT COUNT(1) FROM usuarios WHERE email = ?`, u.Email).Scan(&existente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao verificar email", "details": err.Error()})
		return
	}
	if existente > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "email já cadastrado"})
		return
	}

	res, err := database.DB.Exec(`INSERT INTO usuarios(nome,email) VALUES(?,?)`, u.Nome, u.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao criar usuario", "details": err.Error()})
		return
	}
	id, _ := res.LastInsertId()
	u.ID = int(id)
	c.JSON(http.StatusCreated, u)
}

// UpdateUsuario atualiza nome e email do usuário informado em :id.
func UpdateUsuario(c *gin.Context) {
	id := c.Param("id")
	var atualizado models.Usuario
	if err := c.ShouldBindJSON(&atualizado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O corpo da requisicao nao pode ser lido"})
		return
	}

	if _, err := database.DB.Exec("UPDATE usuarios SET nome=?, email=? WHERE id=?", atualizado.Nome, atualizado.Email, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "atualizado"})
}

// DeleteUsuario remove o registro identificado por :id.
func DeleteUsuario(c *gin.Context) {
	id := c.Param("id")

	if _, err := database.DB.Exec("DELETE FROM usuarios WHERE id=?", id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "removido"})
}
