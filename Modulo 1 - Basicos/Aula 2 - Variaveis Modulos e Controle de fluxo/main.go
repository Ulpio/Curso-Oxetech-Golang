package main // Define como o pacote principal do programa

import "fmt" // Importa o pacote fmt para imprimir na tela

func main() {
	// Declaração de variáveis - forma explícita
	var nome string = "João"
	var idade int = 25

	// Declaração de variáveis - forma inferida
	cidade := "São Paulo" // O Go infere que é string
	ativo := true         // O Go infere que é bool

	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Cidade:", cidade)

	// Controle de fluxo - if/else
	if idade >= 18 {
		fmt.Println("É maior de idade")
	} else {
		fmt.Println("É menor de idade")
	}

	// Switch case
	diaSemana := 3
	switch diaSemana {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda-feira")
	case 3:
		fmt.Println("Terça-feira")
	default:
		fmt.Println("Outro dia")
	}

	// If com declaração inline
	if status := ativo; status {
		fmt.Println("Usuário ativo no sistema")
	}
}
