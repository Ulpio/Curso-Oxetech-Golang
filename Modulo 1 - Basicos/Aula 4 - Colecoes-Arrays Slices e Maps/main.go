package main // Define como o pacote principal do programa

import "fmt" // Importa o pacote fmt para imprimir na tela

func main() {
	// ARRAYS - Tamanho fixo
	fmt.Println("=== ARRAYS ===")

	// Declarando array com tamanho fixo de 3 elementos
	var numeros [3]int
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	fmt.Println("Array de números:", numeros)
	fmt.Println("Tamanho do array:", len(numeros))

	// Array inicializado
	frutas := [4]string{"Maçã", "Banana", "Laranja", "Uva"}
	fmt.Println("Array de frutas:", frutas)

	// Iterando sobre array
	fmt.Println("\nIterando sobre frutas:")
	for indice, fruta := range frutas {
		fmt.Printf("Índice %d: %s\n", indice, fruta)
	}

	// SLICES - Tamanho dinâmico
	fmt.Println("\n=== SLICES ===")

	// Criando slice vazio
	var cores []string
	fmt.Println("Slice vazio:", cores)

	// Adicionando elementos ao slice com append
	cores = append(cores, "Vermelho")
	cores = append(cores, "Azul", "Verde")
	fmt.Println("Slice após append:", cores)
	fmt.Println("Tamanho:", len(cores), "Capacidade:", cap(cores))

	// Criando slice com make
	numeros2 := make([]int, 3, 5) // tamanho 3, capacidade 5
	numeros2[0] = 1
	numeros2[1] = 2
	numeros2[2] = 3
	fmt.Println("Slice criado com make:", numeros2)

	// Slice de um array ou outro slice
	semana := []string{"Dom", "Seg", "Ter", "Qua", "Qui", "Sex", "Sáb"}
	diasUteis := semana[1:6] // do índice 1 até 5 (6 não incluído)
	fmt.Println("Dias úteis:", diasUteis)

	// MAPS - Estrutura chave-valor
	fmt.Println("\n=== MAPS ===")

	// Criando map vazio
	var idades map[string]int
	idades = make(map[string]int)

	// Adicionando elementos
	idades["João"] = 25
	idades["Maria"] = 30
	idades["Pedro"] = 28
	fmt.Println("Map de idades:", idades)

	// Criando map inicializado
	telefones := map[string]string{
		"João":  "1111-1111",
		"Maria": "2222-2222",
		"Pedro": "3333-3333",
	}
	fmt.Println("Map de telefones:", telefones)

	// Acessando valores
	fmt.Println("Telefone de Maria:", telefones["Maria"])

	// Verificando se chave existe
	tel, existe := telefones["Ana"]
	if existe {
		fmt.Println("Telefone de Ana:", tel)
	} else {
		fmt.Println("Ana não está no mapa")
	}

	// Iterando sobre map
	fmt.Println("\nIterando sobre telefones:")
	for nome, telefone := range telefones {
		fmt.Printf("%s: %s\n", nome, telefone)
	}

	// Deletando elemento do map
	delete(telefones, "Pedro")
	fmt.Println("Telefones após deletar Pedro:", telefones)

	// Exemplo prático: contador de palavras
	fmt.Println("\n=== EXEMPLO PRÁTICO ===")
	palavras := []string{"go", "python", "go", "java", "python", "go"}
	contador := make(map[string]int)

	for _, palavra := range palavras {
		contador[palavra]++
	}

	fmt.Println("Contador de palavras:", contador)
}
