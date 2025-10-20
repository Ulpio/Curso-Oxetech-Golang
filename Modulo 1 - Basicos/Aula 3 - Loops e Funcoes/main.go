package main // Define como o pacote principal do programa

import "fmt" // Importa o pacote fmt para imprimir na tela

// Função simples que não retorna nada
func saudar(nome string) {
	fmt.Println("Olá,", nome)
}

// Função que retorna um valor
func somar(a int, b int) int {
	return a + b
}

// Função que retorna múltiplos valores
func calcular(a int, b int) (int, int, int) {
	soma := a + b
	subtracao := a - b
	multiplicacao := a * b
	return soma, subtracao, multiplicacao
}

// Função com retorno nomeado
func dividir(a, b float64) (resultado float64, erro string) {
	if b == 0 {
		erro = "Divisão por zero não é permitida"
		return
	}
	resultado = a / b
	return
}

func main() {
	// Loop for tradicional
	fmt.Println("Loop for tradicional:")
	for i := 1; i <= 5; i++ {
		fmt.Println("Contador:", i)
	}

	// Loop while (usando for)
	fmt.Println("\nLoop while:")
	contador := 0
	for contador < 3 {
		fmt.Println("Contador while:", contador)
		contador++
	}

	// Loop infinito com break
	fmt.Println("\nLoop com break:")
	numero := 0
	for {
		if numero >= 3 {
			break // Sai do loop
		}
		fmt.Println("Número:", numero)
		numero++
	}

	// Loop com continue
	fmt.Println("\nLoop com continue:")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // Pula a iteração quando i == 2
		}
		fmt.Println("Valor:", i)
	}

	// Chamando funções
	fmt.Println("\nChamando funções:")
	saudar("Maria")

	resultado := somar(10, 5)
	fmt.Println("Soma:", resultado)

	soma, sub, mult := calcular(10, 3)
	fmt.Println("Soma:", soma, "Subtração:", sub, "Multiplicação:", mult)

	res, err := dividir(10, 2)
	if err != "" {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Divisão:", res)
	}
}
