package main // Define como o pacote principal do programa

import "fmt" // Importa o pacote fmt para imprimir na tela

// Struct simples - representa uma Pessoa
type Pessoa struct {
	Nome  string
	Idade int
	Email string
}

// Struct para Endereço
type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

// Struct com composição - Cliente possui um Endereço
type Cliente struct {
	Nome     string
	CPF      string
	Endereco Endereco // Composição: Cliente tem um Endereço
}

// Struct com composição embutida (embedded)
type Funcionario struct {
	Pessoa  // Composição embutida - Funcionario "é uma" Pessoa
	Cargo   string
	Salario float64
}

// Struct para Produto
type Produto struct {
	Nome  string
	Preco float64
}

// Função que retorna uma struct
func novaPessoa(nome string, idade int) Pessoa {
	return Pessoa{
		Nome:  nome,
		Idade: idade,
		Email: nome + "@email.com",
	}
}

func main() {
	// Criando struct de forma literal
	fmt.Println("=== STRUCT SIMPLES ===")
	pessoa1 := Pessoa{
		Nome:  "João Silva",
		Idade: 30,
		Email: "joao@email.com",
	}
	fmt.Println("Pessoa 1:", pessoa1)
	fmt.Println("Nome:", pessoa1.Nome)
	fmt.Println("Idade:", pessoa1.Idade)

	// Criando struct sem nomear os campos (deve seguir a ordem)
	pessoa2 := Pessoa{"Maria Santos", 25, "maria@email.com"}
	fmt.Println("\nPessoa 2:", pessoa2)

	// Criando struct usando função construtora
	pessoa3 := novaPessoa("Pedro", 28)
	fmt.Println("\nPessoa 3:", pessoa3)

	// Struct vazia (valores zero)
	var pessoa4 Pessoa
	fmt.Println("\nPessoa 4 (vazia):", pessoa4)

	// Modificando valores
	pessoa4.Nome = "Ana Costa"
	pessoa4.Idade = 35
	fmt.Println("Pessoa 4 (modificada):", pessoa4)

	// COMPOSIÇÃO
	fmt.Println("\n=== COMPOSIÇÃO ===")

	// Criando cliente com endereço
	cliente := Cliente{
		Nome: "Carlos Oliveira",
		CPF:  "123.456.789-00",
		Endereco: Endereco{
			Rua:    "Av. Paulista",
			Numero: 1000,
			Cidade: "São Paulo",
			Estado: "SP",
		},
	}
	fmt.Println("Cliente:", cliente)
	fmt.Println("Cidade do cliente:", cliente.Endereco.Cidade)

	// COMPOSIÇÃO EMBUTIDA
	fmt.Println("\n=== COMPOSIÇÃO EMBUTIDA ===")

	funcionario := Funcionario{
		Pessoa: Pessoa{
			Nome:  "Fernanda Lima",
			Idade: 27,
			Email: "fernanda@empresa.com",
		},
		Cargo:   "Desenvolvedora",
		Salario: 8000.50,
	}

	fmt.Println("Funcionário:", funcionario)
	// Com composição embutida, podemos acessar campos diretamente
	fmt.Println("Nome do funcionário:", funcionario.Nome) // Acesso direto!
	fmt.Println("Cargo:", funcionario.Cargo)
	fmt.Println("Salário:", funcionario.Salario)

	// SLICE DE STRUCTS
	fmt.Println("\n=== SLICE DE STRUCTS ===")

	produtos := []Produto{
		{Nome: "Notebook", Preco: 3500.00},
		{Nome: "Mouse", Preco: 50.00},
		{Nome: "Teclado", Preco: 150.00},
	}

	fmt.Println("Lista de produtos:")
	for i, produto := range produtos {
		fmt.Printf("%d. %s - R$ %.2f\n", i+1, produto.Nome, produto.Preco)
	}

	// Calculando total
	var total float64
	for _, produto := range produtos {
		total += produto.Preco
	}
	fmt.Printf("\nTotal: R$ %.2f\n", total)
}
