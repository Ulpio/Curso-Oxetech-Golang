package main // Define como o pacote principal do programa

import "fmt" // Importa o pacote fmt para imprimir na tela

// Struct Usuario
type Usuario struct {
	Nome  string
	Idade int
}

// Método com receptor por valor - NÃO modifica o original
func (u Usuario) ImprimeDados() {
	fmt.Printf("Nome: %s, Idade: %d\n", u.Nome, u.Idade)
}

// Método com receptor por ponteiro - MODIFICA o original
func (u *Usuario) Aniversario() {
	u.Idade++ // Incrementa a idade do usuário original
}

// Método com receptor por ponteiro para alterar nome
func (u *Usuario) AlterarNome(novoNome string) {
	u.Nome = novoNome
}

// Struct Circulo
type Circulo struct {
	Raio float64
}

// Método com receptor por valor - calcula área
func (c Circulo) Area() float64 {
	return 3.14 * c.Raio * c.Raio
}

// Método com receptor por ponteiro - altera o raio
func (c *Circulo) Dobrar() {
	c.Raio = c.Raio * 2
}

// Função que recebe ponteiro como parâmetro
func incrementarIdade(u *Usuario) {
	u.Idade++
}

// Função que recebe valor (não modifica o original)
func tentarIncrementar(u Usuario) {
	u.Idade++ // Isso não afeta o original
}

func main() {
	fmt.Println("=== PONTEIROS ===")

	// Criando variáveis
	x := 10
	fmt.Println("Valor de x:", x)

	// Criando ponteiro para x
	p := &x // & retorna o endereço de memória
	fmt.Println("Endereço de x:", p)
	fmt.Println("Valor apontado por p:", *p) // * desreferencia o ponteiro

	// Modificando através do ponteiro
	*p = 20
	fmt.Println("Novo valor de x:", x) // x foi modificado!

	// MÉTODOS COM RECEPTORES
	fmt.Println("\n=== MÉTODOS COM RECEPTORES ===")

	usuario := Usuario{Nome: "Ulpio", Idade: 23}
	fmt.Println("\nUsuário inicial:")
	usuario.ImprimeDados()

	// Chamando método com receptor por ponteiro
	fmt.Println("\nDepois do aniversário:")
	usuario.Aniversario() // Go automaticamente converte para &usuario
	usuario.ImprimeDados()

	// Alterando nome
	usuario.AlterarNome("Ulpio Silva")
	fmt.Println("\nDepois de alterar o nome:")
	usuario.ImprimeDados()

	// DIFERENÇA ENTRE VALOR E PONTEIRO
	fmt.Println("\n=== VALOR vs PONTEIRO ===")

	usuario2 := Usuario{Nome: "Maria", Idade: 25}
	fmt.Println("\nUsuário 2 inicial:")
	usuario2.ImprimeDados()

	// Tentando incrementar por valor (não funciona)
	tentarIncrementar(usuario2)
	fmt.Println("Após tentarIncrementar (por valor):")
	usuario2.ImprimeDados() // Idade não mudou!

	// Incrementando por ponteiro (funciona)
	incrementarIdade(&usuario2)
	fmt.Println("Após incrementarIdade (por ponteiro):")
	usuario2.ImprimeDados() // Idade mudou!

	// MÉTODOS COM STRUCT CIRCULO
	fmt.Println("\n=== EXEMPLO COM CÍRCULO ===")

	circulo := Circulo{Raio: 5}
	fmt.Printf("Círculo com raio %.2f\n", circulo.Raio)
	fmt.Printf("Área: %.2f\n", circulo.Area())

	// Dobrando o raio
	circulo.Dobrar()
	fmt.Printf("\nApós dobrar o raio: %.2f\n", circulo.Raio)
	fmt.Printf("Nova área: %.2f\n", circulo.Area())

	// QUANDO USAR PONTEIROS?
	fmt.Println("\n=== QUANDO USAR PONTEIROS? ===")
	fmt.Println("1. Quando você precisa MODIFICAR o valor original")
	fmt.Println("2. Quando a struct é GRANDE (evita cópia)")
	fmt.Println("3. Para CONSISTÊNCIA (se alguns métodos usam *, todos devem usar)")
}
