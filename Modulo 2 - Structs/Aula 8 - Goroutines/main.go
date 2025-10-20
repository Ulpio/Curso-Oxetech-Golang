package main // Define como o pacote principal do programa

import (
	"fmt"       // Importa o pacote fmt para imprimir na tela
	"math/rand" // Importa o pacote rand para gerar números aleatórios
	"time"      // Importa o pacote time para trabalhar com tempo
)

// Função que simula o download de um arquivo
func baixarArquivo(nome string, c chan string) {
	tempo := time.Duration(rand.Intn(1000))
	time.Sleep(tempo * time.Millisecond)

	// Envia o resultado para o canal
	c <- fmt.Sprintf("%s baixado em %dms", nome, tempo)
}

// Função que imprime números
func imprimirNumeros(prefixo string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", prefixo, i)
		time.Sleep(100 * time.Millisecond)
	}
}

// Função que envia números para um canal
func gerarNumeros(c chan int) {
	for i := 1; i <= 5; i++ {
		c <- i // Envia número para o canal
		time.Sleep(200 * time.Millisecond)
	}
	close(c) // Fecha o canal quando terminar
}

// Função que soma números e retorna via canal
func somar(a, b int, resultado chan int) {
	time.Sleep(500 * time.Millisecond) // Simula processamento
	resultado <- a + b
}

// Função que processa dados com canal bidirecional
func processarDados(entrada chan int, saida chan string) {
	for numero := range entrada {
		resultado := fmt.Sprintf("Processado: %d x 2 = %d", numero, numero*2)
		saida <- resultado
	}
	close(saida)
}

func main() {
	fmt.Println("=== GOROUTINES E CANAIS ===")

	// GOROUTINES SIMPLES
	fmt.Println("--- Exemplo 1: Goroutines Simples ---")

	// Executando funções em paralelo com goroutines
	go imprimirNumeros("Goroutine A")
	go imprimirNumeros("Goroutine B")
	imprimirNumeros("Main") // Main também executa

	// Aguarda um pouco para ver as goroutines executarem
	time.Sleep(1 * time.Second)

	// CANAIS (CHANNELS)
	fmt.Println("\n--- Exemplo 2: Download de Arquivos (Canais) ---")

	arquivos := []string{"foto.png", "video.mp4", "musica.mp3"}
	canal := make(chan string)

	// Inicia downloads em paralelo
	for _, arquivo := range arquivos {
		go baixarArquivo(arquivo, canal)
	}

	// Recebe os resultados
	for i := 0; i < len(arquivos); i++ {
		fmt.Println(<-canal) // Bloqueia até receber um valor
	}

	// CANAIS COM RANGE E CLOSE
	fmt.Println("\n--- Exemplo 3: Canal com Range ---")

	numeros := make(chan int)
	go gerarNumeros(numeros)

	// Range itera até o canal ser fechado
	for num := range numeros {
		fmt.Println("Recebido:", num)
	}

	// MÚLTIPLAS GOROUTINES COM CANAIS
	fmt.Println("\n--- Exemplo 4: Cálculos Paralelos ---")

	resultado1 := make(chan int)
	resultado2 := make(chan int)
	resultado3 := make(chan int)

	// Executando somas em paralelo
	go somar(10, 20, resultado1)
	go somar(30, 40, resultado2)
	go somar(50, 60, resultado3)

	// Recebendo resultados (ordem não importa)
	fmt.Println("Resultado 1:", <-resultado1)
	fmt.Println("Resultado 2:", <-resultado2)
	fmt.Println("Resultado 3:", <-resultado3)

	// PIPELINE COM CANAIS
	fmt.Println("\n--- Exemplo 5: Pipeline de Processamento ---")

	entrada := make(chan int)
	saida := make(chan string)

	// Inicia o processador
	go processarDados(entrada, saida)

	// Inicia goroutine para enviar dados
	go func() {
		for i := 1; i <= 3; i++ {
			entrada <- i
		}
		close(entrada)
	}()

	// Recebe e imprime resultados processados
	for resultado := range saida {
		fmt.Println(resultado)
	}

	// CANAL COM BUFFER
	fmt.Println("\n--- Exemplo 6: Canal com Buffer ---")

	// Canal com capacidade para 3 elementos
	canalBuffer := make(chan string, 3)

	// Envia 3 valores sem bloquear (porque tem buffer)
	canalBuffer <- "Mensagem 1"
	canalBuffer <- "Mensagem 2"
	canalBuffer <- "Mensagem 3"

	fmt.Println(<-canalBuffer)
	fmt.Println(<-canalBuffer)
	fmt.Println(<-canalBuffer)

	// SELECT - MULTIPLEXAÇÃO DE CANAIS
	fmt.Println("\n--- Exemplo 7: Select ---")

	canal1 := make(chan string)
	canal2 := make(chan string)

	go func() {
		time.Sleep(300 * time.Millisecond)
		canal1 <- "Resultado do canal 1"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		canal2 <- "Resultado do canal 2"
	}()

	// Select espera pelo primeiro canal que estiver pronto
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-canal1:
			fmt.Println(msg1)
		case msg2 := <-canal2:
			fmt.Println(msg2)
		}
	}

	fmt.Println("=== CONCEITOS IMPORTANTES ===")
	fmt.Println("• Goroutines são threads leves gerenciadas pelo Go")
	fmt.Println("• Use 'go' antes de uma função para executá-la em paralelo")
	fmt.Println("• Canais permitem comunicação segura entre goroutines")
	fmt.Println("• Canais sem buffer bloqueiam até haver receptor/emissor")
	fmt.Println("• Canais com buffer só bloqueiam quando cheios/vazios")
	fmt.Println("• Use 'close()' para fechar um canal")
	fmt.Println("• Use 'select' para trabalhar com múltiplos canais")
}
