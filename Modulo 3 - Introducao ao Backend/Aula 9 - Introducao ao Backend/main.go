// Exemplo básico de servidor HTTP com rotas simples usando net/http.
package main

import (
	"fmt"
	"net/http"
)

// main configura as rotas HTTP e inicia um servidor simples.
func main() {
	// Rota raiz: retorna uma mensagem de boas-vindas.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bem vindo ao primeiro servidor do curso!!")
	})

	// Rota /alunos: demonstra outra resposta estática.
	http.HandleFunc("/alunos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Os alunos do oxetech são desenrolados!")
	})

	// Rota /metodos: exemplo de uso de switch para diferenciar verbos HTTP.
	http.HandleFunc("/metodos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "O método utilizado foi o GET!")
		case http.MethodPost:
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "O método utilizado foi o POST!")
		default:
			// Para métodos não tratados, retornamos status 405 (Method Not Allowed).
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Método %s não suportado", r.Method)
		}
	})

	// ListenAndServe inicia o servidor HTTP na porta 8080;
	// o segundo parâmetro nil indica que usaremos o DefaultServeMux.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
