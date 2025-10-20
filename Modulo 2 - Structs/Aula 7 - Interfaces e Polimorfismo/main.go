package main // Define como o pacote principal do programa

import "fmt" // Importa o pacote fmt para imprimir na tela

// Interface Desenhavel - define um contrato
// Qualquer tipo que implementar o método Desenhar() é um Desenhavel
type Desenhavel interface {
	Desenhar()
}

// Interface para calcular área
type FormaGeometrica interface {
	Area() float64
	Perimetro() float64
}

// Struct Circulo
type Circulo struct {
	Raio float64
}

// Circulo implementa Desenhavel
func (c Circulo) Desenhar() {
	fmt.Printf("🔵 Desenhando um círculo com raio %.2f\n", c.Raio)
}

// Circulo implementa FormaGeometrica
func (c Circulo) Area() float64 {
	return 3.14 * c.Raio * c.Raio
}

func (c Circulo) Perimetro() float64 {
	return 2 * 3.14 * c.Raio
}

// Struct Retangulo
type Retangulo struct {
	Altura  float64
	Largura float64
}

// Retangulo implementa Desenhavel
func (r Retangulo) Desenhar() {
	fmt.Printf("◻️  Desenhando um retângulo %.2f x %.2f\n", r.Altura, r.Largura)
}

// Retangulo implementa FormaGeometrica
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Altura + r.Largura)
}

// Struct Triangulo
type Triangulo struct {
	Base   float64
	Altura float64
}

// Triangulo implementa Desenhavel
func (t Triangulo) Desenhar() {
	fmt.Printf("🔺 Desenhando um triângulo com base %.2f e altura %.2f\n", t.Base, t.Altura)
}

// Triangulo implementa FormaGeometrica
func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

func (t Triangulo) Perimetro() float64 {
	// Simplificação: assumindo triângulo equilátero
	return t.Base * 3
}

// Função que aceita qualquer Desenhavel (POLIMORFISMO)
func MostrarDesenho(d Desenhavel) {
	d.Desenhar()
}

// Função que aceita qualquer FormaGeometrica
func ImprimirInformacoes(f FormaGeometrica) {
	fmt.Printf("  Área: %.2f\n", f.Area())
	fmt.Printf("  Perímetro: %.2f\n", f.Perimetro())
}

// Type assertion - verificar o tipo concreto
func ConferirTipo(d Desenhavel) {
	// Usando type assertion com verificação
	if circulo, ok := d.(Circulo); ok {
		fmt.Printf("✓ É um Círculo com raio %.2f\n", circulo.Raio)
	}
	if retangulo, ok := d.(Retangulo); ok {
		fmt.Printf("✓ É um Retângulo %.2fx%.2f\n", retangulo.Altura, retangulo.Largura)
	}
	if triangulo, ok := d.(Triangulo); ok {
		fmt.Printf("✓ É um Triângulo com base %.2f\n", triangulo.Base)
	}
}

// Type switch - outra forma de verificar tipos
func IdentificarForma(d Desenhavel) {
	switch forma := d.(type) {
	case Circulo:
		fmt.Printf("Type switch: Círculo com raio %.2f\n", forma.Raio)
	case Retangulo:
		fmt.Printf("Type switch: Retângulo %.2fx%.2f\n", forma.Altura, forma.Largura)
	case Triangulo:
		fmt.Printf("Type switch: Triângulo com base %.2f\n", forma.Base)
	default:
		fmt.Println("Type switch: Forma desconhecida")
	}
}

func main() {
	fmt.Println("=== INTERFACES E POLIMORFISMO ===")

	// Criando diferentes formas
	circulo := Circulo{Raio: 5}
	retangulo := Retangulo{Altura: 4, Largura: 6}
	triangulo := Triangulo{Base: 6, Altura: 4}

	// POLIMORFISMO - a mesma função aceita diferentes tipos
	fmt.Println("--- Desenhando formas ---")
	MostrarDesenho(circulo)
	MostrarDesenho(retangulo)
	MostrarDesenho(triangulo)

	// SLICE DE INTERFACES
	fmt.Println("\n--- Slice de Desenháveis ---")
	formas := []Desenhavel{circulo, retangulo, triangulo}

	for i, forma := range formas {
		fmt.Printf("\nForma %d:\n", i+1)
		forma.Desenhar()
	}

	// TYPE ASSERTION
	fmt.Println("\n--- Type Assertion ---")
	ConferirTipo(circulo)
	ConferirTipo(retangulo)
	ConferirTipo(triangulo)

	// TYPE SWITCH
	fmt.Println("\n--- Type Switch ---")
	IdentificarForma(circulo)
	IdentificarForma(retangulo)
	IdentificarForma(triangulo)

	// MÚLTIPLAS INTERFACES
	fmt.Println("\n--- Formas Geométricas (Interface FormaGeometrica) ---")

	// Slice de FormaGeometrica
	geometricas := []FormaGeometrica{circulo, retangulo, triangulo}

	for i, forma := range geometricas {
		fmt.Printf("\nForma Geométrica %d:\n", i+1)
		ImprimirInformacoes(forma)
	}

	// INTERFACE VAZIA
	fmt.Println("\n--- Interface Vazia (interface{}) ---")

	// interface{} aceita QUALQUER tipo
	var qualquerCoisa interface{}

	qualquerCoisa = 42
	fmt.Println("Interface vazia com int:", qualquerCoisa)

	qualquerCoisa = "Hello"
	fmt.Println("Interface vazia com string:", qualquerCoisa)

	qualquerCoisa = circulo
	fmt.Println("Interface vazia com Circulo:", qualquerCoisa)

	fmt.Println("=== VANTAGENS DAS INTERFACES ===")
	fmt.Println("1. POLIMORFISMO - diferentes tipos, mesma interface")
	fmt.Println("2. DESACOPLAMENTO - código menos dependente de implementações")
	fmt.Println("3. TESTABILIDADE - facilita a criação de mocks")
	fmt.Println("4. FLEXIBILIDADE - adicionar novos tipos sem modificar código existente")
}
