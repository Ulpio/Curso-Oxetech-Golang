// Programa de exemplo que expõe uma API REST básica usando a biblioteca net/http.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Usuario representa o modelo usado na API em memória.
type Usuario struct {
	ID    uint
	Nome  string
	Idade int
}

// usuarios simula um armazenamento em memória para os dados.
var usuarios []Usuario

// idCounter controla a geração incremental de IDs.
var idCounter = 1

// getUsuarios devolve todos os usuários cadastrados.
func getUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Resposta será JSON.
	json.NewEncoder(w).Encode(usuarios)                // Serializa o slice em JSON.
}

// createUsuarios cria um novo usuário a partir do corpo da requisição.
func createUsuarios(w http.ResponseWriter, r *http.Request) {
	var novo Usuario // Estrutura que receberá os dados enviados.

	if err := json.NewDecoder(r.Body).Decode(&novo); err != nil { // Faz o parse do JSON.
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	novo.ID = uint(idCounter)         // Define o ID sequencial.
	idCounter++                       // Incrementa o próximo ID.
	usuarios = append(usuarios, novo) // Persiste em memória.

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // Responde 201 Created.
	json.NewEncoder(w).Encode(novo)
}

// updateUsuario atualiza um usuário existente usando o ID da rota.
func updateUsuario(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/usuarios/") // Extrai o ID da URL.
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	var atualizado Usuario
	if err := json.NewDecoder(r.Body).Decode(&atualizado); err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	for i, u := range usuarios {
		if u.ID == uint(id) {
			atualizado.ID = uint(id) // Garante a manutenção do ID original.
			usuarios[i] = atualizado // Substitui o registro no slice.
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(atualizado)
			return
		}
	}
	http.Error(w, "Usuario não encontrado", http.StatusNotFound)
}

// deleteUsuario remove um usuário da lista pelo ID informado na rota.
func deleteUsuario(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	for i, u := range usuarios {
		if u.ID == uint(id) {
			usuarios = append(usuarios[:i], usuarios[i+1:]...) // Remove do slice.
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Usuario não encontrado", http.StatusNotFound)
}

func main() {
	usuarios = append(usuarios, Usuario{ID: 1, Nome: "Ulpio", Idade: 23})
	idCounter = 2

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUsuarios(w, r)
		case http.MethodPost:
			createUsuarios(w, r)
		default:
			http.Error(w, "Metodo nao permitido", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/usuarios/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			updateUsuario(w, r)
		case http.MethodDelete:
			deleteUsuario(w, r)
		default:
			http.Error(w, "Metodo nao permitido", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
